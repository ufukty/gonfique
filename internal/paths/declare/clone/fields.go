package clone

import "go/ast"

func Field(n *ast.Field) *ast.Field {
	if n == nil {
		return nil
	}
	return &ast.Field{
		Doc:     CommentGroup(n.Doc),     // ast.CommentGroup{},
		Names:   Idents(n.Names),         // []*ast.Ident{},
		Type:    Expr(n.Type),            // nil,
		Tag:     BasicLit(n.Tag),         // ast.BasicLit{},
		Comment: CommentGroup(n.Comment), // ast.CommentGroup{},
	}
}

func FieldList(n *ast.FieldList) *ast.FieldList {
	if n == nil {
		return nil
	}
	return &ast.FieldList{
		Opening: n.Opening,
		List:    Fields(n.List),
		Closing: n.Closing,
	}
}
