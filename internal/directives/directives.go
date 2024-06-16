package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/directives/accessors"
)

func Apply(b *bundle.Bundle, verbose bool) error {
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
	if err := populateNamedTypeExprs(b); err != nil {
		return fmt.Errorf("checking for named directive: %w", err)
	}
	if err := checkConflictsForParentRefs(b); err != nil {
		return fmt.Errorf("checking conflicts for adding parent refs: %w", err)
	}
	if verbose {
		debug(b)
	}
	implementTypeDeclarations(b)
	if err := replaceTypeExpressionsWithIdents(b); err != nil {
		return fmt.Errorf("declaring named types: %w", err)
	}
	if err := accessors.Implement(b); err != nil {
		return fmt.Errorf("implementing accessor methods: %w", err)
	}
	if err := addParentRefs(b); err != nil {
		return fmt.Errorf("adding parent refs as fields: %w", err)
	}

	return nil
}
