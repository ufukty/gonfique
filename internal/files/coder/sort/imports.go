package sort

import (
	"go/ast"
	"slices"
	"strings"
)

// does [s] contain a dot in first segment?
func thirdparty(s string) bool {
	return strings.Contains(strings.Split(s, "/")[0], ".")
}

func Imports(specs []ast.Spec) {
	slices.SortFunc(specs, func(a, b ast.Spec) int {
		sa := a.(*ast.ImportSpec).Path.Value
		sb := b.(*ast.ImportSpec).Path.Value
		if !thirdparty(sa) && thirdparty(sb) {
			return -1
		} else if thirdparty(sa) && !thirdparty(sb) {
			return 1
		} else if sa < sb {
			return -1
		} else if sa == sb {
			return 0
		} else {
			return 1
		}
	})
}
