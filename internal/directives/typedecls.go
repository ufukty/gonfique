package directives

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

func (d *Directives) implementTypeDeclarations() {
	uniq := map[models.TypeName]ast.Expr{}
	for _, kp := range d.NeededToBeDeclared {
		uniq[d.TypenamesElected[kp]] = d.TypeExprs[kp]
	}
	uniq = datas.MergeMaps(uniq, d.NamedTypeExprs)

	for tn, expr := range uniq {
		d.b.Named = append(d.b.Named, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{
				Name: tn.Ident(),
				Type: expr,
			}},
		})
	}
}

func (d *Directives) replaceTypeExpressionsWithIdents() error {
	for tn, kps := range d.TypenameUsers {
		for _, kp := range kps {
			holder := d.Holders[kp]
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
