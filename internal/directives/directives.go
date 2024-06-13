package directives

import (
	"fmt"
	"slices"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/directives/accessors"
	"github.com/ufukty/gonfique/internal/directives/parent"
	"github.com/ufukty/gonfique/internal/directives/typedecls"
)

func debug(b *bundle.Bundle) {
	fmt.Println("elected types:")
	for tn, kps := range b.Usages {
		fmt.Printf("  %s:\n", tn)
		slices.Sort(kps)
		for _, kp := range kps {
			fmt.Printf("    %s\n", kp)
		}
	}
}

func Apply(b *bundle.Bundle, dbg bool) error {
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
	if dbg {
		debug(b)
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
