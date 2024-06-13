package parent

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
)

func TypenameRequirements(b *bundle.Bundle) error {
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
