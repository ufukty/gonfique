package pkg

import (
	"fmt"
	"go/ast"
	"go/token"
)

func Iterators(cfg *ast.TypeSpec, isolated *ast.GenDecl) ([]*ast.FuncDecl, error) {
	fds := []*ast.FuncDecl{}
	for _, gd := range []*ast.GenDecl{{Tok: token.TYPE, Specs: []ast.Spec{cfg}}, isolated} {
		for _, s := range gd.Specs {
			if ts, ok := s.(*ast.TypeSpec); ok {
				if st, ok := ts.Type.(*ast.StructType); ok {
					var cti *ast.Ident
					for _, f := range st.Fields.List {
						if ti, ok := f.Type.(*ast.Ident); ok {
							if cti == nil {
								cti = ti
								continue
							} else if ti.Name == cti.Name {
								continue
							}
						}
						cti = nil
						break
					}
					// if the all fields have same Ident in their types;
					// generate a FuncDecl which its body consists by a ReturnStmt of map[string]cti
					// the map has the exact same amount of Fields struct type has
					if cti != nil {
						elements := []ast.Expr{}
						for _, f := range st.Fields.List {
							keyname, err := stripKeyname(f.Tag.Value)
							if err != nil {
								return nil, fmt.Errorf("could not strip the keyname in %s.%s field tag list: %w", ts.Name.Name, f.Names[0].Name, err)
							}
							elements = append(elements, &ast.KeyValueExpr{
								Key:   &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", keyname)},
								Value: &ast.SelectorExpr{X: &ast.Ident{Name: "a"}, Sel: f.Names[0]},
							})
						}
						fds = append(fds, &ast.FuncDecl{
							Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "a"}}, Type: ts.Name}}},
							Name: &ast.Ident{Name: "Range"},
							Type: &ast.FuncType{
								Params:  &ast.FieldList{},
								Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: cti}}}},
							},
							Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{
								Results: []ast.Expr{&ast.CompositeLit{
									Type: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: cti},
									Elts: elements,
								}},
							}}},
						})
					}
				}
			}
		}
	}
	return fds, nil
}
