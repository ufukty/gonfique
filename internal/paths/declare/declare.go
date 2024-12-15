package declare

import (
	"cmp"
	"fmt"
	"go/ast"
	"go/token"
	"slices"

	"github.com/ufukty/gonfique/internal/compares"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"golang.org/x/exp/maps"
)

func getType(holder ast.Node) ast.Expr {
	switch h := holder.(type) {
	case *ast.Field:
		return h.Type
	case *ast.ArrayType:
		return h.Elt
	default:
		panic("implementation error")
	}
}

func groupSchemas(users []resolve.Path, holders map[resolve.Path]ast.Node) map[ast.Expr][]resolve.Path {
	groups := map[ast.Expr][]resolve.Path{}
	for _, rp := range users {
		found := false
		for alt := range groups {
			if compares.Compare(alt, getType(holders[rp])) {
				groups[alt] = append(groups[alt], rp)
				found = true
				break
			}
		}
		if !found {
			groups[getType(holders[rp])] = []resolve.Path{rp}
		}
	}
	return groups
}

func group(directives map[resolve.Path]config.Typename, holders map[resolve.Path]ast.Node) map[config.Typename]map[ast.Expr][]resolve.Path {
	schemas := map[config.Typename]map[ast.Expr][]resolve.Path{}
	for tn, rps := range datas.Revmap(directives) {
		schemas[tn] = groupSchemas(rps, holders)
	}
	return schemas
}

func ternary(cond bool, t, f string) string {
	if cond {
		return t
	}
	return f
}

// doesn't recur on structs
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
	default:
		return "..."
	}
}

func summarize(n ast.Node) string {
	msg := ""
	switch n := n.(type) {
	case *ast.StructType:
		msg += "struct{"
		for i, f := range n.Fields.List {
			for i, id := range f.Names {
				msg += id.Name
				if i != len(f.Names)-1 {
					msg += ", "
				}
			}
			msg += " " + further(f.Type)
			if i != len(n.Fields.List)-1 {
				msg += "; "
			}
		}
		msg += "}"
	case *ast.ArrayType:
		msg += "[]" + further(n.Elt)
	case *ast.Ident:
		msg += n.Name
	}
	return msg
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
		heading := ternary(i != len(tns)-1, "├── ", "└── ")
		inherit := ternary(i != len(tns)-1, "|   ", "    ")

		users := cs[tns[i]]
		msg += fmt.Sprintf("%stypename: %s\n", heading, tns[i])
		summaries := map[ast.Node]string{}
		for t := range users {
			summaries[t] = summarize(t)
		}
		types := maps.Keys(users)
		slices.SortFunc(types, func(a, b ast.Expr) int {
			return cmp.Compare(summaries[a], summaries[b])
		})
		for j := 0; j < len(types); j++ {
			heading := inherit + ternary(j != len(types)-1, "├── ", "└── ")
			inherit := inherit + ternary(j != len(types)-1, "|   ", "    ")

			msg += fmt.Sprintf("%stype expression: %s\n", heading, summaries[types[j]])
			rps := users[types[j]]
			slices.Sort(rps)
			for k, rp := range rps {
				heading := inherit + ternary(k != len(rps)-1, "├── ", "└── ")
				msg += fmt.Sprintf("%s%s\n", heading, rp)
			}
		}
	}
	return msg
}

func conflicts(schemas map[config.Typename]map[ast.Expr][]resolve.Path) error {
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

func pick(schemas map[config.Typename]map[ast.Expr][]resolve.Path) map[config.Typename]ast.Expr {
	picks := map[config.Typename]ast.Expr{}
	for tn, users := range schemas {
		picks[tn] = maps.Keys(users)[0]
	}
	return picks
}

func set(holder ast.Node, typeExpr ast.Expr) {
	switch h := holder.(type) {
	case *ast.Field:
		h.Type = typeExpr
	case *ast.ArrayType:
		h.Elt = typeExpr
	}
}

func Declare(directives map[resolve.Path]config.Typename, holders map[resolve.Path]ast.Node) ([]*ast.GenDecl, error) {
	schemas := group(directives, holders)

	err := conflicts(schemas)
	if err != nil {
		return nil, fmt.Errorf("checking conflicting schemas:\n%s", err.Error())
	}

	picks := pick(schemas)
	decls := []*ast.GenDecl{}

	for tn, exp := range picks {
		decls = append(decls, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{Name: tn.Ident(), Type: exp},
			},
		})
	}

	for rp, tn := range directives {
		set(holders[rp], tn.Ident())
	}

	return decls, nil
}
