package pkg

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

func ReadConfigYaml(src string) (*ast.TypeSpec, error) {
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
		Type: Transform(reflect.ValueOf(y)),
	}, nil
}

func WriteConfigGo(dst string, cfg *ast.TypeSpec, pkgname string) error {
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
