package transform

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
)

type transformer struct {
	isTimeUsed bool
	keys       map[ast.Node]string // corresponding keys for ASTs
	tagname    string
	fieldnames map[ast.Node]models.FieldName
}

func (tr *transformer) arrayType(v reflect.Value) ast.Expr {
	var m ast.Expr
	for i := 0; i < v.Len(); i++ {
		iv := v.Index(i)
		t := tr.transform(iv)
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
				m = &ast.StructType{Fields: combine(stM.Fields, stT.Fields)}
			}
		}
	}
	if m == nil {
		return &ast.ArrayType{Elt: ast.NewIdent("any")}
	}
	return &ast.ArrayType{Elt: m}
}

func (tr *transformer) structType(v reflect.Value) *ast.StructType {
	st := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{},
		},
	}
	iter := v.MapRange()
	for iter.Next() {
		ik := iter.Key()
		iv := iter.Value()
		fieldname := models.FieldName(namings.SafeFieldName(ik.String()))
		f := &ast.Field{
			Names: []*ast.Ident{fieldname.Ident()},
			Type:  tr.transform(iv),
			Tag: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("`%s:%q`", tr.tagname, ik.String()),
			},
		}
		st.Fields.List = append(st.Fields.List, f)
		tr.keys[f] = ik.String()
		tr.fieldnames[f] = fieldname
	}
	sort(st.Fields)
	return st
}

func (tr *transformer) stringType(v reflect.Value) ast.Expr {
	s := v.Interface().(string)
	if s == "0" { // BECAUSE: time.ParseDuration("0") doesn't return error
		return ast.NewIdent("string")
	}
	if _, err := time.ParseDuration(s); err == nil {
		tr.isTimeUsed = true
		return &ast.SelectorExpr{
			X:   ast.NewIdent("time"),
			Sel: ast.NewIdent("Duration"),
		}
	}
	return ast.NewIdent("string") // generic string
}

func (tr *transformer) transform(v reflect.Value) ast.Expr {
	if !v.IsValid() {
		fmt.Fprintf(os.Stderr, "Notice: Seen an invalid value (%q) and assigned 'any' as type. This may caused by input file contain a 'null' as value.\n", v.String())
		return ast.NewIdent("any")
	}
	t := v.Type()
	switch t.Kind() {
	case reflect.Interface:
		return tr.transform(v.Elem())
	case reflect.Map:
		return tr.structType(v)
	case reflect.Slice:
		return tr.arrayType(v)
	case reflect.Bool:
		return ast.NewIdent("bool")
	case reflect.String:
		return tr.stringType(v)
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

// reconstructs a reflect-value's type in ast.TypeSpec.
// limited with types used by YAML decoder.
func Transform(b *bundle.Bundle) {
	t := &transformer{
		isTimeUsed: false,
		keys:       map[ast.Node]string{},
		tagname:    string(b.Encoding),
		fieldnames: map[ast.Node]models.FieldName{},
	}
	b.CfgType = t.transform(reflect.ValueOf(b.Cfgcontent))
	if t.isTimeUsed {
		b.AddImports("time")
	}
	b.OriginalKeys = t.keys
	b.Fieldnames = t.fieldnames
}
