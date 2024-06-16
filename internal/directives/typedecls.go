package directives

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

func implementTypeDeclarations(b *bundle.Bundle) {
	uniq := map[models.TypeName]ast.Expr{}
	for _, kp := range b.NeededToBeDeclared {
		uniq[b.ElectedTypenames[kp]] = b.TypeExprs[kp]
	}
	uniq = datas.MergeMaps(uniq, b.NamedTypeExprs)
	
	for tn, expr := range uniq {
		b.Named = append(b.Named, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{
				Name: tn.Ident(),
				Type: expr,
			}},
		})
	}
}

func replaceTypeExpressionsWithIdents(b *bundle.Bundle) error {
	for tn, kps := range b.TypenameUsers {
		for _, kp := range kps {
			holder := b.Holders[kp]
			switch h := holder.(type) {
			case *ast.Field:
				h.Type = tn.Ident()
			case *ast.ArrayType:
				h.Elt = tn.Ident()
			default:
				return fmt.Errorf("replacing inline type definition with the name of declared type: unrecognized holder type: %T", holder)
			}
		}
	}
	return nil
}
