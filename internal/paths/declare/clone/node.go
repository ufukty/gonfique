package clone

import (
	"fmt"
	"go/ast"
)

func Decl(n ast.Decl) ast.Decl {
	switch n := n.(type) {
	case *ast.BadDecl:
		return BadDecl(n)
	case *ast.GenDecl:
		return GenDecl(n)
	case *ast.FuncDecl:
		return FuncDecl(n)
	default:
		panic(fmt.Sprintf("unknown declaration type: %T", n))
	}
}

func Expr(n ast.Expr) ast.Expr {
	switch n := n.(type) {
	case *ast.BadExpr:
		return BadExpr(n)
	case *ast.Ident:
		return Ident(n)
	case *ast.Ellipsis:
		return Ellipsis(n)
	case *ast.BasicLit:
		return BasicLit(n)
	case *ast.FuncLit:
		return FuncLit(n)
	case *ast.CompositeLit:
		return CompositeLit(n)
	case *ast.ParenExpr:
		return ParenExpr(n)
	case *ast.SelectorExpr:
		return SelectorExpr(n)
	case *ast.IndexExpr:
		return IndexExpr(n)
	case *ast.IndexListExpr:
		return IndexListExpr(n)
	case *ast.SliceExpr:
		return SliceExpr(n)
	case *ast.TypeAssertExpr:
		return TypeAssertExpr(n)
	case *ast.CallExpr:
		return CallExpr(n)
	case *ast.StarExpr:
		return StarExpr(n)
	case *ast.UnaryExpr:
		return UnaryExpr(n)
	case *ast.BinaryExpr:
		return BinaryExpr(n)
	case *ast.KeyValueExpr:
		return KeyValueExpr(n)
	case *ast.ArrayType:
		return ArrayType(n)
	case *ast.StructType:
		return StructType(n)
	case *ast.FuncType:
		return FuncType(n)
	case *ast.InterfaceType:
		return InterfaceType(n)
	case *ast.MapType:
		return MapType(n)
	case *ast.ChanType:
		return ChanType(n)
	default:
		panic(fmt.Sprintf("unknown expression type: %T", n))
	}
}

func Stmt(n ast.Stmt) ast.Stmt {
	switch n := n.(type) {
	case *ast.BadStmt:
		return BadStmt(n)
	case *ast.DeclStmt:
		return DeclStmt(n)
	case *ast.EmptyStmt:
		return EmptyStmt(n)
	case *ast.LabeledStmt:
		return LabeledStmt(n)
	case *ast.ExprStmt:
		return ExprStmt(n)
	case *ast.SendStmt:
		return SendStmt(n)
	case *ast.IncDecStmt:
		return IncDecStmt(n)
	case *ast.AssignStmt:
		return AssignStmt(n)
	case *ast.GoStmt:
		return GoStmt(n)
	case *ast.DeferStmt:
		return DeferStmt(n)
	case *ast.ReturnStmt:
		return ReturnStmt(n)
	case *ast.BranchStmt:
		return BranchStmt(n)
	case *ast.BlockStmt:
		return BlockStmt(n)
	case *ast.IfStmt:
		return IfStmt(n)
	case *ast.CaseClause:
		return CaseClause(n)
	case *ast.SwitchStmt:
		return SwitchStmt(n)
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt(n)
	case *ast.CommClause:
		return CommClause(n)
	case *ast.SelectStmt:
		return SelectStmt(n)
	case *ast.ForStmt:
		return ForStmt(n)
	case *ast.RangeStmt:
		return RangeStmt(n)
	default:
		panic(fmt.Sprintf("unknown statement type: %T", n))
	}
}

func Spec(n ast.Spec) ast.Spec {
	switch n := n.(type) {
	case *ast.ImportSpec:
		return ImportSpec(n)
	case *ast.ValueSpec:
		return ValueSpec(n)
	case *ast.TypeSpec:
		return TypeSpec(n)
	default:
		panic(fmt.Sprintf("unknown specification type: %T", n))
	}
}

func Node(n ast.Node) ast.Node {
	switch n := n.(type) {
	case ast.Decl:
		return Decl(n)
	case ast.Expr:
		return Expr(n)
	case ast.Spec:
		return Spec(n)
	case ast.Stmt:
		return Stmt(n)
	default:
		panic(fmt.Sprintf("unknown node type: %T", n))
	}
}
