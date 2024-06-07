package bundle

import (
	"go/ast"

	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
	"github.com/ufukty/gonfique/internal/transform"
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

	b.CfgType, b.Imports, b.OriginalKeys = transform.Transform(cfgcontent, encoding)
	b.Imports = append(b.Imports, "fmt", "os") // ReadConfig
	if b.Encoding == models.Yaml {
		b.Imports = append(b.Imports, "gopkg.in/yaml.v3")
	} else {
		b.Imports = append(b.Imports, "encoding/json")
	}

	return b
}
