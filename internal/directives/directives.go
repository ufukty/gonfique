package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/namings"
	"golang.org/x/exp/maps"
)

func Apply(b *bundle.Bundle) error {
	populateKeypathsAndHolders(b)
	err := populateExprs(b)
	if err != nil {
		return fmt.Errorf("collecting type expressions for each keypaths: %w", err)
	}
	if err = expandKeypathsInDirectives(b); err != nil {
		return fmt.Errorf("expanding: %w", err)
	}
	markNeededNamedTypes(b)
	b.GeneratedTypenames = namings.GenerateTypenames(maps.Values(b.Keypaths))

	populateProvidedTypeNames(b)
	electTypeNames(b)

	err = implementNamedTypeDeclarations(b)
	if err != nil {
		return fmt.Errorf("declaring named types: %w", err)
	}
	if err = implementAccessors(b); err != nil {
		return fmt.Errorf("implement: %w", err)
	}

	return nil
}
