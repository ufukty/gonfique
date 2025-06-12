package iterator

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/transform"
)

type Agent struct {
	Decls map[config.Typename]*ast.FuncDecl
}

func New() *Agent {
	return &Agent{
		Decls: map[config.Typename]*ast.FuncDecl{},
	}
}

// returns nil if all field types are not same
func fieldtype(fls []*ast.Field) (*ast.Ident, error) {
	var common *ast.Ident
	for _, f := range fls {
		if f.Type == nil {
			return nil, fmt.Errorf("field with uninitialized type")
		}
		t, ok := f.Type.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf("field with non-Ident type")
		}
		if common == nil {
			common = t
		} else if t.Name != common.Name {
			return nil, fmt.Errorf("fields with conflicting Ident types")
		}
	}
	return common, nil
}

func (a *Agent) Implement(ti *transform.Info, tn config.Typename, decl *ast.GenDecl) error {
	if decl == nil || len(decl.Specs) == 0 {
		return fmt.Errorf("specs not found")
	}
	ts, ok := decl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return fmt.Errorf("expected TypeSpec, got %T", decl.Specs[0])
	}
	if ts.Type == nil {
		return fmt.Errorf("type expression not found")
	}
	st, ok := ts.Type.(*ast.StructType)
	if !ok {
		return fmt.Errorf("expected StructType, got %T", ts.Type)
	}
	if st.Fields == nil || st.Fields.List == nil {
		return fmt.Errorf("fields or field list is uninitialized")
	}
	// if the all fields have same Ident in their types;
	// generate a FuncDecl which its body consists by a ReturnStmt of map[string]ct
	// the map has the exact same amount of Fields struct type has
	ct, err := fieldtype(st.Fields.List)
	if err != nil {
		return fmt.Errorf("comparing field types: %w", err)
	}
	a.Decls[tn] = fd(tn, st.Fields.List, ct, ti.Keys)
	return nil
}
