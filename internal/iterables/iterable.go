package iterables

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/files"
)

func DetectIterators(file *files.File) error {
	fds := []*ast.FuncDecl{}
	gds := []*ast.GenDecl{}
	if file.Isolated != nil {
		gds = append(gds, file.Isolated)
	}
	gds = append(gds, &ast.GenDecl{ // temporary
		Tok:   token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{Name: ast.NewIdent(file.TypeName), Type: file.Cfg}},
	})
	receivername := ast.NewIdent(file.TypeNameInitial)
	for _, gd := range gds {
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
							keyname, ok := file.OriginalKeys[f]
							if !ok {
								return fmt.Errorf("could not retrieve the original keyname for %s.%s (AST %p)", ts.Name.Name, f.Names[0].Name, f)
							}
							elements = append(elements, &ast.KeyValueExpr{
								Key:   &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", keyname)},
								Value: &ast.SelectorExpr{X: receivername, Sel: f.Names[0]},
							})
						}
						fds = append(fds, &ast.FuncDecl{
							Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{receivername}, Type: ts.Name}}},
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
	file.Iterators = fds
	return nil
}
