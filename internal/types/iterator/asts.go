package iterator

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/files/config"
)

func quotes(s string) string {
	return fmt.Sprintf("%q", s)
}

func fd(tn config.Typename, fs []*ast.Field, commonType *ast.Ident, keys map[ast.Node]string) *ast.FuncDecl {
	key, value := ast.NewIdent("string"), commonType

	elts := []ast.Expr{}
	for _, f := range fs {
		kve := &ast.KeyValueExpr{
			Key:   &ast.BasicLit{Kind: token.STRING, Value: quotes(keys[f])},
			Value: &ast.SelectorExpr{X: tn.RecvName(), Sel: f.Names[0]},
		}
		elts = append(elts, kve)
	}

	fd := &ast.FuncDecl{
		Recv: &ast.FieldList{List: []*ast.Field{{
			Names: []*ast.Ident{tn.RecvName()},
			Type:  tn.Ident(),
		}}},
		Name: &ast.Ident{Name: "Fields"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
			Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.IndexListExpr{
				X:       &ast.SelectorExpr{X: &ast.Ident{Name: "iter"}, Sel: &ast.Ident{Name: "Seq2"}},
				Indices: []ast.Expr{key, value},
			}}}},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.FuncLit{
						Type: &ast.FuncType{Params: &ast.FieldList{List: []*ast.Field{{
							Names: []*ast.Ident{{Name: "yield"}},
							Type: &ast.FuncType{
								Params: &ast.FieldList{List: []*ast.Field{
									{Type: key},
									{Type: value},
								}},
								Results: &ast.FieldList{List: []*ast.Field{
									{Type: &ast.Ident{Name: "bool"}},
								}},
							},
						}}}},
						Body: &ast.BlockStmt{List: []ast.Stmt{
							&ast.AssignStmt{
								Lhs: []ast.Expr{&ast.Ident{Name: "mp"}},
								Tok: token.DEFINE,
								Rhs: []ast.Expr{
									&ast.CompositeLit{
										Type: &ast.MapType{Key: key, Value: value},
										Elts: elts,
									},
								},
							},
							&ast.RangeStmt{
								Key: &ast.Ident{Name: "k"}, Value: &ast.Ident{Name: "v"},
								Tok: token.DEFINE,
								X:   &ast.Ident{Name: "mp"},
								Body: &ast.BlockStmt{List: []ast.Stmt{
									&ast.IfStmt{
										Cond: &ast.UnaryExpr{
											Op: token.NOT,
											X: &ast.CallExpr{
												Fun:  &ast.Ident{Name: "yield"},
												Args: []ast.Expr{&ast.Ident{Name: "k"}, &ast.Ident{Name: "v"}},
											},
										},
										Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{}}},
									},
								}},
							},
						}},
					},
				},
			}},
		},
	}

	return fd
}
