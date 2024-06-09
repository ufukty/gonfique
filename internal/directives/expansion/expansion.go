package expansion

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/matcher"
	"github.com/ufukty/gonfique/internal/models"
)

func ExpandKeypathsInDirectives(b *bundle.Bundle) error {
	b.Expansions = map[models.WildcardKeypath][]ast.Node{}
	for kp := range *b.Df {
		matches, err := matcher.FindTypeDefHoldersForKeypath(b.CfgType, kp, b.OriginalKeys)
		if err != nil {
			return fmt.Errorf("matching the rule: %w", err)
		}
		if len(matches) == 0 {
			fmt.Printf("No match for keypath: %s\n", kp)
		}
		b.Expansions[kp] = matches
	}
	return nil
}
