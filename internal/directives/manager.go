package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/files"
)

type Manager struct {
	File          *files.File
	DirectiveFile *DirectiveFile
	UserAssigned  map[ast.Node]string // user assigned types specified by `type` directive
}

func NewManager(f *files.File, df *DirectiveFile) *Manager {
	return &Manager{
		File:          f,
		DirectiveFile: df,
		UserAssigned:  map[ast.Node]string{},
	}
}

func (m *Manager) accessors() error {
	return nil
}

func (m *Manager) typeAssigning() error {
	return nil
}

func (m *Manager) embedding() error {
	return nil
}

func (m *Manager) named() error {
	return nil
}

func (m *Manager) ApplyDirectives() error {
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
