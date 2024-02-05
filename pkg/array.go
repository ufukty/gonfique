package pkg

import (
	"fmt"
	"go/ast"
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
		t := toAst(iv)
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
