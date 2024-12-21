package declare

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/holders"
	"github.com/ufukty/gonfique/internal/paths/declare/clone"
	"github.com/ufukty/gonfique/internal/paths/resolve"
)

func (a *Agent) Declare(holder ast.Node, last string, tn config.Typename, rp resolve.Path) (*ast.TypeSpec, error) {
	expr, err := holders.Get(holder, last)
	if err != nil {
		return nil, fmt.Errorf("checking existing type: %w", err)
	}

	err = holders.Set(holder, last, tn.Ident())
	if err != nil {
		return nil, fmt.Errorf("replacing type expression with declared type: %w", err)
	}

	a.exprs[rp] = expr

	_, declared := a.users[tn]
	if declared {
		a.users[tn] = append(a.users[tn], rp)
		return nil, nil

	} else {
		a.users[tn] = []resolve.Path{rp}
		ts := &ast.TypeSpec{Name: tn.Ident(), Type: clone.Expr(expr)}
		a.Decls = append(a.Decls, &ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{ts},
		})
		return ts, nil
	}
}
