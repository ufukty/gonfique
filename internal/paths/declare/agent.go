package declare

import (
	"go/ast"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"golang.org/x/exp/maps"
)

// keeps track of type expression and typename conflicts.
// provides troubleshoot information back to user.
type Agent struct {
	Decls []*ast.GenDecl
	exprs map[resolve.Path]ast.Expr
	users map[config.Typename][]resolve.Path
}

func New() *Agent {
	return &Agent{
		Decls: []*ast.GenDecl{},
		exprs: map[resolve.Path]ast.Expr{},
		users: map[config.Typename][]resolve.Path{},
	}
}

func (a *Agent) Typenames() []config.Typename {
	return maps.Keys(a.users)
}
