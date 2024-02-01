package pkg

import (
	"go/ast"
	"slices"
)

func sort(idts []*ast.Ident) {
	slices.SortFunc(idts, func(x, y *ast.Ident) int {
		if x.Name > y.Name {
			return 1
		} else if x.Name < y.Name {
			return -1
		} else {
			return 0
		}
	})
}
