package files

import (
	"go/ast"

	"github.com/ufukty/gonfique/internal/transform"
)

type File struct {
	Encoding transform.Encoding

	TypeName        string
	TypeNameInitial string

	Keys           map[ast.Node]string // mappings
	TypeDefHolders map[string]ast.Node // keypath -> Field, ArrayType
	TypeDefs       map[string]ast.Node // keypath -> StructType, ArrayType, Ident

	Imports []string // package paths
	Cfg     ast.Expr // config type, needed to be placed in a TypeSpec

	Isolated  *ast.GenDecl    // organization
	Iterators []*ast.FuncDecl // .Range() methods
	Named     []*ast.GenDecl  // mappings, directives
	Accessors []*ast.FuncDecl // directives
}
