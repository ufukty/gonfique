package typedecls

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/compares"
	"github.com/ufukty/gonfique/internal/models"
)

func replaceInlineTypeDefWithIdent(holder ast.Node, ident *ast.Ident) error {
	switch h := holder.(type) {
	case *ast.Field:
		h.Type = ident
	case *ast.ArrayType:
		h.Elt = ident
	default:
		return fmt.Errorf("unrecognized holder type: %s", reflect.TypeOf(holder).String())
	}
	return nil
}

func Implement(b *bundle.Bundle) error {
	if b.Holders == nil {
		return fmt.Errorf("holders is needed")
	}

	typesToInstances := map[models.TypeName][]models.FlattenKeypath{}

	for kp, tn := range b.ElectedTypenames {
		if _, ok := typesToInstances[tn]; !ok {
			typesToInstances[tn] = []models.FlattenKeypath{}
		}
		typesToInstances[tn] = append(typesToInstances[tn], kp)
	}

	// TODO: check all kps share the same type expression
	for tn, kps := range typesToInstances {
		for i := 1; i < len(kps); i++ {
			if !compares.Compare(b.TypeExprs[kps[0]], b.TypeExprs[kps[i]]) {
				return fmt.Errorf("%q and %q doesn't share the same schema, but required to share same type %q", kps[0], kps[i], tn)
			}
		}
	}

	neededToBeDeclared := map[models.TypeName]ast.Expr{}
	for _, kp := range b.NeededToBeDeclared {
		neededToBeDeclared[b.ElectedTypenames[kp]] = b.TypeExprs[kp]
	}

	for tn, expr := range neededToBeDeclared {
		b.Named = append(b.Named, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{
				Name: tn.Ident(),
				Type: expr,
			}},
		})
	}

	for tn, kps := range typesToInstances {
		for _, kp := range kps {
			if err := replaceInlineTypeDefWithIdent(b.Holders[kp], tn.Ident()); err != nil {
				return fmt.Errorf("replacing inline type definition with the name of declared type: %w", err)
			}
		}
	}
	return nil
}
