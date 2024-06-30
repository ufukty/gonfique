package iterables

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/namings"
)

// returns nil if all field types are not same
func getCommonTypeOfFields(st *ast.StructType) *ast.Ident {
	if st.Fields == nil || st.Fields.List == nil {
		return nil
	}
	var ct *ast.Ident
	for _, f := range st.Fields.List {
		if f.Type == nil {
			return nil
		}
		if t, ok := f.Type.(*ast.Ident); ok {
			if ct == nil {
				ct = t
			} else if t.Name != ct.Name {
				return nil
			}
		} else {
			return nil
		}
	}
	return ct
}

func generateIterator(ts *ast.TypeSpec, commonType *ast.Ident, originalKeys map[ast.Node]string) *ast.FuncDecl {
	typeSpecNameInitial := ast.NewIdent(namings.Initial(ts.Name.Name))
	keyValuePairs := []ast.Expr{}
	for _, f := range ts.Type.(*ast.StructType).Fields.List {
		keyValuePairs = append(keyValuePairs, &ast.KeyValueExpr{
			Key:   &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", originalKeys[f])},
			Value: &ast.SelectorExpr{X: typeSpecNameInitial, Sel: f.Names[0]},
		})
	}
	return &ast.FuncDecl{
		Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{typeSpecNameInitial}, Type: ts.Name}}},
		Name: &ast.Ident{Name: "Range"},
		Type: &ast.FuncType{
			Params:  &ast.FieldList{},
			Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: commonType}}}},
		},
		Body: &ast.BlockStmt{List: []ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{&ast.CompositeLit{
					Type: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: commonType},
					Elts: keyValuePairs,
				}},
			},
		}},
	}
}

func ImplementIterators(b *bundle.Bundle) error {
	fds := []*ast.FuncDecl{}
	gds := []*ast.GenDecl{}
	if b.Isolated != nil {
		gds = append(gds, b.Isolated)
	}
	if b.Named != nil {
		gds = append(gds, b.Named...)
	}
	gds = append(gds, &ast.GenDecl{ // temporary
		Tok:   token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{Name: ast.NewIdent(b.TypeName), Type: b.CfgType}},
	})
	for _, gd := range gds {
		for _, s := range gd.Specs {
			if ts, ok := s.(*ast.TypeSpec); ok && ts.Type != nil {
				if st, ok := ts.Type.(*ast.StructType); ok {
					// if the all fields have same Ident in their types;
					// generate a FuncDecl which its body consists by a ReturnStmt of map[string]ct
					// the map has the exact same amount of Fields struct type has
					if ct := getCommonTypeOfFields(st); ct != nil {
						fds = append(fds, generateIterator(ts, ct, b.OriginalKeys))
					}
				}
			}
		}
	}
	b.Iterators = fds
	return nil
}
