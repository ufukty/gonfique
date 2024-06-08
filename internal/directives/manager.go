package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/directives/accessors"
	"github.com/ufukty/gonfique/internal/directives/assignment"
	"github.com/ufukty/gonfique/internal/directives/embedding"
	"github.com/ufukty/gonfique/internal/directives/named"
	"github.com/ufukty/gonfique/internal/files"
)

func Apply(b *bundle.Bundle, df *files.DirectiveFile) error {
	if err := named.Implement(); err != nil {
		return fmt.Errorf(": %w", err)
	}
	if err := accessors.Implement(b); err != nil {
		return fmt.Errorf(": %w", err)
	}
	if err := assignment.Implement(); err != nil {
		return fmt.Errorf(": %w", err)
	}
	if err := embedding.Implement(); err != nil {
		return fmt.Errorf(": %w", err)
	}
	return nil
}
