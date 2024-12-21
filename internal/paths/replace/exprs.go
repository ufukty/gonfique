package replace

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"

	"github.com/ufukty/gonfique/internal/holders"
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

func (a *Agent) Expression(v string, holder ast.Node, last string) error {
	v2, err := parse(v)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	e, err := expr(v2.Typename)
	if err != nil {
		return fmt.Errorf("building ast for typename: %w", err)
	}
	holders.Set(holder, last, e)
	if v2.ImportPath != "" {
		a.Imports = append(a.Imports, v2.ImportPath)
	}
	return nil
}
