package pkg

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"reflect"
	"strings"

	"github.com/ufukty/gonfique/pkg/fieldlist"
)

func fieldsByTags(fl *ast.FieldList) map[string]*ast.Field {
	tgs := map[string]*ast.Field{}
	for _, f := range fl.List {
		tgs[f.Tag.Value] = f
	}
	return tgs
}

func areMergeable(a, b *ast.FieldList) error {
	bfs := fieldsByTags(b)
	conflicts := []string{}
	for _, af := range a.List {
		if bf, ok := bfs[af.Tag.Value]; ok {
			if !compare(af.Type, bf.Type) {
				conflicts = append(conflicts, af.Names[0].Name) // FIXME: ".Name" is the transformed version of the user-provided key
			}
		}
	}
	if len(conflicts) > 0 {
		return fmt.Errorf(strings.Join(conflicts, ", "))
	}
	return nil
}

func arrayType(v reflect.Value) ast.Expr {
	var m ast.Expr
	for i := 0; i < v.Len(); i++ {
		iv := v.Index(i)
		t := Transform(iv)
		if m == nil {
			m = t
			continue
		}
		stM, isSructM := m.(*ast.StructType)
		stT, isSructT := t.(*ast.StructType)
		if isSructT && isSructM {
			err := areMergeable(stM.Fields, stT.Fields)
			if err != nil {
				log.Println(fmt.Errorf("assigning 'any' to array type because of at least 2 items' type are different: %w", err))
				return &ast.ArrayType{Elt: ast.NewIdent("any")}
			} else {
				m = &ast.StructType{Fields: fieldlist.Combine(stM.Fields, stT.Fields)}
			}
		}
	}
	if m == nil {
		return &ast.ArrayType{Elt: ast.NewIdent("any")}
	}
	return &ast.ArrayType{Elt: m}
}

func structType(v reflect.Value) *ast.StructType {
	st := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{},
		},
	}
	iter := v.MapRange()
	for iter.Next() {
		ik := iter.Key()
		iv := iter.Value()
		st.Fields.List = append(st.Fields.List, &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(safeFieldName(ik.String()))},
			Type:  Transform(iv),
			Tag: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("`yaml:%q`", ik.String()),
			},
		})
		fieldlist.Sort(st.Fields)
	}
	return st
}

// reconstructs a reflect-value's type in ast.TypeSpec.
// limited with types used by YAML decoder.
func Transform(v reflect.Value) ast.Expr {
	t := v.Type()
	switch t.Kind() {
	case reflect.Interface:
		return Transform(v.Elem())
	case reflect.Map:
		return structType(v)
	case reflect.Slice:
		return arrayType(v)
	case reflect.Bool:
		return ast.NewIdent("bool")
	case reflect.String:
		return ast.NewIdent("string")
	case reflect.Int:
		return ast.NewIdent("int")
	case reflect.Int32:
		return ast.NewIdent("int32")
	case reflect.Int64:
		return ast.NewIdent("int64")
	case reflect.Uint:
		return ast.NewIdent("uint")
	case reflect.Uint32:
		return ast.NewIdent("uint32")
	case reflect.Uint64:
		return ast.NewIdent("uint64")
	case reflect.Float32:
		return ast.NewIdent("float32")
	case reflect.Float64:
		return ast.NewIdent("float64")
	default:
		log.Println("unhandled reflect kind", t)
	}
	return nil
}
