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
	"github.com/ufukty/gonfique/internal/files/coder/sort"
	"github.com/ufukty/gonfique/internal/files/config/meta"
	"github.com/ufukty/gonfique/internal/files/input"
	"github.com/ufukty/gonfique/internal/namings"
	"golang.org/x/exp/maps"
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

func uniq[K comparable](ss []K) []K {
	m := make(map[K]any, len(ss))
	for _, s := range ss {
		m[s] = nil
	}
	return maps.Keys(m)
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
	imports = uniq(imports)

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

func (c Coder) createGenDecls(dst *ast.File) {
	decls := []ast.Decl{}
	if c.Iterators != nil {
		for _, fd := range c.Iterators {
			decls = append(decls, fd)
		}
	}
	if c.Auto != nil {
		for _, n := range c.Auto {
			decls = append(decls, n)
		}
	}
	if c.Named != nil {
		for _, n := range c.Named {
			decls = append(decls, n)
		}
	}
	if c.Accessors != nil {
		for _, fd := range c.Accessors {
			decls = append(decls, fd)
		}
	}
	if c.Config != nil {
		decls = append(decls, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{
				Name: ast.NewIdent(c.Meta.Type),
				Type: c.Config,
			}},
		})
	}
	dst.Decls = append(dst.Decls, sort.Decls(decls)...)
}

var typedecls = regexp.MustCompile(`(?m)$(\n(?://.*\n)*type)`)
var imports = regexp.MustCompile(`(?m)import \(((?:\s+"[\w]+(?:/[\w.]+)*"\n)*)((?:\s+"[\w.]+(?:/[\w.]+)*"\n)*)\)`)

func post(s string) string {
	s = fmt.Sprintf("// Code generated by gonfique %s. DO NOT EDIT.\n\n%s", version.Version, s)
	s = typedecls.ReplaceAllString(s, "\n$1")
	s = imports.ReplaceAllString(s, "import ($1\n$2)") // split packages starts with domains
	return s
}

func (c Coder) Write(dst string) error {
	c.ti = ast.NewIdent(namings.Initial(c.Meta.Type))

	f := &ast.File{
		Name:  ast.NewIdent(c.Meta.Package),
		Decls: []ast.Decl{},
	}
	c.addImports(f)
	c.createGenDecls(f)

	c.addParentRefAssignmentsFunction(f)
	if err := c.addReaderFunction(f); err != nil {
		return fmt.Errorf("reader: %w", err)
	}

	b := bytes.NewBufferString("")
	err := printer.Fprint(b, token.NewFileSet(), f)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	p := post(b.String())
	fs, err := format.Source([]byte(p))
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
