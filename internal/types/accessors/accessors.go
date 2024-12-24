package accessors

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/transform"
)

type Agent struct {
	Decls map[config.Typename][]*ast.FuncDecl
}

func New() *Agent {
	return &Agent{
		Decls: map[config.Typename][]*ast.FuncDecl{},
	}
}

func (a *Agent) Implement(ti *transform.Info, tn config.Typename, t *ast.GenDecl, fields []config.Fieldname) error {
	if len(t.Specs) == 0 {
		return fmt.Errorf("specs are empty")
	}
	ts, ok := t.Specs[0].(*ast.TypeSpec)
	if !ok {
		return fmt.Errorf("expected type spec in generic declaration")
	}
	if ts.Type == nil {
		return fmt.Errorf("type expression is empty")
	}
	st, ok := ts.Type.(*ast.StructType)
	if !ok {
		return fmt.Errorf("expected a StructType, got %T", ts.Type)
	}
	if st.Fields == nil || st.Fields.List == nil {
		return fmt.Errorf("fields or field list is uninitialized")
	}

	types := map[config.Fieldname]ast.Expr{}
	for _, f := range st.Fields.List {
		switch f.Type.(type) {
		case *ast.Ident, *ast.StarExpr, *ast.SelectorExpr:
		default:
			return fmt.Errorf("found type of an accessor requested field is defined inline: %q", ti.Keys[f]) // TODO: consider allowing inline type definitions
		}
		types[ti.Fieldnames[f]] = f.Type
	}

	a.Decls[tn] = []*ast.FuncDecl{}
	for fn, ft := range types {
		a.Decls[tn] = append(a.Decls[tn], get(tn, fn, ft), set(tn, fn, ft))
	}

	return nil
}
