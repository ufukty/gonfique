package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/files"
)

type manager struct {
	Bundle        *bundle.Bundle
	DirectiveFile *files.DirectiveFile
	UserAssigned  []ast.Node // user assigned types specified by `type` directive
}

func newManager(b *bundle.Bundle, df *files.DirectiveFile) *manager {
	return &manager{
		Bundle:        b,
		DirectiveFile: df,
		UserAssigned:  []ast.Node{},
	}
}

func (m *manager) accessors() error {
	return nil
}

func (m *manager) typeAssigning() error {
	return nil
}

func (m *manager) embedding() error {
	return nil
}

func (m *manager) named() error {
	return nil
}

func (m *manager) ApplyDirectives() error {
	if err := m.named(); err != nil {
		return fmt.Errorf(": %w", err)
	}
	if err := m.accessors(); err != nil {
		return fmt.Errorf(": %w", err)
	}
	if err := m.typeAssigning(); err != nil {
		return fmt.Errorf(": %w", err)
	}
	if err := m.embedding(); err != nil {
		return fmt.Errorf(": %w", err)
	}
	return nil
}

func Apply(b *bundle.Bundle, df *files.DirectiveFile) error {
	return newManager(b, df).ApplyDirectives()
}
