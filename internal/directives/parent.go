package directives

import (
	"go/ast"
	"go/token"
)

func (d *Directives) addParentRefs() error {
	for tn, details := range d.ParametersForTypenames.Parent {
		pf := &ast.Field{
			Names: []*ast.Ident{details.Fieldname.Ident()},
			Type: &ast.StarExpr{
				Star: token.NoPos,
				X:    details.ParentType.Ident(),
			},
			Tag: &ast.BasicLit{Kind: token.STRING, Value: "`yaml:\"-\"`"},
		}
		ty := d.TypeExprs[tn].(*ast.StructType)
		ty.Fields.List = append(ty.Fields.List, pf)
	}

	return nil
}
