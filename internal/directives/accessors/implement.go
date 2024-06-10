package accessors

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
)

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

					fieldtypename, ok := b.GeneratedTypenames[kp.WithField(fieldname)]
					if !ok {
						return fmt.Errorf("can't find the assigned type name for field %q in struct: %s", fieldname, wildcardkp)
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
