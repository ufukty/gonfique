package parent

import (
	"fmt"

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
				b.NeededToBeDeclared = append(b.NeededToBeDeclared, kp)          // it to declare
				b.NeededToBeReferred = append(b.NeededToBeReferred, kp.Parent()) // parent to refer
			}
		}
	}
	return nil
}
