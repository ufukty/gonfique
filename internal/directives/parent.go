package directives

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

func parentEnabledTypenames(b *bundle.Bundle) []models.TypeName {
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

func checkConflictsForParentRefs(b *bundle.Bundle) error {
	enabled := parentEnabledTypenames(b)
	for _, tn := range enabled {
		ptns := []models.TypeName{}
		for _, user := range b.TypenameUsers[tn] {
			ptns = append(ptns, b.ElectedTypenames[user.Parent()])
		}
		simplified := datas.Uniq(ptns)
		if len(simplified) > 1 {
			return fmt.Errorf("users of type %q have parents with different types: %v", tn, simplified)
		}
	}
	return nil
}

func addParentRefs(b *bundle.Bundle) error {
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

func parent(b *bundle.Bundle) error {
	for wckp, drs := range *b.Df {
		if drs.Parent != "" {
			kps, ok := b.Expansions[wckp]
			if !ok {
				return fmt.Errorf("expansions are not found for wildcard containing keypath: %s", wckp)
			}
			for _, kp := range kps {
				if _, ok = b.TypeExprs[kp].(*ast.StructType); !ok {
					fmt.Printf("warning: keypath %q directs to add a parent ref to a non-struct type (%s) is ignored\n", wckp, kp)
					continue
				}
				b.NeededToBeDeclared = append(b.NeededToBeDeclared, kp)          // itself to declare
				b.NeededToBeReferred = append(b.NeededToBeReferred, kp.Parent()) // parent to refer
			}
		}
	}
	return nil
}
