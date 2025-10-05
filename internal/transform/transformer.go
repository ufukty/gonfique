package transform

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"reflect"

	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/files/input/encoders"
	"go.ufukty.com/gonfique/internal/transform/combine"
	"go.ufukty.com/gonfique/internal/transform/safe"
)

type Info struct {
	Type       ast.Expr
	Keys       map[ast.Node]string
	Fieldnames map[ast.Node]config.Fieldname
}

type transformer struct {
	keys       map[ast.Node]string // corresponding keys for ASTs
	fieldnames map[ast.Node]config.Fieldname
	tagname    string
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
				m = &ast.StructType{Fields: combine.FieldLists(stM.Fields, stT.Fields)}
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
		fieldname := safe.FieldName(ik.String())
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

// reconstructs a reflect-value's type in ast.TypeSpec.
// limited with types used by YAML decoder.
func Transform(d any, encoding encoders.Encoding) *Info {
	tr := transformer{
		keys:       map[ast.Node]string{},
		fieldnames: map[ast.Node]config.Fieldname{},
		tagname:    string(encoding),
	}
	ty := tr.transform(reflect.ValueOf(d))
	return &Info{
		Type:       ty,
		Keys:       tr.keys,
		Fieldnames: tr.fieldnames,
	}
}
