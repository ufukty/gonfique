package replace

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"

	"github.com/ufukty/gonfique/internal/paths/resolve"
)

type value struct {
	Typename   string
	ImportPath string
}

var spaces = regexp.MustCompile(" +")

func (v *value) from(s string) error {
	ss := strings.Split(spaces.ReplaceAllString(s, ";"), ";")
	if len(ss) > 2 {
		return fmt.Errorf("more than 2 values")
	}
	v.Typename = ss[0]
	if len(ss) == 2 {
		v.ImportPath = ss[1]
	}
	return nil
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

func Expressions(directives map[resolve.Path]string, holders map[resolve.Path]ast.Node) ([]string, error) {
	v2 := value{}
	imports := []string{}
	for rp, v1 := range directives {
		err := v2.from(v1)
		if err != nil {
			return nil, fmt.Errorf("parsing: %w", err)
		}
		e, err := expr(v2.Typename)
		if err != nil {
			return nil, fmt.Errorf("building ast for typename: %w", err)
		}
		switch h := holders[rp].(type) {
		case *ast.Field:
			h.Type = e
		case *ast.ArrayType:
			h.Elt = e
		}
		if v2.ImportPath != "" {
			imports = append(imports, v2.ImportPath)
		}
	}
	return imports, nil
}
