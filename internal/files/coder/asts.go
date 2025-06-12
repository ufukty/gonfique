package coder

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/files/input/encoders"
)

func (c Coder) addParentRefAssignmentsFunction(dst *ast.File) {
	if len(c.ParentRefAssignments) == 0 {
		return
	}
	fd := &ast.FuncDecl{
		Name: ast.NewIdent("parentRefAssignments"),
		Type: &ast.FuncType{Params: &ast.FieldList{
			List: []*ast.Field{{
				Names: []*ast.Ident{c.ti},
				Type:  &ast.StarExpr{X: ast.NewIdent(c.Meta.Type)},
			}},
		}},
		Body: &ast.BlockStmt{List: c.ParentRefAssignments},
	}
	dst.Decls = append(dst.Decls, fd)
}

func (c Coder) addReaderFunction(dst *ast.File) error {
	var decoder *ast.SelectorExpr
	switch c.Encoding {
	case encoders.Json:
		decoder = &ast.SelectorExpr{
			X:   &ast.Ident{Name: "json"},
			Sel: &ast.Ident{Name: "NewDecoder"},
		}
	case encoders.Yaml:
		decoder = &ast.SelectorExpr{
			X:   &ast.Ident{Name: "yaml"},
			Sel: &ast.Ident{Name: "NewDecoder"},
		}
	default:
		return fmt.Errorf("unknown encoding: %q", c.Encoding)
	}

	fd := &ast.FuncDecl{
		Name: &ast.Ident{Name: "Read" + c.Meta.Type},
		Type: &ast.FuncType{
			Params: &ast.FieldList{List: []*ast.Field{{
				Names: []*ast.Ident{{Name: "path"}},
				Type:  &ast.Ident{Name: "string"},
			}}},
			Results: &ast.FieldList{List: []*ast.Field{
				{Type: &ast.StarExpr{X: &ast.Ident{Name: c.Meta.Type}}},
				{Type: &ast.Ident{Name: "error"}},
			}},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "file"}, &ast.Ident{Name: "err"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "os"}, Sel: &ast.Ident{Name: "Open"}},
							Args: []ast.Expr{&ast.Ident{Name: "path"}},
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{X: &ast.Ident{Name: "err"}, Op: token.NEQ, Y: &ast.Ident{Name: "nil"}},
					Body: &ast.BlockStmt{List: []ast.Stmt{
						&ast.ReturnStmt{Results: []ast.Expr{
							&ast.Ident{Name: "nil"},
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "fmt"}, Sel: &ast.Ident{Name: "Errorf"}},
								Args: []ast.Expr{
									&ast.BasicLit{Kind: token.STRING, Value: quotes("opening config file: %w")},
									&ast.Ident{Name: "err"},
								},
							},
						}},
					}},
				},
				&ast.DeferStmt{
					Call: &ast.CallExpr{
						Fun:  &ast.SelectorExpr{X: ast.NewIdent("file"), Sel: ast.NewIdent("Close")},
						Args: []ast.Expr{},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{c.ti},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.UnaryExpr{Op: token.AND, X: &ast.CompositeLit{
							Type: &ast.Ident{Name: c.Meta.Type},
						}},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "err"}},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   &ast.CallExpr{Fun: decoder, Args: []ast.Expr{&ast.Ident{Name: "file"}}},
							Sel: &ast.Ident{Name: "Decode"},
						},
						Args: []ast.Expr{c.ti},
					}},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{X: &ast.Ident{Name: "err"}, Op: token.NEQ, Y: &ast.Ident{Name: "nil"}},
					Body: &ast.BlockStmt{List: []ast.Stmt{
						&ast.ReturnStmt{Results: []ast.Expr{
							&ast.Ident{Name: "nil"},
							&ast.CallExpr{
								Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "fmt"}, Sel: &ast.Ident{Name: "Errorf"}},
								Args: []ast.Expr{
									&ast.BasicLit{Kind: token.STRING, Value: quotes("decoding config file: %w")},
									&ast.Ident{Name: "err"},
								},
							},
						}},
					}},
				},
			},
		},
	}

	if len(c.ParentRefAssignments) > 0 {
		fd.Body.List = append(fd.Body.List, &ast.ExprStmt{
			X: &ast.CallExpr{Fun: ast.NewIdent("parentRefAssignments"), Args: []ast.Expr{c.ti}},
		})
	}

	fd.Body.List = append(fd.Body.List, &ast.ReturnStmt{
		Results: []ast.Expr{c.ti, &ast.Ident{Name: "nil"}},
	})

	dst.Decls = append(dst.Decls, fd)

	return nil
}
