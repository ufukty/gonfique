package export

import (
	"fmt"
	"go/ast"
	"slices"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/export/auto"
	"github.com/ufukty/gonfique/internal/paths/mapper/resolve"
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
