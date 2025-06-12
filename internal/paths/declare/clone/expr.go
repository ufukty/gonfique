package clone

import "go/ast"

func BadExpr(n *ast.BadExpr) *ast.BadExpr {
	if n == nil {
		return nil
	}
	return &ast.BadExpr{
		From: n.From, // 0,
		To:   n.To,   // 0,
	}
}

func Ident(n *ast.Ident) *ast.Ident {
	if n == nil {
		return nil
	}
	return &ast.Ident{
		NamePos: n.NamePos, // 0,
		Name:    n.Name,    // "",
	}
}

func Ellipsis(n *ast.Ellipsis) *ast.Ellipsis {
	if n == nil {
		return nil
	}
	return &ast.Ellipsis{
		Ellipsis: n.Ellipsis,  // 0,
		Elt:      Expr(n.Elt), // nil,
	}
}

func BasicLit(n *ast.BasicLit) *ast.BasicLit {
	if n == nil {
		return nil
	}
	return &ast.BasicLit{
		ValuePos: n.ValuePos, // 0,
		Kind:     n.Kind,     // 0,
		Value:    n.Value,    // "",
	}
}

func FuncLit(n *ast.FuncLit) *ast.FuncLit {
	if n == nil {
		return nil
	}
	return &ast.FuncLit{
		Type: FuncType(n.Type),  // ast.FuncType{},
		Body: BlockStmt(n.Body), // ast.BlockStmt{},
	}
}

func CompositeLit(n *ast.CompositeLit) *ast.CompositeLit {
	if n == nil {
		return nil
	}
	return &ast.CompositeLit{
		Type:       Expr(n.Type),  // nil,
		Lbrace:     n.Lbrace,      // 0,
		Elts:       Exprs(n.Elts), // []ast.Expr{},
		Rbrace:     n.Rbrace,      // 0,
		Incomplete: n.Incomplete,  // false,
	}
}

func ParenExpr(n *ast.ParenExpr) *ast.ParenExpr {
	if n == nil {
		return nil
	}
	return &ast.ParenExpr{
		Lparen: n.Lparen,  // 0,
		X:      Expr(n.X), // nil,
		Rparen: n.Rparen,  // 0,
	}
}

func SelectorExpr(n *ast.SelectorExpr) *ast.SelectorExpr {
	if n == nil {
		return nil
	}
	return &ast.SelectorExpr{
		X:   Expr(n.X),    // nil,
		Sel: Ident(n.Sel), // ast.Ident{},
	}
}

func IndexExpr(n *ast.IndexExpr) *ast.IndexExpr {
	if n == nil {
		return nil
	}
	return &ast.IndexExpr{
		X:      Expr(n.X),     // nil,
		Lbrack: n.Lbrack,      // 0,
		Index:  Expr(n.Index), // nil,
		Rbrack: n.Rbrack,      // 0,
	}
}

func IndexListExpr(n *ast.IndexListExpr) *ast.IndexListExpr {
	if n == nil {
		return nil
	}
	return &ast.IndexListExpr{
		X:       Expr(n.X),        // nil,
		Lbrack:  n.Lbrack,         // 0,
		Indices: Exprs(n.Indices), // []ast.Expr{},
		Rbrack:  n.Rbrack,         // 0,
	}
}

func SliceExpr(n *ast.SliceExpr) *ast.SliceExpr {
	if n == nil {
		return nil
	}
	return &ast.SliceExpr{
		X:      Expr(n.X),    // nil,
		Lbrack: n.Lbrack,     // 0,
		Low:    Expr(n.Low),  // nil,
		High:   Expr(n.High), // nil,
		Max:    Expr(n.Max),  // nil,
		Slice3: n.Slice3,     // false,
		Rbrack: n.Rbrack,     // 0,
	}
}

func TypeAssertExpr(n *ast.TypeAssertExpr) *ast.TypeAssertExpr {
	if n == nil {
		return nil
	}
	return &ast.TypeAssertExpr{
		X:      Expr(n.X),    // nil,
		Lparen: n.Lparen,     // 0,
		Type:   Expr(n.Type), // nil,
		Rparen: n.Rparen,     // 0,
	}
}

func CallExpr(n *ast.CallExpr) *ast.CallExpr {
	if n == nil {
		return nil
	}
	return &ast.CallExpr{
		Fun:      Expr(n.Fun),   // nil,
		Lparen:   n.Lparen,      // 0,
		Args:     Exprs(n.Args), // []ast.Expr{},
		Ellipsis: n.Ellipsis,    // 0,
		Rparen:   n.Rparen,      // 0,
	}
}

func StarExpr(n *ast.StarExpr) *ast.StarExpr {
	if n == nil {
		return nil
	}
	return &ast.StarExpr{
		Star: n.Star,    // 0,
		X:    Expr(n.X), // nil,
	}
}

func UnaryExpr(n *ast.UnaryExpr) *ast.UnaryExpr {
	if n == nil {
		return nil
	}
	return &ast.UnaryExpr{
		OpPos: n.OpPos,   // 0,
		Op:    n.Op,      // 0,
		X:     Expr(n.X), // nil,
	}
}

func BinaryExpr(n *ast.BinaryExpr) *ast.BinaryExpr {
	if n == nil {
		return nil
	}
	return &ast.BinaryExpr{
		X:     Expr(n.X), // nil,
		OpPos: n.OpPos,   // 0,
		Op:    n.Op,      // 0,
		Y:     Expr(n.Y), // nil,
	}
}

func KeyValueExpr(n *ast.KeyValueExpr) *ast.KeyValueExpr {
	if n == nil {
		return nil
	}
	return &ast.KeyValueExpr{
		Key:   Expr(n.Key),   // nil,
		Colon: n.Colon,       // 0,
		Value: Expr(n.Value), // nil,
	}
}

func ArrayType(n *ast.ArrayType) *ast.ArrayType {
	if n == nil {
		return nil
	}
	return &ast.ArrayType{
		Lbrack: n.Lbrack,    // 0,
		Len:    Expr(n.Len), // nil,
		Elt:    Expr(n.Elt), // nil,
	}
}

func StructType(n *ast.StructType) *ast.StructType {
	if n == nil {
		return nil
	}
	return &ast.StructType{
		Struct:     n.Struct,            // 0,
		Fields:     FieldList(n.Fields), // ast.FieldList{},
		Incomplete: n.Incomplete,        // false,
	}
}

func FuncType(n *ast.FuncType) *ast.FuncType {
	if n == nil {
		return nil
	}
	return &ast.FuncType{
		Func:       n.Func,                  // 0,
		TypeParams: FieldList(n.TypeParams), // ast.FieldList{},
		Params:     FieldList(n.Params),     // ast.FieldList{},
		Results:    FieldList(n.Results),    // ast.FieldList{},
	}
}

func InterfaceType(n *ast.InterfaceType) *ast.InterfaceType {
	if n == nil {
		return nil
	}
	return &ast.InterfaceType{
		Interface:  n.Interface,          // 0,
		Methods:    FieldList(n.Methods), // ast.FieldList{},
		Incomplete: n.Incomplete,         // false,
	}
}

func MapType(n *ast.MapType) *ast.MapType {
	if n == nil {
		return nil
	}
	return &ast.MapType{
		Map:   n.Map,         // 0,
		Key:   Expr(n.Key),   // nil,
		Value: Expr(n.Value), // nil,
	}
}

func ChanType(n *ast.ChanType) *ast.ChanType {
	if n == nil {
		return nil
	}
	return &ast.ChanType{
		Begin: n.Begin,       // 0,
		Arrow: n.Arrow,       // 0,
		Dir:   n.Dir,         // 0,
		Value: Expr(n.Value), // nil,
	}
}
