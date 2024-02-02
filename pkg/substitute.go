package pkg

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/exp/maps"

	"golang.org/x/tools/go/ast/astutil"
)

func linearizeFieldList(fl *ast.FieldList) map[*ast.Ident]ast.Expr {
	l := map[*ast.Ident]ast.Expr{}
	for _, f := range fl.List {
		for _, id := range f.Names {
			l[id] = f.Type
		}
	}
	return l
}

func compare(a, b any) bool {
	switch a := a.(type) {
	case *ast.Ident:
		if b, ok := b.(*ast.Ident); ok {
			return a.Name == b.Name
		}
	case *ast.ArrayType:
		if b, ok := b.(*ast.ArrayType); ok {
			return (a.Len == b.Len) && (a.Elt != nil && b.Elt != nil && compare(a.Elt, b.Elt))
		}
	case *ast.MapType:
		if b, ok := b.(*ast.MapType); ok {
			return compare(a.Key, b.Key) && compare(a.Value, b.Value)
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
				if !compare(afk[i], bfk[i]) || !compare(af[afk[i]], bf[bfk[i]]) {
					return false
				}
			}
			return true
		}
	case *ast.StructType:
		if b, ok := b.(*ast.StructType); ok {
			return compare(a.Fields, b.Fields)
		}
	}
	return false
}

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
			if c.Node() != nil && compare(c.Node(), e.Type) {
				c.Replace(e.Name)
			}
		}
		return true
	})
}
