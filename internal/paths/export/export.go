package export

import (
	"fmt"
	"go/ast"
	"go/token"
	"slices"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/holders"
	"github.com/ufukty/gonfique/internal/paths/export/auto"
	"github.com/ufukty/gonfique/internal/paths/resolve"
)

func (a *Agent) Type(rp resolve.Path, reserved []config.Typename, holder ast.Node, termination string) error {
	tn, ok := auto.Typename(rp, slices.Concat(reserved, a.typenames))
	if !ok {
		return fmt.Errorf("could not produce typename for %s", rp)
	}

	expr, err := holders.Get(holder, termination)
	if err != nil {
		return fmt.Errorf("getting type expression of target: %w", err)
	}
	err = holders.Set(holder, termination, tn.Ident())
	if err != nil {
		return fmt.Errorf("replacing type def with typename on target: %w", err)
	}
	gd := &ast.GenDecl{
		Doc:   &ast.CommentGroup{List: []*ast.Comment{{Text: fmt.Sprintf("// exported for %s", rp)}}},
		Tok:   token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{Name: tn.Ident(), Type: expr}},
	}

	a.Decls = append(a.Decls, gd)
	a.typenames = append(a.typenames, tn)
	return nil
}
