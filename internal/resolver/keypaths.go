package resolver

import (
	"go/ast"
	"strings"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

type resolver struct {
	originalKeys map[ast.Node]string
	keypaths     map[ast.Node]models.FlattenKeypath
}

func newresolver(originalKeys map[ast.Node]string) *resolver {
	return &resolver{
		originalKeys: originalKeys,
		keypaths:     map[ast.Node]models.FlattenKeypath{},
	}
}

func (r *resolver) dfs(n ast.Node, holder ast.Node, path []string) {
	if holder != nil {
		r.keypaths[holder] = models.FlattenKeypath(strings.Join(path, "."))
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

func AllKeypathsForHolders(b *bundle.Bundle) {
	resolver := newresolver(b.OriginalKeys)
	resolver.dfs(b.CfgType, nil, []string{})
	b.Keypaths = resolver.keypaths
	b.Holders = datas.Invmap(b.Keypaths)
}
