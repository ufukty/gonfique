package pkg

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

func rec(v reflect.Value) ast.Expr {
	t := v.Type()
	switch t.Kind() {
	case reflect.Interface:
		return rec(v.Elem())
	case reflect.Map:
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
				Names: []*ast.Ident{ast.NewIdent(ik.String())},
				Type:  rec(iv),
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf("`yaml:%q`", ik.String()),
				},
			})
		}
		return st
	case reflect.Slice:
		return ast.NewIdent("any")
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

func GetTypeSpecForConfig(src string) (*ast.TypeSpec, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer f.Close()
	var y any
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decoding input file: %w", err)
	}
	return &ast.TypeSpec{
		Name: ast.NewIdent("Config"),
		Type: rec(reflect.ValueOf(y)),
	}, nil
}

func WriteConfigTypeSpecIntoFile(dst string, cfg *ast.TypeSpec, pkgname string) error {
	f := &ast.File{
		Name: ast.NewIdent(pkgname),
		Decls: []ast.Decl{
			imports,
			&ast.GenDecl{
				Tok:   token.TYPE,
				Specs: []ast.Spec{cfg},
			},
			readerFunc,
		},
	}

	o, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer o.Close()

	err = format.Node(o, token.NewFileSet(), f)
	if err != nil {
		return fmt.Errorf("writing into output file: %w", err)
	}

	return nil
}
