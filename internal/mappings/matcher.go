package mappings

import (
	"fmt"
	"go/ast"
	"log"
	"reflect"
	"slices"
	"strings"
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
func matchTypeDefHolderHelper(n ast.Node, rule []string, pathway []string, keys map[ast.Node]string) ([]matchitem, error) {
	if len(rule) == 0 {
		return []matchitem{}, nil
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
				ckey, ok := keys[f]
				if !ok {
					return nil, fmt.Errorf("original key name is not stored for generated AST node: %p", f)
				}
				mis, err := matchTypeDefHolderHelper(f, rule, append(pathway, ckey), keys)
				if err != nil {
					return nil, fmt.Errorf("checking next segments for '**': %w", err)
				}
				matches = append(matches, mis...)
				mis, err = matchTypeDefHolderHelper(f, degraded(rule), append(pathway, ckey), keys)
				if err != nil {
					return nil, fmt.Errorf("checking at segment results for '**': %w", err)
				}
				matches = append(matches, mis...)
				// TODO: add call to check "[]" appended rule
			}
		}

	case "*":
		switch t := t.(type) {
		case *ast.StructType:
			for _, f := range t.Fields.List {
				ckey, ok := keys[f]
				if !ok {
					return nil, fmt.Errorf("original key name is not stored for generated AST node: %p", f)
				}
				if len(rule) == 1 {
					matches = append(matches, matchitem{f, append(pathway, ckey)})
				} else {
					mis, err := matchTypeDefHolderHelper(f, rule[1:], append(pathway, ckey), keys)
					if err != nil {
						return nil, fmt.Errorf("checking next segments for '*': %w", err)
					}
					matches = append(matches, mis...)
				}
			}
		}

	case "[]": // other selectors like [birthday], [photo,email] are reserved for later use
		if at, ok := t.(*ast.ArrayType); ok {
			if len(rule) == 1 { // should be leaf
				matches = append(matches, matchitem{at, append(pathway, "[]")})
			} else {
				mis, err := matchTypeDefHolderHelper(at, rule[1:], append(pathway, "[]"), keys)
				if err != nil {
					return nil, fmt.Errorf("checking matches for '[]': %w", err)
				}
				matches = append(matches, mis...)
			}
		}

	default:
		if st, ok := t.(*ast.StructType); ok {
			for _, f := range st.Fields.List {
				ckey, ok := keys[f]
				if !ok {
					return nil, fmt.Errorf("could not retrieve the original keyname for %s (AST %p)", f.Names[0].Name, f)
				}
				if ckey == segment {
					if len(rule) == 1 { // should be leaf
						matches = append(matches, matchitem{f, append(pathway, ckey)})
					} else {
						mis, err := matchTypeDefHolderHelper(f, rule[1:], append(pathway, ckey), keys)
						if err != nil {
							return nil, fmt.Errorf("checking matches at next segments for %q: %w", strings.Join(rule[1:], "."), err)
						}
						matches = append(matches, mis...)
					}
				}
			}
		}
	}
	return matches, nil
}

// accepts processed form of Config type AST which:
//   - should not have multiple names per ast.Field
//   - array types should be defined by combining compatible item fields
func matchTypeDefHolder(cfg *ast.TypeSpec, rule string, keys map[ast.Node]string) ([]matchitem, error) {
	segments := strings.Split(rule, ".")
	if len(segments) == 0 {
		return []matchitem{}, nil
	}
	return matchTypeDefHolderHelper(cfg, segments, []string{}, keys)
}
