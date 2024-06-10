package accessors

import (
	"fmt"
	"go/ast"
	"slices"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
)

func lookFieldTypename(b *bundle.Bundle, kp models.FlattenKeypath) (models.TypeName, error) {
	if slices.Contains(b.NeedsToBeNamed, kp) {
		tn, ok := b.GeneratedTypenames[kp]
		if !ok {
			return "", fmt.Errorf("generated typename is not found for keyapth: %s", kp)
		}
		return tn, nil
	} else if ident, ok := b.TypeExprs[kp].(*ast.Ident); ok {
		return models.TypeName(ident.Name), nil
	} else {
		return "", fmt.Errorf("type name is not found")
	}
}

func Implement(b *bundle.Bundle) error {
	if b.Df == nil {
		return fmt.Errorf("directive file is not populated")
	} else if b.GeneratedTypenames == nil {
		return fmt.Errorf("typenames is missing")
	}
	b.Accessors = []*ast.FuncDecl{}

	for wildcardkp, directives := range *b.Df {
		if directives.Accessors != nil {
			matches, ok := b.Expansions[wildcardkp]
			if !ok {
				return fmt.Errorf("no match for keypath: %s", wildcardkp)
			}
			for _, match := range matches {
				kp, ok := b.Keypaths[match]
				if !ok {
					return fmt.Errorf("no flatten keypath found for wildcard keypath and match: %s / %s", wildcardkp, b.Keypaths[match])
				}

				structtypename, ok := b.GeneratedTypenames[kp]
				if !ok {
					return fmt.Errorf("can't find the assigned type name for struct: %s", wildcardkp)
				}

				for _, fieldname := range directives.Accessors {

					fieldtypename, err := lookFieldTypename(b, kp.WithField(fieldname))
					if err != nil {
						return fmt.Errorf("looking for correct typename: %w", err)
					}

					b.Accessors = append(b.Accessors,
						generateGetter(structtypename, fieldname, fieldtypename),
						generateSetter(structtypename, fieldname, fieldtypename),
					)
				}
			}
		}
	}
	return nil
}
