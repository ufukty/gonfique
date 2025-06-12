package holders

import (
	"fmt"
	"go/ast"
)

type Node struct {
	Holder      ast.Node
	Termination string
}

func (n *Node) Get() (ast.Expr, error) {
	switch h := n.Holder.(type) {
	case *ast.Field:
		return h.Type, nil
	case *ast.ArrayType:
		return h.Elt, nil
	case *ast.MapType:
		switch n.Termination {
		case "[key]":
			return h.Key, nil
		case "[value]":
			return h.Value, nil
		}
	case *ast.TypeSpec:
		return h.Type, nil
	}
	return nil, fmt.Errorf("unknown holder type (%T) or path termination (%s)", n.Holder, n.Termination)
}

func (n *Node) Set(expr ast.Expr) error {
	switch h := n.Holder.(type) {
	case *ast.Field:
		h.Type = expr
		return nil
	case *ast.ArrayType:
		h.Elt = expr
		return nil
	case *ast.MapType:
		switch n.Termination {
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
	return fmt.Errorf("unknown holder type (%T) or path termination (%s)", n.Holder, n.Termination)
}
