package accessors

import (
	"go/ast"
	"go/token"

	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/namings"
)

func get(tn config.Typename, fn config.Fieldname, ft ast.Expr) *ast.FuncDecl {
	recvname := namings.Initial(string(tn))
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{{
				Names: []*ast.Ident{{Name: recvname}},
				Type:  tn.Ident(),
			}},
		},
		Name: &ast.Ident{Name: "Get" + string(fn)},
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
			Results: &ast.FieldList{
				List: []*ast.Field{{
					Type: ft,
				}},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.SelectorExpr{
							X:   ast.NewIdent(recvname),
							Sel: fn.Ident(),
						},
					},
				},
			},
		},
	}
}

func set(tn config.Typename, fn config.Fieldname, ft ast.Expr) *ast.FuncDecl {
	recvname := namings.Initial(string(tn))
	paramname := "v"
	if recvname == "v" {
		paramname = "value"
	}
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{{
				Names: []*ast.Ident{{Name: recvname}},
				Type:  &ast.StarExpr{X: tn.Ident()},
			}},
		},
		Name: &ast.Ident{Name: "Set" + string(fn)},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{{
					Names: []*ast.Ident{{Name: paramname}},
					Type:  ft,
				}},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.SelectorExpr{
						X:   ast.NewIdent(recvname),
						Sel: fn.Ident(),
					}},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{&ast.Ident{Name: paramname}},
				},
			},
		},
	}
}
