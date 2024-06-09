package coder

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"slices"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
)

func addImports(dst *ast.File, imports []string) {
	imports = append(imports, "fmt", "os") // ReadConfig
	specs := []ast.Spec{}
	slices.Sort(imports)
	imports = datas.Uniq(imports)
	for _, imp := range imports {
		specs = append(specs, &ast.ImportSpec{
			Name:    nil,
			Comment: nil,
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprintf("\"%s\"", imp),
			},
		})
	}
	dst.Decls = append(dst.Decls, &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: specs,
	})
}

func addIsolatedTypeSpecifications(dst *ast.File, isolated *ast.GenDecl) {
	if isolated != nil {
		dst.Decls = append(dst.Decls, isolated)
	}
}

func addIteratorMethods(dst *ast.File, iterators []*ast.FuncDecl) {
	for _, fd := range iterators {
		dst.Decls = append(dst.Decls, fd)
	}
}

func addNamedTypeSpecifications(dst *ast.File, named []*ast.GenDecl) {
	slices.SortFunc(named, func(a, b *ast.GenDecl) int {
		if a.Specs[0].(*ast.TypeSpec).Name.Name > b.Specs[0].(*ast.TypeSpec).Name.Name {
			return 1
		} else {
			return -1
		}
	})
	for _, n := range named {
		dst.Decls = append(dst.Decls, n)
	}
}

func addConfig(dst *ast.File, cfg ast.Expr, typeName string) {
	dst.Decls = append(dst.Decls, &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{
			Name: ast.NewIdent(typeName),
			Type: cfg,
		}},
	})
}

func addAccessors(dst *ast.File, accessors []*ast.FuncDecl) {
	for _, fd := range accessors {
		dst.Decls = append(dst.Decls, fd)
	}
}

func Write(b *bundle.Bundle, dst, pkgname string) error {
	af := &ast.File{
		Name:  ast.NewIdent(pkgname),
		Decls: []ast.Decl{},
	}

	addImports(af, b.Imports)
	addIsolatedTypeSpecifications(af, b.Isolated)
	addIteratorMethods(af, b.Iterators)
	addNamedTypeSpecifications(af, b.Named)
	addConfig(af, b.CfgType, b.TypeName)
	addReaderFunction(b, af)
	if b.Accessors != nil {
		addAccessors(af, b.Accessors)
	}

	o, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer o.Close()

	err = format.Node(o, token.NewFileSet(), af)
	if err != nil {
		return fmt.Errorf("writing into output file: %w", err)
	}

	return nil
}
