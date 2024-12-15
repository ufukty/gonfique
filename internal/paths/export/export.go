package export

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/namings/auto"
	"github.com/ufukty/gonfique/internal/paths/resolve"
)

func get(holder ast.Node) (ast.Expr, error) {
	switch h := holder.(type) {
	case *ast.Field:
		return h.Type, nil
	case *ast.ArrayType:
		return h.Elt, nil
	default:
		return nil, fmt.Errorf("unknown holder type: %T", holder)
	}
}

func set(holder ast.Node, expr ast.Expr) error {
	switch h := holder.(type) {
	case *ast.Field:
		h.Type = expr
		return nil
	case *ast.ArrayType:
		h.Elt = expr
		return nil
	default:
		return fmt.Errorf("unknown holder type: %T", holder)
	}
}

func Types(targets []resolve.Path, reserved []config.Typename, holders map[resolve.Path]ast.Node) ([]*ast.GenDecl, error) {
	decls := []*ast.GenDecl{}

	types := auto.GenerateTypenames(targets, reserved)
	for rp, tn := range types {
		expr, err := get(holders[rp])
		if err != nil {
			return nil, fmt.Errorf("getting type expression of target: %w", err)
		}
		decls = append(decls, &ast.GenDecl{
			Doc:   &ast.CommentGroup{List: []*ast.Comment{{Text: fmt.Sprintf("// exported for %s", rp)}}},
			Tok:   token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{Name: tn.Ident(), Type: expr}},
		})
		err = set(holders[rp], tn.Ident())
		if err != nil {
			return nil, fmt.Errorf("replacing type def with typename on target: %w", err)
		}
	}

	return decls, nil
}
