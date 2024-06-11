package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/namings"
	"golang.org/x/exp/maps"
)

func Apply(b *bundle.Bundle) error {
	AllKeypathsForHolders(b)
	err := PopulateExprs(b)
	if err != nil {
		return fmt.Errorf("collecting type expressions for each keypaths: %w", err)
	}
	if err = ExpandKeypathsInDirectives(b); err != nil {
		return fmt.Errorf("expanding: %w", err)
	}
	MarkNeededNamedTypes(b)
	b.GeneratedTypenames = namings.GenerateTypenames(maps.Values(b.Keypaths))

	populateProvidedTypeNames(b)
	electTypeNames(b)

	err = ImplementNamedTypeDeclarations(b)
	if err != nil {
		return fmt.Errorf("declaring named types: %w", err)
	}
	if err = ImplementAccessors(b); err != nil {
		return fmt.Errorf("implement: %w", err)
	}

	return nil
}
