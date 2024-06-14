package parent

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

func ParentEnabledTypenames(b *bundle.Bundle) []models.TypeName {
	enabled := map[models.TypeName]bool{}
	for wckp, dirs := range *b.Df {
		if dirs.Parent != "" {
			for _, kp := range b.Expansions[wckp] {
				ty := b.TypeExprs[kp]
				_, ok := ty.(*ast.StructType)
				if !ok {
					continue
				}
				enabled[b.ElectedTypenames[kp]] = true
			}
		}
	}
	return maps.Keys(enabled)
}

func CheckConflicts(b *bundle.Bundle) error {
	typenameusers := datas.Revmap(b.ElectedTypenames)
	enabled := ParentEnabledTypenames(b)
	for _, tn := range enabled {
		ptns := []models.TypeName{}
		for _, user := range typenameusers[tn] {
			ptns = append(ptns, b.ElectedTypenames[user.Parent()])
		}
		simplified := datas.Uniq(ptns)
		if len(simplified) > 1 {
			return fmt.Errorf("can't decide the type of parent ref because parents of all users of the type %q use different types: %v", tn, simplified)
		}
	}
	return nil
}

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
