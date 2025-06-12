package clone

import "go/ast"

func ImportSpec(n *ast.ImportSpec) *ast.ImportSpec {
	if n == nil {
		return nil
	}
	return &ast.ImportSpec{
		Doc:     CommentGroup(n.Doc),     // &ast.CommentGroup{},
		Name:    Ident(n.Name),           // &ast.Ident{},
		Path:    BasicLit(n.Path),        // &ast.BasicLit{},
		Comment: CommentGroup(n.Comment), // &ast.CommentGroup{},
		EndPos:  n.EndPos,                // 0,
	}
}

func ValueSpec(n *ast.ValueSpec) *ast.ValueSpec {
	if n == nil {
		return nil
	}
	return &ast.ValueSpec{
		Doc:     CommentGroup(n.Doc),     // &ast.CommentGroup{},
		Names:   Idents(n.Names),         // []*ast.Ident{},
		Type:    Expr(n.Type),            // nil,
		Values:  Exprs(n.Values),         // []ast.Expr{},
		Comment: CommentGroup(n.Comment), // &ast.CommentGroup{},
	}
}

func TypeSpec(n *ast.TypeSpec) *ast.TypeSpec {
	if n == nil {
		return nil
	}
	return &ast.TypeSpec{
		Doc:        CommentGroup(n.Doc),     // &ast.CommentGroup{},
		Name:       Ident(n.Name),           // &ast.Ident{},
		TypeParams: FieldList(n.TypeParams), // &ast.FieldList{},
		Assign:     n.Assign,                // 0,
		Type:       Expr(n.Type),            // nil,
		Comment:    CommentGroup(n.Comment), // &ast.CommentGroup{},
	}
}
