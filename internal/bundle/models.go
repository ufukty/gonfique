package bundle

import (
	"go/ast"

	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
)

type Bundle struct {
	Encoding models.Encoding

	TypeName        string
	TypeNameInitial string

	OriginalKeys map[ast.Node]string           // holder -> key
	Fieldnames   map[ast.Node]models.FieldName // populated by transformer

	// type declarations
	Cfgcontent any      // produced by yaml.Decoder
	CfgType    ast.Expr // config type, needed to be placed in a TypeSpec

	// function declarations
	Isolated *ast.GenDecl   // organization
	Named    []*ast.GenDecl // mappings, directives
	Imports  []string       // package paths

	Iterators []*ast.FuncDecl // .Range() methods
	Accessors []*ast.FuncDecl // directives

	Df                   *directivefile.DirectiveFile
	ParentRefAssignStmts []ast.Stmt
}

func New(typename string) *Bundle {
	return &Bundle{
		TypeName:        typename,
		TypeNameInitial: namings.Initial(typename),
		Imports:         []string{},
	}
}

func (b *Bundle) AddImports(path ...string) {
	b.Imports = append(b.Imports, path...)
}
