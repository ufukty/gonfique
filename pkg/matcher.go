package pkg

import (
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

// return items are either *ast.Field or *ast.ArrayType.
// use a typeswitch to replace .Type or .Elt fields.
func match(n ast.Node, rule []string) []ast.Node {
	if len(rule) == 0 {
		return []ast.Node{}
	}

	var t ast.Expr
	switch n := n.(type) {
	case *ast.TypeSpec:
		t = n.Type
	case *ast.Field:
		t = n.Type
	// case *ast.Ident:
	// 	log.Fatalln("to implement 2")
	case *ast.ArrayType: // only when the rule was "[]" previously
		t = n.Elt
	default:
		log.Fatalln("unhandled type", reflect.TypeOf(n))
	}

	matches := []ast.Node{}

	switch segment := rule[0]; segment {
	case "**": // works like Levenshtein (DP)
		if st, ok := t.(*ast.StructType); ok {
			for _, f := range st.Fields.List {
				matches = append(matches, match(f, rule)...)
				matches = append(matches, match(f, degraded(rule))...)
				// TODO: add call to check "[]" appended rule
			}
		}

	case "*":
		switch t := t.(type) {
		case *ast.StructType:
			for _, f := range t.Fields.List {
				matches = append(matches, match(f, rule[1:])...)
			}
		}

	case "[]": // other selectors like [birthday], [photo,email] are reserved for later use
		if at, ok := t.(*ast.ArrayType); ok {
			if len(rule) == 1 { // should be leaf
				matches = append(matches, n)
			} else {
				matches = append(matches, match(at, rule[1:])...)
			}
		}

	default:
		if st, ok := t.(*ast.StructType); ok {
			for _, f := range st.Fields.List {
				ckey, err := stripKeyname(f.Tag.Value)
				if err != nil {
					log.Fatalf("could not get the key name out of field tag for %s", f.Tag.Value)
				}
				if ckey == segment {
					if len(rule) == 1 { // should be leaf
						matches = append(matches, f)
					} else {
						matches = append(matches, match(f, rule[1:])...)
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
func Match(cfg *ast.TypeSpec, rule string) []ast.Node {
	segments := strings.Split(rule, ".")
	if len(segments) == 0 {
		return []ast.Node{}
	} else if l := segments[len(segments)-1]; l == "*" || l == "**" {
		return []ast.Node{}
	}
	cds := match(cfg, segments)
	return cds
}
