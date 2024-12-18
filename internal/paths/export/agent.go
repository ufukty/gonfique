package export

import (
	"go/ast"

	"github.com/ufukty/gonfique/internal/files/config"
)

type Agent struct {
	typenames []config.Typename
	Decls     []*ast.GenDecl
}

func New() *Agent {
	return &Agent{
		typenames: []config.Typename{},
		Decls:     []*ast.GenDecl{},
	}
}
