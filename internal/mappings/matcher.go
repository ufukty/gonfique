package mappings

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

type matchitem struct {
	holder  ast.Node // smallest piece of AST that holds both the type definition and the YAML key name (except *ast.ArrayType)
	pathway []string
}

func callMatcherHelperOnFields(st *ast.StructType, rule []string, pathway []string, keys map[ast.Node]string) ([]matchitem, error) {
	matches := []matchitem{}
	if st.Fields != nil && st.Fields.List != nil {
		for _, f := range st.Fields.List {
			submatches, err := matchTypeDefHolderHelper(f, rule, append(pathway, keys[f]), keys)
			if err != nil {
				return nil, fmt.Errorf("recurring to fields: %w", err)
			}
			matches = append(matches, submatches...)
		}
	}
	return matches, nil
}

// n is Field or ArrayType
// return is Field or ArrayType.
// use a typeswitch to replace .Type or .Elt fields.
func matchTypeDefHolderHelper(n ast.Node, rule []string, pathway []string, keys map[ast.Node]string) ([]matchitem, error) {
	fmt.Printf("%q  %q\n", strings.Join(pathway, "."), strings.Join(rule, "."))
	// if strings.Join(rule, ".") == "name" && pathway[len(pathway)-1] == "name" {
	// 	fmt.Println("name")
	// }
	matches := []matchitem{}

	if len(rule) == 0 {
		return matches, nil
	}

	if st, ok := n.(*ast.StructType); ok {
		return callMatcherHelperOnFields(st, rule, pathway, keys)
	}

	if len(rule) == 1 {

		switch segment := rule[0]; segment {
		case "**":
			// match everything in current depth except *ast.Ident{"string"}
			switch n.(type) {
			case *ast.ArrayType, *ast.Field:
				matches = append(matches, matchitem{n, append(pathway, segment)})
			}

			// keep the "**" for next depth in recursion
			var next ast.Expr
			var orgkey string
			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				next = f.Type
				orgkey = keys[f]
			} else if at, ok := n.(*ast.ArrayType); ok && at.Elt != nil {
				next = at.Elt
				orgkey = "[]"
			}
			if next != nil {
				submatches, err := matchTypeDefHolderHelper(next, rule, append(pathway, orgkey), keys)
				if err != nil {
					return nil, fmt.Errorf("passing the '**' for next depths of all keys or the array's item type: %w", err)
				}
				if len(submatches) > 0 {
					matches = append(matches, submatches...)
				}
			}

		case "*":
			if f, ok := n.(*ast.Field); ok {
				return []matchitem{{f, append(pathway, segment)}}, nil
			}
			fmt.Println("non-field leaf for * ending")

		case "[]":
			if at, ok := n.(*ast.ArrayType); ok {
				if at.Elt != nil {
					return []matchitem{{at, append(pathway, segment)}}, nil
				}
			}
			fmt.Println("non-array leaf for [] ending")

		default:
			if f, ok := n.(*ast.Field); ok {
				if keys[f] == segment && f.Type != nil {
					return []matchitem{{f, append(pathway, keys[f])}}, nil
				}
			}
			fmt.Println("no match for leaf")
		}

	} else if len(rule) > 1 {

		switch segment := rule[0]; segment {
		case "**": // ., *, **
			// consume the "**" at the current depth and continue recurring without it
			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				if st, ok := f.Type.(*ast.StructType); ok {
					submatches, err := callMatcherHelperOnFields(st, rule[1:], pathway, keys)
					if err != nil {
						return nil, fmt.Errorf("consuming '**' and matching all keys in current dict: %w", err)
					}
					if len(submatches) > 0 {
						matches = append(matches, submatches...)
					}
				} else if at, ok := f.Type.(*ast.ArrayType); ok && at.Elt != nil {
					submatches, err := matchTypeDefHolderHelper(at.Elt, rule[1:], append(pathway, "[]"), keys)
					if err != nil {
						return nil, fmt.Errorf("consuming '**' and recurring into array's item type: %w", err)
					}
					if len(submatches) > 0 {
						matches = append(matches, submatches...)
					}
				}
			}

			// keep the "**" for next depth in recursion
			var next ast.Expr
			var orgkey string
			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				next = f.Type
				orgkey = keys[f]
			} else if at, ok := n.(*ast.ArrayType); ok && at.Elt != nil {
				next = at.Elt
				orgkey = "[]"
			}
			if next != nil {
				submatches, err := matchTypeDefHolderHelper(next, rule, append(pathway, orgkey), keys)
				if err != nil {
					return nil, fmt.Errorf("passing the '**' for next depths of all keys or the array's item type: %w", err)
				}
				if len(submatches) > 0 {
					matches = append(matches, submatches...)
				}
			}

		case "*":
			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				if st, ok := f.Type.(*ast.StructType); ok {
					return callMatcherHelperOnFields(st, rule[1:], pathway, keys)
				}
			}

		case "[]":
			if at, ok := n.(*ast.ArrayType); ok && at.Elt != nil {
				submatches, err := matchTypeDefHolderHelper(at.Elt, rule[1:], append(pathway, "[]"), keys)
				if err != nil {
					return nil, fmt.Errorf("checking matches for '[]': %w", err)
				}
				matches = append(matches, submatches...)
			}

		default:
			if f, ok := n.(*ast.Field); ok {
				orgkey, ok := keys[f]
				if !ok {
					return nil, fmt.Errorf("could not retrieve the original keyname for %s (AST %p)", f.Names[0].Name, f)
				}
				if orgkey == segment && f.Type != nil {
					return matchTypeDefHolderHelper(f.Type, rule[1:], append(pathway, segment), keys)
				}
			}
		}
	}
	return matches, nil
}

func uniq(src []matchitem) []matchitem {
	mp := map[*matchitem]bool{}
	for i := 0; i < len(src); i++ {
		mp[&(src[i])] = true
	}
	keys := []matchitem{}
	for mi := range mp {
		keys = append(keys, *mi)
	}
	return keys
}

// root should be either of ArrayType or StructType
// accepts processed form of Config type AST which:
//   - should not have multiple names per ast.Field
//   - array types should be defined by combining compatible item fields
func matchTypeDefHolder(root ast.Expr, rule string, keys map[ast.Node]string) ([]matchitem, error) {
	switch root.(type) {
	case *ast.ArrayType, *ast.StructType:
		break
	default:
		return nil, fmt.Errorf("unsupported root type: %s", reflect.TypeOf(root).String())
	}
	segments := strings.Split(rule, ".")
	if len(segments) == 0 {
		return []matchitem{}, nil
	}
	mis, err := matchTypeDefHolderHelper(root, segments, []string{}, keys)
	if err != nil {
		return nil, fmt.Errorf("checking against rules: %w", err)
	}
	return uniq(mis), nil
}
