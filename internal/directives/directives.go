package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/directives/accessors"
	"github.com/ufukty/gonfique/internal/directives/parent"
	"github.com/ufukty/gonfique/internal/directives/typedecls"
)

func Apply(b *bundle.Bundle) error {
	populateKeypathsAndHolders(b)

	if err := populateExprs(b); err != nil {
		return fmt.Errorf("collecting type expressions for each keypaths: %w", err)
	}
	if err := expandKeypathsInDirectives(b); err != nil {
		return fmt.Errorf("expanding keypaths: %w", err)
	}
	if err := typenames(b); err != nil {
		return fmt.Errorf("listing, declaring typenames and swapping definitions: %w", err)
	}
	if err := typedecls.Implement(b); err != nil {
		return fmt.Errorf("declaring named types: %w", err)
	}
	if err := accessors.Implement(b); err != nil {
		return fmt.Errorf("implementing accessor methods: %w", err)
	}
	if err := parent.Implement(b); err != nil {
		return fmt.Errorf("adding parent refs as fields: %w", err)
	}

	return nil
}
