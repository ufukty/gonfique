package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/directives/accessors"
	"github.com/ufukty/gonfique/internal/directives/typedecls"
)

func Apply(b *bundle.Bundle) error {
	populateKeypathsAndHolders(b)
	err := populateExprs(b)
	if err != nil {
		return fmt.Errorf("collecting type expressions for each keypaths: %w", err)
	}
	if err = expandKeypathsInDirectives(b); err != nil {
		return fmt.Errorf("expanding keypaths: %w", err)
	}
	if err := checkTypenameRequirements(b); err != nil {
		return fmt.Errorf("checking for typenames needed to be either referred or declared: %w", err)
	}

	checkTypenameRequirements(b)
	electTypenames(b)

	err = typedecls.Implement(b)
	if err != nil {
		return fmt.Errorf("declaring named types: %w", err)
	}
	if err = accessors.Implement(b); err != nil {
		return fmt.Errorf("implement: %w", err)
	}

	return nil
}
