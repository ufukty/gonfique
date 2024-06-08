package accessors

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
)

func Implement(b *bundle.Bundle) error {
	for kp, directives := range *b.Df {
		if directives.Accessors != nil {

			structtypename, ok := b.Typenames[kp]
			if !ok {
				return fmt.Errorf("can't find the assigned type name for struct: %s", kp)
			}

			for _, fieldname := range directives.Accessors {

				fieldtypename, ok := b.Typenames[kp.WithField(fieldname)]
				if !ok {
					return fmt.Errorf("can't find the assigned type name for field %q in struct: %s", fieldname, kp)
				}

				b.Accessors = append(b.Accessors,
					generateGetter(structtypename, fieldname, fieldtypename),
					generateSetter(structtypename, fieldname, fieldtypename),
				)
			}
		}
	}
	return nil
}
