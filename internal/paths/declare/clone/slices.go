package clone

import "go/ast"

func Comments(s []*ast.Comment) []*ast.Comment {
	if s == nil {
		return nil
	}
	s2 := make([]*ast.Comment, len(s))
	for i, e := range s {
		s2[i] = Comment(e)
	}
	return s2
}

func Decls(s []ast.Decl) []ast.Decl {
	if s == nil {
		return nil
	}
	s2 := make([]ast.Decl, len(s))
	for i, e := range s {
		s2[i] = Decl(e)
	}
	return s2
}

func Exprs(s []ast.Expr) []ast.Expr {
	if s == nil {
		return nil
	}
	s2 := make([]ast.Expr, len(s))
	for i, e := range s {
		s2[i] = Expr(e)
	}
	return s2
}

func Fields(s []*ast.Field) []*ast.Field {
	if s == nil {
		return nil
	}
	s2 := make([]*ast.Field, len(s))
	for i, e := range s {
		s2[i] = Field(e)
	}
	return s2
}

func Idents(s []*ast.Ident) []*ast.Ident {
	if s == nil {
		return nil
	}
	s2 := make([]*ast.Ident, len(s))
	for i, e := range s {
		s2[i] = Ident(e)
	}
	return s2
}

func Stmts(s []ast.Stmt) []ast.Stmt {
	if s == nil {
		return nil
	}
	s2 := make([]ast.Stmt, len(s))
	for i, e := range s {
		s2[i] = Stmt(e)
	}
	return s2
}

func Specs(s []ast.Spec) []ast.Spec {
	if s == nil {
		return nil
	}
	s2 := make([]ast.Spec, len(s))
	for i, e := range s {
		s2[i] = Spec(e)
	}
	return s2
}
