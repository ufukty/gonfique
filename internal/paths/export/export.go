package export

import (
	"fmt"
	"go/ast"
	"go/token"
	"slices"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/export/auto"
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

func set(holder ast.Node, termination string, expr ast.Expr) error {
	switch h := holder.(type) {
	case *ast.Field:
		h.Type = expr
		return nil
	case *ast.ArrayType:
		h.Elt = expr
		return nil
	case *ast.MapType:
		switch termination {
		case "[key]":
			h.Key = expr
			return nil
		case "[value]":
			h.Value = expr
			return nil
		}
	}
	return fmt.Errorf("unknown holder type (%T) or path termination (%s)", holder, termination)
}

func (a *Agent) Type(rp resolve.Path, reserved []config.Typename, holder ast.Node, termination string) error {
	tn, ok := auto.Typename(rp, slices.Concat(reserved, a.typenames))
	if !ok {
		return fmt.Errorf("could not produce typename for %s", rp)
	}

	expr, err := get(holder, termination)
	if err != nil {
		return fmt.Errorf("getting type expression of target: %w", err)
	}
	err = set(holder, termination, tn.Ident())
	if err != nil {
		return fmt.Errorf("replacing type def with typename on target: %w", err)
	}
	gd := &ast.GenDecl{
		Doc:   &ast.CommentGroup{List: []*ast.Comment{{Text: fmt.Sprintf("// exported for %s", rp)}}},
		Tok:   token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{Name: tn.Ident(), Type: expr}},
	}

	a.Decls = append(a.Decls, gd)
	a.typenames = append(a.typenames, tn)
	return nil
}
