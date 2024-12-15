package coder

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/printer"
	"go/token"
	"os"
	"regexp"
	"slices"

	"github.com/ufukty/gonfique/cmd/gonfique/commands/version"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/files/coder/sort"
	"github.com/ufukty/gonfique/internal/files/config/meta"
	"github.com/ufukty/gonfique/internal/files/input"
	"github.com/ufukty/gonfique/internal/namings"
)

type Coder struct {
	ti *ast.Ident

	Meta     meta.Meta
	Encoding input.Encoding

	Config ast.Expr

	Imports     []string
	Named, Auto []*ast.GenDecl

	Accessors, Iterators []*ast.FuncDecl
	ParentRefAssignments []ast.Stmt
}

func quotes(s string) string {
	return fmt.Sprintf("%q", s)
}

func (c Coder) addImports(dst *ast.File) {
	imports := slices.Clone(c.Imports)

	imports = append(imports, "fmt", "os") // ReadConfig
	switch c.Encoding {
	case input.Yaml:
		imports = append(imports, "gopkg.in/yaml.v3")
	case input.Json:
		imports = append(imports, "encoding/json")
	}

	slices.Sort(imports)
	imports = datas.Uniq(imports)

	specs := []ast.Spec{}
	for _, imp := range imports {
		specs = append(specs, &ast.ImportSpec{
			Path: &ast.BasicLit{Kind: token.STRING, Value: quotes(imp)},
		})
	}
	sort.Imports(specs)

	dst.Decls = append(dst.Decls, &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: specs,
	})
}

func (c Coder) addIteratorMethods(dst *ast.File) {
	if c.Iterators == nil {
		return
	}
	for _, fd := range c.Iterators {
		dst.Decls = append(dst.Decls, fd)
	}
}

func (c Coder) addAutoTypes(dst *ast.File) {
	if c.Auto == nil {
		return
	}
	sort.FuncDecls(c.Auto)
	for _, n := range c.Auto {
		dst.Decls = append(dst.Decls, n)
	}
}

func (c Coder) addNamedTypes(dst *ast.File) {
	if c.Named == nil {
		return
	}
	sort.FuncDecls(c.Named)
	for _, n := range c.Named {
		dst.Decls = append(dst.Decls, n)
	}
}

func (c Coder) addConfig(dst *ast.File) {
	dst.Decls = append(dst.Decls, &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{
			Name: ast.NewIdent(c.Meta.Type),
			Type: c.Config,
		}},
	})
}

func (c Coder) addAccessors(dst *ast.File) {
	if len(c.Accessors) == 0 {
		return
	}
	sort.Accessors(c.Accessors)
	for _, fd := range c.Accessors {
		dst.Decls = append(dst.Decls, fd)
	}
}

var typedecls = regexp.MustCompile(`(?m)$(\n(?://.*\n)*type)`)

func post(s string) string {
	s = fmt.Sprintf("// Code generated by gonfique %s. DO NOT EDIT.\n\n%s", version.Version, s)
	s = typedecls.ReplaceAllString(s, "\n$1")
	return s
}

func (c Coder) Write(dst string) error {
	c.ti = ast.NewIdent(namings.Initial(c.Meta.Type))

	f := &ast.File{
		Name:  ast.NewIdent(c.Meta.Package),
		Decls: []ast.Decl{},
	}

	c.addImports(f)
	c.addIteratorMethods(f)
	c.addAutoTypes(f)
	c.addNamedTypes(f)
	c.addConfig(f)
	c.addParentRefAssignmentsFunction(f)
	if err := c.addReaderFunction(f); err != nil {
		return fmt.Errorf("reader: %w", err)
	}
	c.addAccessors(f)

	b := bytes.NewBufferString("")
	err := printer.Fprint(b, token.NewFileSet(), f)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	fs, err := format.Source([]byte(post(b.String())))
	if err != nil {
		return fmt.Errorf("format: %w", err)
	}

	o, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer o.Close()
	fmt.Fprintf(o, "%s", fs)

	return nil
}
