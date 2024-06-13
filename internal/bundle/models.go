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

	Keypaths map[ast.Node]models.FlattenKeypath // holder -> keypath (resolver)
	Holders  map[models.FlattenKeypath]ast.Node // keypath -> Field, ArrayType (inverse Keypaths)

	// either declared or a built-in basic types
	NeededToBeReferred []models.FlattenKeypath
	// those can't stay inline (if they are composite)
	NeededToBeDeclared []models.FlattenKeypath

	// final typenames to refer (user-provided > auto-generated > basic/built-in)
	ElectedTypenames map[models.FlattenKeypath]models.TypeName
	Usages           map[models.TypeName][]models.FlattenKeypath

	Expansions map[models.WildcardKeypath][]models.FlattenKeypath // matches
	TypeExprs  map[models.FlattenKeypath]ast.Expr

	Imports []string // package paths

	// type declarations
	Cfgcontent any      // produced by yaml.Decoder
	CfgType    ast.Expr // config type, needed to be placed in a TypeSpec

	// function declarations
	Isolated *ast.GenDecl   // organization
	Named    []*ast.GenDecl // mappings, directives

	Iterators []*ast.FuncDecl // .Range() methods
	Accessors []*ast.FuncDecl // directives

	Df *directivefile.DirectiveFile
}

func New(typename string) *Bundle {
	return &Bundle{
		TypeName:        typename,
		TypeNameInitial: namings.Initial(typename),

		NeededToBeReferred: []models.FlattenKeypath{},
		NeededToBeDeclared: []models.FlattenKeypath{},

		ElectedTypenames: map[models.FlattenKeypath]models.TypeName{},

		Expansions: map[models.WildcardKeypath][]models.FlattenKeypath{},
		TypeExprs:  map[models.FlattenKeypath]ast.Expr{},

		Imports: []string{},
	}
}

func (b *Bundle) AddImports(path ...string) {
	b.Imports = append(b.Imports, path...)
}
