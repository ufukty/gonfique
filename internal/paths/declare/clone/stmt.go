package clone

import "go/ast"

func BadStmt(n *ast.BadStmt) *ast.BadStmt {
	if n == nil {
		return n
	}
	return &ast.BadStmt{
		From: n.From, // 0,
		To:   n.To,   // 0,
	}
}

func DeclStmt(n *ast.DeclStmt) *ast.DeclStmt {
	if n == nil {
		return n
	}
	return &ast.DeclStmt{
		Decl: n.Decl, // nil,
	}
}

func EmptyStmt(n *ast.EmptyStmt) *ast.EmptyStmt {
	if n == nil {
		return n
	}
	return &ast.EmptyStmt{
		Semicolon: n.Semicolon, // 0,
		Implicit:  n.Implicit,  // false,
	}
}

func LabeledStmt(n *ast.LabeledStmt) *ast.LabeledStmt {
	if n == nil {
		return n
	}
	return &ast.LabeledStmt{
		Label: Ident(n.Label), // ast.Ident{},
		Colon: n.Colon,        // 0,
		Stmt:  n.Stmt,         // nil,
	}
}

func ExprStmt(n *ast.ExprStmt) *ast.ExprStmt {
	if n == nil {
		return n
	}
	return &ast.ExprStmt{
		X: Expr(n.X), // nil,
	}
}

func SendStmt(n *ast.SendStmt) *ast.SendStmt {
	if n == nil {
		return n
	}
	return &ast.SendStmt{
		Chan:  Expr(n.Chan),  // nil,
		Arrow: n.Arrow,       // 0,
		Value: Expr(n.Value), // nil,
	}
}

func IncDecStmt(n *ast.IncDecStmt) *ast.IncDecStmt {
	if n == nil {
		return n
	}
	return &ast.IncDecStmt{
		X:      Expr(n.X), // nil,
		TokPos: n.TokPos,  // 0,
		Tok:    n.Tok,     // 0,
	}
}

func AssignStmt(n *ast.AssignStmt) *ast.AssignStmt {
	if n == nil {
		return n
	}
	return &ast.AssignStmt{
		Lhs:    Exprs(n.Lhs), // []ast.Expr{},
		TokPos: n.TokPos,     // 0,
		Tok:    n.Tok,        // 0,
		Rhs:    Exprs(n.Rhs), // []ast.Expr{},
	}
}

func GoStmt(n *ast.GoStmt) *ast.GoStmt {
	if n == nil {
		return n
	}
	return &ast.GoStmt{
		Go:   n.Go,             // 0,
		Call: CallExpr(n.Call), // ast.CallExpr{},
	}
}

func DeferStmt(n *ast.DeferStmt) *ast.DeferStmt {
	if n == nil {
		return n
	}
	return &ast.DeferStmt{
		Defer: n.Defer,          // 0,
		Call:  CallExpr(n.Call), // ast.CallExpr{},
	}
}

func ReturnStmt(n *ast.ReturnStmt) *ast.ReturnStmt {
	if n == nil {
		return n
	}
	return &ast.ReturnStmt{
		Return:  n.Return,         // 0,
		Results: Exprs(n.Results), // []ast.Expr{},
	}
}

func BranchStmt(n *ast.BranchStmt) *ast.BranchStmt {
	if n == nil {
		return n
	}
	return &ast.BranchStmt{
		TokPos: n.TokPos,       // 0,
		Tok:    n.Tok,          // 0,
		Label:  Ident(n.Label), // ast.Ident{},
	}
}

func BlockStmt(n *ast.BlockStmt) *ast.BlockStmt {
	if n == nil {
		return n
	}
	return &ast.BlockStmt{
		Lbrace: n.Lbrace,      // 0,
		List:   Stmts(n.List), // []ast.Stmt{},
		Rbrace: n.Rbrace,      // 0,
	}
}

func IfStmt(n *ast.IfStmt) *ast.IfStmt {
	if n == nil {
		return n
	}
	return &ast.IfStmt{
		If:   n.If,              // 0,
		Init: Stmt(n.Init),      // nil,
		Cond: Expr(n.Cond),      // nil,
		Body: BlockStmt(n.Body), // ast.BlockStmt{},
		Else: Stmt(n.Else),      // nil,
	}
}

func CaseClause(n *ast.CaseClause) *ast.CaseClause {
	if n == nil {
		return n
	}
	return &ast.CaseClause{
		Case:  n.Case,        // 0,
		List:  Exprs(n.List), // []ast.Expr{},
		Colon: n.Colon,       // 0,
		Body:  Stmts(n.Body), // []ast.Stmt{},
	}
}

func SwitchStmt(n *ast.SwitchStmt) *ast.SwitchStmt {
	if n == nil {
		return n
	}
	return &ast.SwitchStmt{
		Switch: n.Switch,          // 0,
		Init:   Stmt(n.Init),      // nil,
		Tag:    Expr(n.Tag),       // nil,
		Body:   BlockStmt(n.Body), // ast.BlockStmt{},
	}
}

func TypeSwitchStmt(n *ast.TypeSwitchStmt) *ast.TypeSwitchStmt {
	if n == nil {
		return n
	}
	return &ast.TypeSwitchStmt{
		Switch: n.Switch,          // 0,
		Init:   Stmt(n.Init),      // nil,
		Assign: Stmt(n.Assign),    // nil,
		Body:   BlockStmt(n.Body), // ast.BlockStmt{},
	}
}

func CommClause(n *ast.CommClause) *ast.CommClause {
	if n == nil {
		return n
	}
	return &ast.CommClause{
		Case:  n.Case,        // 0,
		Comm:  Stmt(n.Comm),  // nil,
		Colon: n.Colon,       // 0,
		Body:  Stmts(n.Body), // []ast.Stmt{},
	}
}

func SelectStmt(n *ast.SelectStmt) *ast.SelectStmt {
	if n == nil {
		return n
	}
	return &ast.SelectStmt{
		Select: n.Select,          // 0,
		Body:   BlockStmt(n.Body), // ast.BlockStmt{},
	}
}

func ForStmt(n *ast.ForStmt) *ast.ForStmt {
	if n == nil {
		return n
	}
	return &ast.ForStmt{
		For:  n.For,             // 0,
		Init: Stmt(n.Init),      // nil,
		Cond: Expr(n.Cond),      // nil,
		Post: Stmt(n.Post),      // nil,
		Body: BlockStmt(n.Body), // ast.BlockStmt{},
	}
}

func RangeStmt(n *ast.RangeStmt) *ast.RangeStmt {
	if n == nil {
		return n
	}
	return &ast.RangeStmt{
		For:    n.For,             // 0,
		Key:    Expr(n.Key),       // nil,
		Value:  Expr(n.Value),     // nil,
		TokPos: n.TokPos,          // 0,
		Tok:    n.Tok,             // 0,
		Range:  n.Range,           // 0,
		X:      Expr(n.X),         // nil,
		Body:   BlockStmt(n.Body), // ast.BlockStmt{},
	}
}
