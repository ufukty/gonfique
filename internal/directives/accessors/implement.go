package accessors

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
)

func Implement(b *bundle.Bundle) error {
	if b.Df == nil {
		return fmt.Errorf("directive file is not populated")
	} else if b.ElectedTypenames == nil {
		return fmt.Errorf("elected type names are missing")
	}
	b.Accessors = []*ast.FuncDecl{}

	fieldsfortypes := map[models.TypeName]map[models.FieldName]models.TypeName{}
	for wckp, directives := range *b.Df {
		if directives.Accessors != nil {
			for _, kp := range b.Expansions[wckp] {
				tn := b.ElectedTypenames[kp]
				if _, ok := fieldsfortypes[tn]; !ok {
					fieldsfortypes[tn] = map[models.FieldName]models.TypeName{}
				}
				for _, fp := range directives.Accessors {
					fkp := kp.WithFieldPath(fp)
					ftn := b.ElectedTypenames[fkp]
					fn := b.Fieldnames[b.Holders[fkp]]
					fieldsfortypes[tn][fn] = ftn
				}
			}
		}
	}

	for tn, fields := range fieldsfortypes {
		for fn, ftn := range fields {
			b.Accessors = append(b.Accessors,
				generateGetter(tn, fn, ftn),
				generateSetter(tn, fn, ftn),
			)
		}
	}

	return nil
}
