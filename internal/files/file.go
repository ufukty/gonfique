package files

import "go/ast"

type EncodingLang int

const (
	Yaml = EncodingLang(iota)
	Json
)

type File struct {
	Lang          EncodingLang
	ConfigContent any
	Cfg           ast.Expr        // config type, needed to be placed in a TypeSpec
	Named         []*ast.GenDecl  // Named types
	Isolated      *ast.GenDecl    // Product of organization process
	Iterators     []*ast.FuncDecl // .Range() methods
	Imports       []string        // package paths
}
