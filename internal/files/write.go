package files

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"slices"

	"github.com/ufukty/gonfique/internal/transform"
)

func addImports(dst *ast.File, imports []string) {
	specs := []ast.Spec{}
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

func addConfig(dst *ast.File, cfg ast.Expr) {
	dst.Decls = append(dst.Decls, &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{
			Name: ast.NewIdent("Config"),
			Type: cfg,
		}},
	})
}

func (f *File) addReaderFunction(dst *ast.File) {
	if f.Encoding == transform.Yaml {
		dst.Decls = append(dst.Decls, readerFuncForYaml)
	} else if f.Encoding == transform.Json {
		dst.Decls = append(dst.Decls, readerFuncForJson)
	}
}

func (f *File) Write(dst string, pkgname string) error {
	af := &ast.File{
		Name:  ast.NewIdent(pkgname),
		Decls: []ast.Decl{},
	}

	addImports(af, f.Imports)
	addIsolatedTypeSpecifications(af, f.Isolated)
	addIteratorMethods(af, f.Iterators)
	addNamedTypeSpecifications(af, f.Named)
	addConfig(af, f.Cfg)
	f.addReaderFunction(af)

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
