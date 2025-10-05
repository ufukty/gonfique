package export

import (
	"fmt"
	"go/ast"
	"slices"

	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/paths/export/auto"
	"go.ufukty.com/gonfique/internal/paths/mapper/resolve"
)

type Agent struct {
	reserved  []config.Typename
	typenames map[resolve.Path]config.Typename
	Decls     map[config.Typename]*ast.GenDecl
}

func New(reserved []config.Typename) *Agent {
	return &Agent{
		reserved:  slices.Clone(reserved),
		typenames: map[resolve.Path]config.Typename{},
		Decls:     map[config.Typename]*ast.GenDecl{},
	}
}

func (a *Agent) Reserve(rp resolve.Path) error {
	tn, ok := auto.Typename(rp, a.reserved)
	if !ok {
		return fmt.Errorf("could not produce typename for %s", rp)
	}
	a.reserved = append(a.reserved, tn)
	a.typenames[rp] = tn
	return nil
}
