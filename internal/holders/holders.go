package holders

import (
	"fmt"
	"go/ast"
)

func Get(holder ast.Node, termination string) (ast.Expr, error) {
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
	case *ast.TypeSpec:
		return h.Type, nil
	}
	return nil, fmt.Errorf("unknown holder type (%T) or path termination (%s)", holder, termination)
}

func Set(holder ast.Node, termination string, expr ast.Expr) error {
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
	case *ast.TypeSpec:
		h.Type = expr
		return nil
	}
	return fmt.Errorf("unknown holder type (%T) or path termination (%s)", holder, termination)
}
