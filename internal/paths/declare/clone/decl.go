package clone

import "go/ast"

func BadDecl(n *ast.BadDecl) *ast.BadDecl {
	if n == nil {
		return nil
	}
	return &ast.BadDecl{
		From: n.From, // 0,
		To:   n.To,   // 0,
	}
}

func GenDecl(n *ast.GenDecl) *ast.GenDecl {
	if n == nil {
		return nil
	}
	return &ast.GenDecl{
		Doc:    CommentGroup(n.Doc), // &ast.CommentGroup{},
		TokPos: n.TokPos,            // 0,
		Tok:    n.Tok,               // 0,
		Lparen: n.Lparen,            // 0,
		Specs:  Specs(n.Specs),      // []ast.Spec{},
		Rparen: n.Rparen,            // 0,
	}
}

func FuncDecl(n *ast.FuncDecl) *ast.FuncDecl {
	if n == nil {
		return nil
	}
	return &ast.FuncDecl{
		Doc:  CommentGroup(n.Doc), // &ast.CommentGroup{},
		Recv: FieldList(n.Recv),   // &ast.FieldList{},
		Name: Ident(n.Name),       // &ast.Ident{},
		Type: FuncType(n.Type),    // &ast.FuncType{},
		Body: BlockStmt(n.Body),   // &ast.BlockStmt{},
	}
}
