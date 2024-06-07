package resolver

import (
	"go/ast"
	"strings"

	"github.com/ufukty/gonfique/internal/models"
)

type resolver struct {
	originalKeys map[ast.Node]string
	keypaths     map[ast.Node]models.Keypath
}

func newresolver(originalKeys map[ast.Node]string) *resolver {
	return &resolver{
		originalKeys: originalKeys,
		keypaths:     map[ast.Node]models.Keypath{},
	}
}

func (r *resolver) dfs(n ast.Node, holder ast.Node, path []string) {
	if holder != nil {
		r.keypaths[holder] = models.Keypath(strings.Join(path, "."))
	}

	switch n := n.(type) {
	case *ast.StructType:
		if n.Fields != nil && n.Fields.List != nil {
			for _, f := range n.Fields.List {
				if f != nil && f.Type != nil {
					r.dfs(f.Type, f, append(path, r.originalKeys[f]))
				}
			}
		}

	case *ast.ArrayType:
		r.dfs(n.Elt, n, append(path, "[]"))
	}
}

func AllKeypathsForHolders(cfgtype ast.Expr, originalKeys map[ast.Node]string) map[ast.Node]models.Keypath {
	resolver := newresolver(originalKeys)
	resolver.dfs(cfgtype, nil, []string{})
	return resolver.keypaths
}
