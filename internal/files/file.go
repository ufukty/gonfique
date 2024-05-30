package files

import (
	"go/ast"

	"github.com/ufukty/gonfique/internal/transform"
)

type File struct {
	Encoding      transform.Encoding
	TypeName      string
	Keys          map[ast.Node]string
	ConfigContent any
	Cfg           ast.Expr        // config type, needed to be placed in a TypeSpec
	Named         []*ast.GenDecl  // Named types
	Isolated      *ast.GenDecl    // Product of organization process
	Iterators     []*ast.FuncDecl // .Range() methods
	Imports       []string        // package paths
}
