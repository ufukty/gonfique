package directives

import (
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/models"
)

func (d *Directives) addParentRefs() error {
	type details struct {
		fieldname      models.FieldName
		parenttypename models.TypeName
	}
	toAdd := map[*ast.StructType]details{}

	for wckp, dirs := range *d.b.Df {
		if dirs.Parent != "" {
			for _, kp := range d.Expansions[wckp] {
				ty := d.TypeExprs[kp].(*ast.StructType)
				toAdd[ty] = details{dirs.Parent, models.TypeName(d.TypenamesElected[kp.Parent()])}
			}
		}
	}

	for expr, details := range toAdd {
		expr.Fields.List = append(expr.Fields.List, &ast.Field{
			Names: []*ast.Ident{details.fieldname.Ident()},
			Type: &ast.StarExpr{
				Star: token.NoPos,
				X:    details.parenttypename.Ident(),
			},
			Tag: &ast.BasicLit{Kind: token.STRING, Value: "`yaml:\"-\"`"},
		})
	}

	return nil
}
