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

	OriginalKeys   map[ast.Node]string         // holder -> key
	Keypaths       map[ast.Node]models.Keypath // holder -> keypath
	TypeDefHolders map[models.Keypath]ast.Node // keypath -> Field, ArrayType

	Imports []string // package paths

	// type declarations
	CfgType ast.Expr // config type, needed to be placed in a TypeSpec

	// function declarations
	Isolated  *ast.GenDecl    // organization
	Iterators []*ast.FuncDecl // .Range() methods
	Named     []*ast.GenDecl  // mappings, directives
	Accessors []*ast.FuncDecl // directives
}

func New(cfgcontent any, encoding models.Encoding, typename string) *Bundle {
	cfg, imports, keys := transform.Transform(cfgcontent, encoding)
	imports = append(imports, "fmt", "os") // ReadConfig

	b := &Bundle{
		Encoding:        encoding,
		TypeName:        typename,
		TypeNameInitial: namings.Initial(typename),
		OriginalKeys:    keys,
		Keypaths:        map[ast.Node]models.Keypath{},
		TypeDefHolders:  nil,
		Imports:         imports,
		CfgType:         cfg,
		Isolated:        nil,
		Iterators:       nil,
		Named:           nil,
		Accessors:       nil,
	}

	if b.Encoding == models.Yaml {
		b.Imports = append(b.Imports, "gopkg.in/yaml.v3")
	} else {
		b.Imports = append(b.Imports, "encoding/json")
	}

	return b
}

