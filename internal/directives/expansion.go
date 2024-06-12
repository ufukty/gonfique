package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/matcher"
	"github.com/ufukty/gonfique/internal/models"
)

func expandKeypathsInDirectives(b *bundle.Bundle) error {
	for kp := range *b.Df {
		matches, err := matcher.FindTypeDefHoldersForKeypath(b.CfgType, kp, b.OriginalKeys)
		if err != nil {
			return fmt.Errorf("matching the rule: %w", err)
		}
		if len(matches) == 0 {
			fmt.Printf("No match for keypath: %s\n", kp)
		}
		kps := []models.FlattenKeypath{}
		for _, match := range matches {
			kps = append(kps, b.Keypaths[match])
		}
		b.Expansions[kp] = kps
	}
	return nil
}
