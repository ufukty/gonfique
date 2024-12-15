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

func FuncDecls(named []*ast.GenDecl) {
	slices.SortFunc(named, func(a, b *ast.GenDecl) int {
		if a.Specs[0].(*ast.TypeSpec).Name.Name > b.Specs[0].(*ast.TypeSpec).Name.Name {
			return 1
		} else {
			return -1
		}
	})
}

// sorts accessors by receiver name first then method name
func Accessors(accessors []*ast.FuncDecl) {
	slices.SortFunc(accessors, func(i, j *ast.FuncDecl) int {
		it, jt := i.Recv.List[0].Type, j.Recv.List[0].Type
		if se, ok := it.(*ast.StarExpr); ok {
			it = se.X
		}
		if se, ok := jt.(*ast.StarExpr); ok {
			jt = se.X
		}
		in := it.(*ast.Ident).Name
		jn := jt.(*ast.Ident).Name
		if in < jn {
			return -1
		} else if in > jn {
			return 1
		}
		ifn := i.Name.Name
		jfn := j.Name.Name
		if i.Name.Name < j.Name.Name {
			return -1
		} else if ifn > jfn {
			return 1
		}
		return 0
	})
}
