package bundle

import (
	"go/ast"

	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
)

type Bundle struct {
	Encoding models.Encoding

	TypeName        string
	TypeNameInitial string

	OriginalKeys map[ast.Node]string         // holder -> key
	Keypaths     map[ast.Node]models.Keypath // holder -> keypath
	Holders      map[models.Keypath]ast.Node // keypath -> Field, ArrayType

	Imports []string // package paths

	// type declarations
	CfgType ast.Expr // config type, needed to be placed in a TypeSpec

	// function declarations
	Isolated *ast.GenDecl   // organization
	Named    []*ast.GenDecl // mappings, directives

	Iterators []*ast.FuncDecl // .Range() methods
	Accessors []*ast.FuncDecl // directives
}

func New(cfgcontent any, encoding models.Encoding, typename string) *Bundle {

	b := &Bundle{
		Encoding:        encoding,
		TypeName:        typename,
		TypeNameInitial: namings.Initial(typename),
	}

	return b
}
