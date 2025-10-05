package declare

import (
	"cmp"
	"fmt"
	"go/ast"
	"slices"

	"go.ufukty.com/gonfique/internal/compares"
	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/paths/mapper/resolve"
	"golang.org/x/exp/maps"
)

func groupSchemas(users []resolve.Path, exprs map[resolve.Path]ast.Expr) map[ast.Expr][]resolve.Path {
	groups := map[ast.Expr][]resolve.Path{}
	for _, rp := range users {
		found := false
		for alt := range groups {
			if compares.Compare(alt, exprs[rp]) {
				groups[alt] = append(groups[alt], rp)
				found = true
				break
			}
		}
		if !found {
			groups[exprs[rp]] = []resolve.Path{rp}
		}
	}
	return groups
}

func group(users map[config.Typename][]resolve.Path, exprs map[resolve.Path]ast.Expr) map[config.Typename]map[ast.Expr][]resolve.Path {
	schemas := map[config.Typename]map[ast.Expr][]resolve.Path{}
	for tn, rps := range users {
		schemas[tn] = groupSchemas(rps, exprs)
	}
	return schemas
}

func ternary(cond bool, t, f string) string {
	if cond {
		return t
	}
	return f
}

// summarize further (doesn't recur on structs)
func further(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.ArrayType:
		return "[]" + further(t.Elt)
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", further(t.X), t.Sel.Name)
	case *ast.StructType:
		return "struct{...}"
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", further(t.Key), further(t.Value))
	default:
		return "..."
	}
}

// prints fields for structs only in first depth
func summarize(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.StructType:
		msg := "struct{ "
		for i, f := range t.Fields.List {
			for i, id := range f.Names {
				msg += id.Name
				if i != len(f.Names)-1 {
					msg += ", "
				}
			}
			msg += " " + further(f.Type)
			if i != len(t.Fields.List)-1 {
				msg += "; "
			}
		}
		msg += " }"
		return msg
	default:
		return further(t)
	}
}

// prints targets
// with conflicting schemas
// directed to be declared with same typename
// in directory tree format
// in sorted order
func format(cs map[config.Typename]map[ast.Expr][]resolve.Path) string {
	msg := ""
	tns := maps.Keys(cs)
	slices.Sort(tns)
	for i := 0; i < len(tns); i++ {
		heading := ternary(i != len(tns)-1, "├─ ", "╰─ ")
		inherit := ternary(i != len(tns)-1, "│  ", "   ")

		users := cs[tns[i]]
		msg += fmt.Sprintf("%sdeclared typename: %s\n", heading, tns[i])
		summaries := map[ast.Node]string{}
		for t := range users {
			summaries[t] = summarize(t)
		}
		types := maps.Keys(users)
		slices.SortFunc(types, func(a, b ast.Expr) int {
			return cmp.Compare(summaries[a], summaries[b])
		})
		for j := 0; j < len(types); j++ {
			heading := inherit + ternary(j != len(types)-1, "├─ ", "╰─ ")
			inherit := inherit + ternary(j != len(types)-1, "│  ", "   ")

			msg += fmt.Sprintf("%stype expression: %s\n", heading, summaries[types[j]])
			rps := users[types[j]]
			slices.Sort(rps)
			for k, rp := range rps {
				heading := inherit + ternary(k != len(rps)-1, "├─ ", "╰─ ")
				msg += fmt.Sprintf("%ssource: %s\n", heading, rp)
			}
		}
	}
	return msg
}

func (a *Agent) Conflicts() error {
	schemas := group(a.users, a.exprs)

	cs := map[config.Typename]map[ast.Expr][]resolve.Path{}
	for tn, groups := range schemas {
		if len(groups) > 1 {
			cs[tn] = groups
		}
	}
	if len(cs) == 0 {
		return nil
	}
	return fmt.Errorf("%s", format(cs))
}
