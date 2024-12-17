package replace

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

type value struct {
	Typename   string
	ImportPath string
}

var spaces = regexp.MustCompile(`\s+`)

func parse(s string) (*value, error) {
	ss := strings.Split(spaces.ReplaceAllString(s, ";"), ";")
	if len(ss) > 2 {
		return nil, fmt.Errorf("more than 2 values")
	}
	v := &value{
		Typename: ss[0],
	}
	if len(ss) == 2 {
		v.ImportPath = ss[1]
	}
	return v, nil
}

func expr(s string) (ast.Expr, error) {
	ss := strings.Split(s, ".")
	switch len(ss) {
	case 1:
		return ast.NewIdent(s), nil
	case 2:
		return &ast.SelectorExpr{X: ast.NewIdent(ss[0]), Sel: ast.NewIdent(ss[1])}, nil
	default:
		return nil, fmt.Errorf("too many dots")
	}
}

func set(holder ast.Node, last string, e ast.Expr) {
	switch h := holder.(type) {
	case *ast.Field:
		h.Type = e
	case *ast.ArrayType:
		h.Elt = e
	case *ast.MapType:
		switch last {
		case "[key]":
			h.Key = e
		case "[value]":
			h.Value = e
		}
	}
}

func Expression(v string, holder ast.Node, last string) (string, error) {
	v2, err := parse(v)
	if err != nil {
		return "", fmt.Errorf("parse: %w", err)
	}
	e, err := expr(v2.Typename)
	if err != nil {
		return "", fmt.Errorf("building ast for typename: %w", err)
	}
	set(holder, last, e)
	if v2.ImportPath != "" {
		return v2.ImportPath, nil
	}
	return "", nil
}
