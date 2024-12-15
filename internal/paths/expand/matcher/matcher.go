package matcher

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/transform"
)

type matcher struct {
	ti *transform.Info
}

func (m matcher) onFields(st *ast.StructType, kp []string) ([]ast.Node, error) {
	matches := []ast.Node{}
	if st.Fields != nil && st.Fields.List != nil {
		for _, f := range st.Fields.List {
			submatches, err := m.holders(f, kp)
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
func (m matcher) holders(n ast.Node, path []string) ([]ast.Node, error) {
	matches := []ast.Node{}

	if len(path) == 0 {
		return matches, nil
	}

	if st, ok := n.(*ast.StructType); ok {
		return m.onFields(st, path)
	}

	if len(path) == 1 {

		switch segment := path[0]; segment {
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
				submatches, err := m.holders(next, path)
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
			if f, ok := n.(*ast.Field); ok && m.ti.Keys[f] == segment {
				return []ast.Node{f}, nil
			}
		}

	} else if len(path) > 1 {

		switch segment := path[0]; segment {
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
					submatches, err := m.holders(next, path[1:])
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
				submatches, err := m.holders(next, path)
				if err != nil {
					return nil, fmt.Errorf("passing the '**' for next depths of all keys or the array's item type: %w", err)
				}
				if len(submatches) > 0 {
					matches = append(matches, submatches...)
				}
			}

		case "*":
			if f, ok := n.(*ast.Field); ok && f.Type != nil {
				return m.holders(f.Type, path[1:])
			}

		case "[]":
			if at, ok := n.(*ast.ArrayType); ok && at.Elt != nil {
				submatches, err := m.holders(at.Elt, path[1:])
				if err != nil {
					return nil, fmt.Errorf("checking matches for '[]': %w", err)
				}
				matches = append(matches, submatches...)
			}

		default:
			if f, ok := n.(*ast.Field); ok {
				orgkey, ok := m.ti.Keys[f]
				if !ok {
					return nil, fmt.Errorf("could not retrieve the original keyname for %s (AST %p)", f.Names[0].Name, f)
				}
				if orgkey == segment && f.Type != nil {
					return m.holders(f.Type, path[1:])
				}
			}
		}
	}
	return matches, nil
}

// Find type definition holderss for path
//
// accepts processed form of Config type AST which:
//   - should not have multiple names per ast.Field
//   - array types should be defined by combining compatible item fields
func (m matcher) FindHolders(p config.Path) ([]ast.Node, error) {
	segments := p.Segments()
	if len(segments) == 0 {
		return nil, fmt.Errorf("empty path")
	}
	mis, err := m.holders(m.ti.Type, segments)
	if err != nil {
		return nil, fmt.Errorf("holders: %w", err)
	}
	return datas.Uniq(mis), nil
}

func New(ti *transform.Info) matcher {
	return matcher{
		ti: ti,
	}
}