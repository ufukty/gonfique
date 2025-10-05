package declare

import (
	"fmt"
	"go/ast"
	"go/token"

	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/holders"
	"go.ufukty.com/gonfique/internal/paths/declare/clone"
	"go.ufukty.com/gonfique/internal/paths/mapper/resolve"
)

func (a *Agent) Declare(h holders.Node, tn config.Typename, rp resolve.Path) (*ast.TypeSpec, error) {
	expr, err := h.Get()
	if err != nil {
		return nil, fmt.Errorf("checking existing type: %w", err)
	}

	err = h.Set(tn.Ident())
	if err != nil {
		return nil, fmt.Errorf("replacing type expression with declared type: %w", err)
	}

	a.exprs[rp] = clone.Expr(expr)

	_, declared := a.users[tn]
	if declared {
		a.users[tn] = append(a.users[tn], rp)
		return nil, nil

	} else {
		a.users[tn] = []resolve.Path{rp}
		ts := &ast.TypeSpec{Name: tn.Ident(), Type: expr}
		a.Decls[tn] = &ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{ts},
		}
		return ts, nil
	}
}
