package files

import (
	"go/ast"
	"go/token"
)

var readerFunc = &ast.FuncDecl{
	Name: &ast.Ident{
		Name: "ReadConfig",
	},
	Type: &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						{
							Name: "path",
						},
					},
					Type: &ast.Ident{
						Name: "string",
					},
				},
			},
		},
		Results: &ast.FieldList{
			List: []*ast.Field{
				{
					Type: &ast.Ident{
						Name: "Config",
					},
				},
				{
					Type: &ast.Ident{
						Name: "error",
					},
				},
			},
		},
	},
	Body: &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.Ident{
						Name: "f",
					},
					&ast.Ident{
						Name: "err",
					},
				},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X: &ast.Ident{
								Name: "os",
							},
							Sel: &ast.Ident{
								Name: "Open",
							},
						},
						Args: []ast.Expr{
							&ast.Ident{
								Name: "path",
							},
						},
					},
				},
			},
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.Ident{
						Name: "err",
					},
					Op: token.NEQ,
					Y: &ast.Ident{
						Name: "nil",
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CompositeLit{
									Type: &ast.Ident{
										Name: "Config",
									},
									Incomplete: false,
								},
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "fmt",
										},
										Sel: &ast.Ident{
											Name: "Errorf",
										},
									},
									Args: []ast.Expr{
										&ast.BasicLit{
											Kind:  token.STRING,
											Value: "\"opening config file: %w\"",
										},
										&ast.Ident{
											Name: "err",
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.Ident{
						Name: "cfg",
					},
				},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.CompositeLit{
						Type: &ast.Ident{
							Name: "Config",
						},
						Incomplete: false,
					},
				},
			},
			&ast.AssignStmt{
				Lhs: []ast.Expr{
					&ast.Ident{
						Name: "err",
					},
				},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X: &ast.Ident{
										Name: "yaml",
									},
									Sel: &ast.Ident{
										Name: "NewDecoder",
									},
								},
								Args: []ast.Expr{
									&ast.Ident{
										Name: "f",
									},
								},
							},
							Sel: &ast.Ident{
								Name: "Decode",
							},
						},
						Args: []ast.Expr{
							&ast.UnaryExpr{
								Op: token.AND,
								X: &ast.Ident{
									Name: "cfg",
								},
							},
						},
					},
				},
			},
			&ast.IfStmt{
				Cond: &ast.BinaryExpr{
					X: &ast.Ident{
						Name: "err",
					},
					Op: token.NEQ,
					Y: &ast.Ident{
						Name: "nil",
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CompositeLit{
									Type: &ast.Ident{
										Name: "Config",
									},
									Incomplete: false,
								},
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.Ident{
											Name: "fmt",
										},
										Sel: &ast.Ident{
											Name: "Errorf",
										},
									},
									Args: []ast.Expr{
										&ast.BasicLit{
											Kind:  token.STRING,
											Value: "\"decoding config file: %w\"",
										},
										&ast.Ident{
											Name: "err",
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.Ident{
						Name: "cfg",
					},
					&ast.Ident{
						Name: "nil",
					},
				},
			},
		},
	},
}
