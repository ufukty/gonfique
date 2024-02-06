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

func bijective26(i int) string {
	const q = 26
	i++
	s := ""
	for i > 0 {
		i--
		s = string(byte(int(byte('A'))+i%q)) + s
		i /= q
	}
	return s
}
