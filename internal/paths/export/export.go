package export

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/holders"
	"github.com/ufukty/gonfique/internal/paths/mapper/resolve"
)

func (a *Agent) Type(h holders.Node, rp resolve.Path, reserved []config.Typename) error {
	tn, ok := a.typenames[rp]
	if !ok {
		return fmt.Errorf("could not fetch the reserved typename for %s", rp)
	}

	expr, err := h.Get()
	if err != nil {
		return fmt.Errorf("getting type expression of target: %w", err)
	}
	err = h.Set(tn.Ident())
	if err != nil {
		return fmt.Errorf("replacing type def with typename on target: %w", err)
	}
	gd := &ast.GenDecl{
		Doc:   &ast.CommentGroup{List: []*ast.Comment{{Text: fmt.Sprintf("// exported for %s", rp)}}},
		Tok:   token.TYPE,
		Specs: []ast.Spec{&ast.TypeSpec{Name: tn.Ident(), Type: expr}},
	}

	a.Decls[tn] = gd
	return nil
}
