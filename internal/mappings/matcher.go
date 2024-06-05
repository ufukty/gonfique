package mappings

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/ufukty/gonfique/internal/files"
)

func callMatcherHelperOnFields(st *ast.StructType, rule []string, keys map[ast.Node]string) ([]ast.Node, error) {
	matches := []ast.Node{}
	if st.Fields != nil && st.Fields.List != nil {
		for _, f := range st.Fields.List {
			submatches, err := matchTypeDefHolderHelper(f, rule, keys)
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
func matchTypeDefHolderHelper(n ast.Node, rule []string, keys map[ast.Node]string) ([]ast.Node, error) {
	matches := []ast.Node{}

	if len(rule) == 0 {
		return matches, nil
	}

	if st, ok := n.(*ast.StructType); ok {
		return callMatcherHelperOnFields(st, rule, keys)
	}

	if len(rule) == 1 {

		switch segment := rule[0]; segment {
		case "**":
			// match everything in current depth except *ast.Ident{"string"}

			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				matches = append(matches, n)
			} else if at, ok := n.(*ast.ArrayType); ok && at.Elt != nil {
				matches = append(matches, at)
			}

			// keep the "**" for next depth in recursion
			var next ast.Expr
			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				next = f.Type
			} else if at, ok := n.(*ast.ArrayType); ok && at.Elt != nil {
				next = at.Elt
			}
			if next != nil {
				submatches, err := matchTypeDefHolderHelper(next, rule, keys)
				if err != nil {
					return nil, fmt.Errorf("passing the '**' for next depths of all keys or the array's item type: %w", err)
				}
				if len(submatches) > 0 {
					matches = append(matches, submatches...)
				}
			}

		case "*":
			if f, ok := n.(*ast.Field); ok {
				return []ast.Node{f}, nil
			}

		case "[]":
			if at, ok := n.(*ast.ArrayType); ok {
				return []ast.Node{at}, nil
			}

		default:
			if f, ok := n.(*ast.Field); ok && keys[f] == segment {
				return []ast.Node{f}, nil
			}
		}

	} else if len(rule) > 1 {

		switch segment := rule[0]; segment {
		case "**": // ., *, **
			// consume the "**" at the current depth and continue recurring without it

			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				var next ast.Expr
				if st, ok := f.Type.(*ast.StructType); ok && st.Fields != nil && st.Fields.List != nil {
					next = st
				} else if at, ok := f.Type.(*ast.ArrayType); ok && at.Elt != nil {
					next = at.Elt
				}
				if next != nil {
					submatches, err := matchTypeDefHolderHelper(next, rule[1:], keys)
					if err != nil {
						return nil, fmt.Errorf("consuming '**' and matching all keys in current dict: %w", err)
					}
					if len(submatches) > 0 {
						matches = append(matches, submatches...)
					}
				}
			}

			// keep the "**" for next depth in recursion
			var next ast.Expr
			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				next = f.Type
			} else if at, ok := n.(*ast.ArrayType); ok && at.Elt != nil {
				next = at.Elt
			}
			if next != nil {
				submatches, err := matchTypeDefHolderHelper(next, rule, keys)
				if err != nil {
					return nil, fmt.Errorf("passing the '**' for next depths of all keys or the array's item type: %w", err)
				}
				if len(submatches) > 0 {
					matches = append(matches, submatches...)
				}
			}

		case "*":
			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				return matchTypeDefHolderHelper(f.Type, rule[1:], keys)
			}

		case "[]":
			if at, ok := n.(*ast.ArrayType); ok && at.Elt != nil {
				submatches, err := matchTypeDefHolderHelper(at.Elt, rule[1:], keys)
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
					return matchTypeDefHolderHelper(f.Type, rule[1:], keys)
				}
			}
		}
	}
	return matches, nil
}

func uniq(nodes []ast.Node) []ast.Node {
	mnodes := map[ast.Node]bool{}
	for i := 0; i < len(nodes); i++ {
		mnodes[nodes[i]] = true
	}
	keys := []ast.Node{}
	for mn := range mnodes {
		keys = append(keys, mn)
	}
	return keys
}

// root should be either of ArrayType or StructType
// accepts processed form of Config type AST which:
//   - should not have multiple names per ast.Field
//   - array types should be defined by combining compatible item fields
func matchTypeDefHolder(root ast.Expr, rule files.Keypath, keys map[ast.Node]string) ([]ast.Node, error) {
	switch root.(type) {
	case *ast.ArrayType, *ast.StructType:
		break
	default:
		return nil, fmt.Errorf("unsupported root type: %s", reflect.TypeOf(root).String())
	}
	segments := strings.Split(string(rule), ".")
	if len(segments) == 0 {
		return []ast.Node{}, fmt.Errorf("empty rule %q", rule)
	}
	mis, err := matchTypeDefHolderHelper(root, segments, keys)
	if err != nil {
		return nil, fmt.Errorf("checking against rules: %w", err)
	}
	return uniq(mis), nil
}
