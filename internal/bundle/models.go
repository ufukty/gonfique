package bundle

import (
	"go/ast"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/transform"
)

type Bundle struct {
	Encoding models.Encoding

	TypeName        string
	TypeNameInitial string

	OriginalKeys   map[ast.Node]string // mappings
	TypeDefHolders map[string]ast.Node // keypath -> Field, ArrayType
	TypeDefs       map[string]ast.Node // keypath -> StructType, ArrayType, Ident

	Imports []string // package paths
	Cfg     ast.Expr // config type, needed to be placed in a TypeSpec

	Isolated  *ast.GenDecl    // organization
	Iterators []*ast.FuncDecl // .Range() methods
	Named     []*ast.GenDecl  // mappings, directives
	Accessors []*ast.FuncDecl // directives
}

func initial(name string) string {
	return strings.ToLower(string(([]rune(name))[0]))
}

func New(cfgcontent any, encoding models.Encoding, typename string) *Bundle {
	if encoding == models.Yaml {
		cfg, imports, keys := transform.Transform(cfgcontent, encoding)
		imports = slices.Concat([]string{"fmt", "os", "gopkg.in/yaml.v3"}, imports)
		slices.Sort(imports)

		return &Bundle{
			Encoding:        models.Yaml,
			OriginalKeys:    keys,
			TypeName:        typename,
			TypeNameInitial: initial(typename),
			Cfg:             cfg,
			Named:           nil,
			Isolated:        nil,
			Iterators:       nil,
			Imports:         imports,
		}

	} else {
		cfg, imports, keys := transform.Transform(cfgcontent, encoding)
		imports = slices.Concat([]string{"fmt", "os", "encoding/json"}, imports)
		slices.Sort(imports)

		return &Bundle{
			Encoding:        models.Json,
			OriginalKeys:    keys,
			TypeName:        typename,
			TypeNameInitial: initial(typename),
			Cfg:             cfg,
			Named:           nil,
			Isolated:        nil,
			Iterators:       nil,
			Imports:         imports,
		}
	}
}
