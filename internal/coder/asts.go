package coder

import (
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
)

func addReaderFunction(b *bundle.Bundle, dst *ast.File) {
	var decoder *ast.SelectorExpr
	switch b.Encoding {
	case models.Json:
		decoder = &ast.SelectorExpr{
			X: &ast.Ident{
				Name: "json",
			},
			Sel: &ast.Ident{
				Name: "NewDecoder",
			},
		}
	case models.Yaml:
		decoder = &ast.SelectorExpr{
			X: &ast.Ident{
				Name: "yaml",
			},
			Sel: &ast.Ident{
				Name: "NewDecoder",
			},
		}
	}

	readerFunc := &ast.FuncDecl{
		Name: &ast.Ident{
			Name: "Read" + b.TypeName,
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
							Name: b.TypeName,
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
							Name: "file",
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
											Name: b.TypeName,
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
				&ast.DeferStmt{
					Call: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("file"),
							Sel: ast.NewIdent("Close"),
						},
						Args: []ast.Expr{},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						&ast.Ident{
							Name: b.TypeNameInitial,
						},
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CompositeLit{
							Type: &ast.Ident{
								Name: b.TypeName,
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
									Fun: decoder,
									Args: []ast.Expr{
										&ast.Ident{
											Name: "file",
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
										Name: b.TypeNameInitial,
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
											Name: b.TypeName,
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
							Name: b.TypeNameInitial,
						},
						&ast.Ident{
							Name: "nil",
						},
					},
				},
			},
		},
	}

	dst.Decls = append(dst.Decls, readerFunc)
}
