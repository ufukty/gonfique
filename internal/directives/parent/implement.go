package parent

import (
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
)

func Implement(b *bundle.Bundle) error {
	type details struct {
		fieldname      models.FieldName
		parenttypename models.TypeName
	}
	toAdd := map[*ast.StructType]details{}

	for wckp, dirs := range *b.Df {
		if dirs.Parent != "" {
			for _, kp := range b.Expansions[wckp] {
				ty := b.TypeExprs[kp]
				st, ok := ty.(*ast.StructType)
				if !ok {
					continue
				}
				toAdd[st] = details{dirs.Parent, models.TypeName(b.ElectedTypenames[kp.Parent()])}
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
