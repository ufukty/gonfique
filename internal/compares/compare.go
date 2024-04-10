package compares

import (
	"go/ast"
	"slices"

	"golang.org/x/exp/maps"
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

func linearizeFieldList(fl *ast.FieldList) map[*ast.Ident]ast.Expr {
	l := map[*ast.Ident]ast.Expr{}
	for _, f := range fl.List {
		for _, id := range f.Names {
			l[id] = f.Type
		}
	}
	return l
}

func Compare(a, b any) bool {
	switch a := a.(type) {
	case *ast.Ident:
		if b, ok := b.(*ast.Ident); ok {
			return a.Name == b.Name
		}
	case *ast.ArrayType:
		if b, ok := b.(*ast.ArrayType); ok {
			return (a.Len == b.Len) && (a.Elt != nil && b.Elt != nil && Compare(a.Elt, b.Elt))
		}
	case *ast.MapType:
		if b, ok := b.(*ast.MapType); ok {
			return Compare(a.Key, b.Key) && Compare(a.Value, b.Value)
		}
	case *ast.FieldList:
		if b, ok := b.(*ast.FieldList); ok {
			if a.NumFields() != b.NumFields() {
				return false
			}
			af := linearizeFieldList(a)
			bf := linearizeFieldList(b)
			afk := maps.Keys(af)
			bfk := maps.Keys(bf)
			if len(afk) != len(bfk) {
				return false
			}
			sort(afk)
			sort(bfk)
			for i := 0; i < len(afk); i++ {
				if !Compare(afk[i], bfk[i]) || !Compare(af[afk[i]], bf[bfk[i]]) {
					return false
				}
			}
			return true
		}
	case *ast.StructType:
		if b, ok := b.(*ast.StructType); ok {
			return Compare(a.Fields, b.Fields)
		}
	}
	return false
}
