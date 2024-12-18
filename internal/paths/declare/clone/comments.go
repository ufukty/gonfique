package clone

import "go/ast"

func Comment(n *ast.Comment) *ast.Comment {
	if n == nil {
		return nil
	}
	return &ast.Comment{
		Slash: n.Slash,
		Text:  n.Text,
	}
}

func CommentGroup(n *ast.CommentGroup) *ast.CommentGroup {
	if n == nil {
		return nil
	}
	return &ast.CommentGroup{
		List: Comments(n.List),
	}
}
