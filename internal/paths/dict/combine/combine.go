package combine

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas/inits"
	"github.com/ufukty/gonfique/internal/transform"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func Combine(ti *transform.Info, types ...ast.Expr) (ast.Expr, error) {
	if len(types) == 0 {
		return nil, fmt.Errorf("no argument")
	}
	if len(types) == 1 {
		return types[0], nil
	}

	switch types[0].(type) {
	case *ast.StructType:
		fls := map[*ast.StructType][]*ast.Field{}
		for _, t := range types {
			st, ok := t.(*ast.StructType)
			if !ok {
				return nil, fmt.Errorf("expected StructType got %T", t)
			}
			if st.Fields == nil || st.Fields.List == nil {
				return nil, fmt.Errorf("uninitialized field list")
			}
			for _, f := range st.Fields.List {
				if f.Type == nil {
					return nil, fmt.Errorf("uninitialized field type in %q", ti.Keys[f.Type])
				}
			}
			fls[st] = st.Fields.List
		}

		sources := map[string]map[*ast.StructType]*ast.Field{} // field sources for keys
		for st, fl := range fls {
			for _, f := range fl {
				inits.Key2(sources, ti.Keys[f], st)
				sources[ti.Keys[f]][st] = f
			}
		}

		fs := []*ast.Field{}
		for key, sts := range sources {
			types := []ast.Expr{}
			var mold *ast.Field
			for _, f := range sts {
				types = append(types, f.Type)
				if mold == nil { // first
					mold = f
				}
			}
			combined, err := Combine(ti, types...)
			if err != nil {
				return nil, fmt.Errorf("key %q: %w", key, err)
			}
			fs = append(fs, &ast.Field{Names: mold.Names, Tag: mold.Tag, Type: combined})
			ti.Keys[combined] = ti.Keys[mold]
		}
		return &ast.StructType{Fields: &ast.FieldList{List: fs}}, nil

	case *ast.ArrayType:
		elts := []ast.Expr{}
		var mold ast.Expr
		for _, t := range types {
			at, ok := t.(*ast.ArrayType)
			if !ok {
				return nil, fmt.Errorf("expected ArrayType got %T", t)
			}
			if at.Elt == nil {
				return nil, fmt.Errorf("uninitialized field list")
			}
			elts = append(elts, at.Elt)
			if mold == nil {
				mold = at.Elt
			}
		}
		elt, err := Combine(ti, elts...)
		if err != nil {
			return nil, fmt.Errorf("element: %w", err)
		}
		ti.Keys[elt] = ti.Keys[mold]
		return &ast.ArrayType{Elt: elt}, nil

	case *ast.Ident:
		var i *ast.Ident
		for _, t := range types {
			i2, ok := t.(*ast.Ident)
			if !ok {
				return nil, fmt.Errorf("expected Ident, got %T", t)
			}
			if i == nil {
				i = i2
			} else if i2.Name != i.Name {
				return nil, fmt.Errorf("ident mismatch %q != %q", i, i2)
			}
		}
		return i, nil

	default:
		return nil, fmt.Errorf("unsupported type (%T)", types[0])
	}
}
