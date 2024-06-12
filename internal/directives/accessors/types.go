package accessors

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
)

func TypenameRequirements(b *bundle.Bundle) error {
	for wckp, drs := range *b.Df {
		if drs.Accessors != nil {
			kps, ok := b.Expansions[wckp]
			if !ok {
				return fmt.Errorf("expansion is found for: %s", wckp)
			}
			for _, kp := range kps {
				b.NeededToBeReferred = append(b.NeededToBeReferred, kp) // struct
				for _, field := range drs.Accessors {
					b.NeededToBeReferred = append(b.NeededToBeReferred, kp.WithFieldPath(field)) // its field
				}
			}
		}
	}
	return nil
}
