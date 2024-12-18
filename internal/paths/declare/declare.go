package declare

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/declare/clone"
	"github.com/ufukty/gonfique/internal/paths/resolve"
)

func get(holder ast.Node, termination string) (ast.Expr, error) {
	switch h := holder.(type) {
	case *ast.Field:
		return h.Type, nil
	case *ast.ArrayType:
		return h.Elt, nil
	case *ast.MapType:
		switch termination {
		case "[key]":
			return h.Key, nil
		case "[value]":
			return h.Value, nil
		}
	}
	return nil, fmt.Errorf("unknown holder type (%T) or path termination (%s)", holder, termination)
}

func set(holder ast.Node, last string, expr ast.Expr) error {
	switch h := holder.(type) {
	case *ast.Field:
		h.Type = expr
		return nil
	case *ast.ArrayType:
		h.Elt = expr
		return nil
	case *ast.MapType:
		switch last {
		case "[key]":
			h.Key = expr
			return nil
		case "[value]":
			h.Value = expr
			return nil
		}
	}
	return fmt.Errorf("unkown holder type (%T) or path termination (%s)", holder, last)
}

func (a *Agent) Declare(holder ast.Node, last string, tn config.Typename, rp resolve.Path) (*ast.TypeSpec, error) {
	expr, err := get(holder, last)
	if err != nil {
		return nil, fmt.Errorf("checking existing type: %w", err)
	}

	err = set(holder, last, tn.Ident())
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
