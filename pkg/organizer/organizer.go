package organizer

import (
	"go/ast"
	"go/token"
	"slices"

	"github.com/ufukty/gonfique/pkg/compares"
	"golang.org/x/tools/go/ast/astutil"
)

func Organize(ts *ast.TypeSpec) *ast.GenDecl {
	store := map[*ast.StructType]*ast.Ident{}
	prevs := []*ast.StructType{}
	astutil.Apply(ts.Type, nil, func(c *astutil.Cursor) bool {
		if c.Node() != nil && c.Node() != ts.Type {
			if st, ok := c.Node().(*ast.StructType); ok {
				i := slices.IndexFunc(prevs, func(prev *ast.StructType) bool {
					return compares.Compare(prev, st)
				})
				if i != -1 {
					c.Replace(store[prevs[i]])
				} else {
					name := ast.NewIdent("autoGen" + bijective26(len(store)))
					prevs = append(prevs, st)
					store[st] = name
					c.Replace(name)
				}
				return true
			}
		}
		return true
	})
	if len(prevs) == 0 {
		return nil
	}
	var gd = &ast.GenDecl{
		Doc: &ast.CommentGroup{[]*ast.Comment{
			{Text: "// IMPORTANT:"},
			{Text: "// Types are defined only for internal purposes."},
			{Text: "// Do not refer auto generated type names from outside."},
			{Text: "// Because they will change as config schema changes."}}},
		Tok:   token.TYPE,
		Specs: []ast.Spec{},
	}
	for _, st := range prevs {
		gd.Specs = append(gd.Specs, &ast.TypeSpec{
			Name: store[st],
			Type: st,
		})
	}
	return gd
}
