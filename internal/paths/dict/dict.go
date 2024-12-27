package dict

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/holders"
	"github.com/ufukty/gonfique/internal/paths/dict/combine"
	"github.com/ufukty/gonfique/internal/transform"
)

func ConvertToMap(h holders.Node, ti *transform.Info) (*ast.MapType, error) {
	t, err := h.Get()
	if err != nil {
		return nil, fmt.Errorf("getting type expression: %w", err)
	}

	st, ok := t.(*ast.StructType)
	if !ok {
		return nil, fmt.Errorf("target is not a struct (%T)", t)
	}

	types := []ast.Expr{}
	if st.Fields != nil && st.Fields.List != nil {
		for _, f := range st.Fields.List {
			if f.Type != nil {
				types = append(types, f.Type)
			}
		}
	}
	if len(types) == 0 {
		h.Set(ast.NewIdent("any"))
	}
	mv, err := combine.Combine(ti, types...)
	if err != nil {
		return nil, fmt.Errorf("combining field types: %w", err)
	}

	mt := &ast.MapType{
		Key:   ast.NewIdent("string"),
		Value: mv,
	}

	err = h.Set(mt)
	if err != nil {
		return nil, fmt.Errorf("assigning generated map type: %w", err)
	}

	return mt, nil
}
