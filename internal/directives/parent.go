package directives

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

func (d *Directives) parentEnabledTypenames() []models.TypeName {
	enabled := map[models.TypeName]bool{}
	for wckp, dirs := range *d.b.Df {
		if dirs.Parent != "" {
			for _, kp := range d.Expansions[wckp] {
				ty := d.TypeExprs[kp]
				_, ok := ty.(*ast.StructType)
				if !ok {
					continue
				}
				enabled[d.ElectedTypenames[kp]] = true
			}
		}
	}
	return maps.Keys(enabled)
}

func (d *Directives) checkConflictsForParentRefs() error {
	enabled := d.parentEnabledTypenames()
	for _, tn := range enabled {
		ptns := []models.TypeName{}
		for _, user := range d.TypenameUsers[tn] {
			ptns = append(ptns, d.ElectedTypenames[user.Parent()])
		}
		simplified := datas.Uniq(ptns)
		if len(simplified) > 1 {
			return fmt.Errorf("users of type %q have parents with different types: %v", tn, simplified)
		}
	}
	return nil
}

func (d *Directives) addParentRefs() error {
	type details struct {
		fieldname      models.FieldName
		parenttypename models.TypeName
	}
	toAdd := map[*ast.StructType]details{}

	for wckp, dirs := range *d.b.Df {
		if dirs.Parent != "" {
			for _, kp := range d.Expansions[wckp] {
				ty := d.TypeExprs[kp]
				st, ok := ty.(*ast.StructType)
				if !ok {
					continue
				}
				toAdd[st] = details{dirs.Parent, models.TypeName(d.ElectedTypenames[kp.Parent()])}
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

func (d *Directives) parent() error {
	for wckp, drs := range *d.b.Df {
		if drs.Parent != "" {
			kps, ok := d.Expansions[wckp]
			if !ok {
				return fmt.Errorf("expansions are not found for wildcard containing keypath: %s", wckp)
			}
			for _, kp := range kps {
				if _, ok = d.TypeExprs[kp].(*ast.StructType); !ok {
					fmt.Printf("warning: keypath %q directs to add a parent ref to a non-struct type (%s) is ignored\n", wckp, kp)
					continue
				}
				d.NeededToBeDeclared = append(d.NeededToBeDeclared, kp)          // itself to declare
				d.NeededToBeReferred = append(d.NeededToBeReferred, kp.Parent()) // parent to refer
			}
		}
	}
	return nil
}
