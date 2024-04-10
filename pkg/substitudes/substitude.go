package substitudes

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/ufukty/gonfique/pkg/compares"
	"golang.org/x/tools/go/ast/astutil"
)

func ReadTypes(path string) ([]*ast.TypeSpec, error) {
	f, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %w", path, err)
	}
	tss := []*ast.TypeSpec{}
	for _, decl := range f.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok {
			if gd.Tok == token.TYPE {
				for _, spec := range gd.Specs {
					if ts, ok := spec.(*ast.TypeSpec); ok {
						tss = append(tss, ts)
					}
				}
			}
		}
	}
	return tss, nil
}

func Substitute(produced *ast.TypeSpec, existing []*ast.TypeSpec) {
	// substitute on dfs traceback
	astutil.Apply(produced.Type, nil, func(c *astutil.Cursor) bool {
		for _, e := range existing {
			if c.Node() != nil && compares.Compare(c.Node(), e.Type) {
				c.Replace(e.Name)
			}
		}
		return true
	})
}
