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

	OriginalKeys map[ast.Node]string                       // holder -> key
	Keypaths     map[ast.Node]models.FlattenKeypath        // holder -> keypath (resolver)
	Holders      map[models.WildcardKeypath]ast.Node       // keypath -> Field, ArrayType (inverse Keypaths)
	Typenames    map[models.FlattenKeypath]models.TypeName // provided by `namings`
	Expansions   map[models.WildcardKeypath][]ast.Node     // keypath (wildcards) -> []match

	Imports []string // package paths

	// type declarations
	Cfgcontent any      // produced by yaml.Decoder
	CfgType    ast.Expr // config type, needed to be placed in a TypeSpec

	// function declarations
	Isolated *ast.GenDecl   // organization
	Named    []*ast.GenDecl // mappings, directives

	Iterators []*ast.FuncDecl // .Range() methods
	Accessors []*ast.FuncDecl // directives

	Df             *directivefile.DirectiveFile
	NeedsToBeNamed []models.FlattenKeypath // filled by directives.preprocess
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
