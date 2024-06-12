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
	} else if b.ElectedTypenames == nil {
		return fmt.Errorf("elected type names are missing")
	}
	b.Accessors = []*ast.FuncDecl{}
	for wildcardkp, directives := range *b.Df {
		if directives.Accessors != nil {
			kps, ok := b.Expansions[wildcardkp]
			if !ok {
				return fmt.Errorf("expansion is not found for %q", wildcardkp)
			}
			for _, kp := range kps {
				structtypename, ok := b.GeneratedTypenames[kp]
				if !ok {
					return fmt.Errorf("generated typename is not found for %q", kp)
				}
				for _, fieldpath := range directives.Accessors {
					fieldtypename, ok := b.ElectedTypenames[kp.WithFieldPath(fieldpath)]
					if !ok {
						return fmt.Errorf("elected typename is not found for %q", kp.WithFieldPath(fieldpath))
					}
					holder, ok := b.Holders[kp.WithFieldPath(fieldpath)]
					if !ok {
						return fmt.Errorf("holder is not found for %q", kp.WithFieldPath(fieldpath))
					}
					electedFieldname := b.Fieldnames[holder]
					b.Accessors = append(b.Accessors,
						generateGetter(structtypename, electedFieldname, fieldtypename),
						generateSetter(structtypename, electedFieldname, fieldtypename),
					)
				}
			}
		}
	}
	return nil
}
