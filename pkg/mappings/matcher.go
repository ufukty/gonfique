package mappings

import (
	"go/ast"
	"log"
	"reflect"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/pkg/namings"
)

func degraded(keys []string) []string {
	keys2 := slices.Clone(keys)
	keys2[0] = "*" // replaces "**"
	return keys2
}

// type Config struct { NOTE: spec
//   Host string `yaml:"host"`
//   Http struct { NOTE: field
//     Paths []struct { NOTE: field -> arrayType
//       Backend struct {
//         Service struct {
//           Name string `yaml:"name"`
//           Port struct {
//             Number int `yaml:"number"`
//           } `yaml:"port"`
//         } `yaml:"service"`
//       } `yaml:"backend"`
//       Path     string `yaml:"path"`
//       PathType string `yaml:"pathType"`
//     } `yaml:"paths"`
//   } `yaml:"http"`
// }

type matchitem struct {
	holder  ast.Node
	pathway []string
}

// return items are either *ast.Field or *ast.ArrayType.
// use a typeswitch to replace .Type or .Elt fields.
func matchTypeDefHolder(n ast.Node, rule []string, pathway []string) []matchitem {
	if len(rule) == 0 {
		return []matchitem{}
	}

	var t ast.Expr
	switch n := n.(type) {
	case *ast.TypeSpec:
		t = n.Type
	case *ast.Field:
		t = n.Type
	case *ast.ArrayType: // only when the rule was "[]" previously
		t = n.Elt
	default:
		log.Fatalln("unhandled type", reflect.TypeOf(n))
	}

	matches := []matchitem{}

	switch segment := rule[0]; segment {
	case "**": // works like Levenshtein (DP)
		if st, ok := t.(*ast.StructType); ok {
			for _, f := range st.Fields.List {
				ckey, err := namings.StripKeyname(f.Tag.Value)
				if err != nil {
					log.Fatalf("could not get the key name out of field tag for %s", f.Tag.Value)
				}
				matches = append(matches, matchTypeDefHolder(f, rule, append(pathway, ckey))...)
				matches = append(matches, matchTypeDefHolder(f, degraded(rule), append(pathway, ckey))...)
				// TODO: add call to check "[]" appended rule
			}
		}

	case "*":
		switch t := t.(type) {
		case *ast.StructType:
			for _, f := range t.Fields.List {
				ckey, err := namings.StripKeyname(f.Tag.Value)
				if err != nil {
					log.Fatalf("could not get the key name out of field tag for %s", f.Tag.Value)
				}
				if len(rule) == 1 {
					matches = append(matches, matchitem{f, append(pathway, ckey)})
				} else {
					matches = append(matches, matchTypeDefHolder(f, rule[1:], append(pathway, ckey))...)
				}
			}
		}

	case "[]": // other selectors like [birthday], [photo,email] are reserved for later use
		if at, ok := t.(*ast.ArrayType); ok {
			if len(rule) == 1 { // should be leaf
				matches = append(matches, matchitem{at, append(pathway, "[]")})
			} else {
				matches = append(matches, matchTypeDefHolder(at, rule[1:], append(pathway, "[]"))...)
			}
		}

	default:
		if st, ok := t.(*ast.StructType); ok {
			for _, f := range st.Fields.List {
				ckey, err := namings.StripKeyname(f.Tag.Value)
				if err != nil {
					log.Fatalf("could not get the key name out of field tag for %s", f.Tag.Value)
				}
				if ckey == segment {
					if len(rule) == 1 { // should be leaf
						matches = append(matches, matchitem{f, append(pathway, ckey)})
					} else {
						matches = append(matches, matchTypeDefHolder(f, rule[1:], append(pathway, ckey))...)
					}
				}
			}
		}
	}
	return matches
}

// accepts processed form of Config type AST which:
//   - should not have multiple names per ast.Field
//   - array types should be defined by combining compatible item fields
func MatchTypeDefinitionHolder(cfg *ast.TypeSpec, rule string) []matchitem {
	segments := strings.Split(rule, ".")
	if len(segments) == 0 {
		return []matchitem{}
	}
	return matchTypeDefHolder(cfg, segments, []string{})
}
