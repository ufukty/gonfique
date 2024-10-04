package directives

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
)

func (d *Directives) checkKeypathsToReferTheirType() {
	for _, kp := range d.Features.Parent {
		d.ToRefer = append(d.ToRefer, kp)          // itself to declare
		d.ToRefer = append(d.ToRefer, kp.Parent()) // parent to refer
	}

	for _, kp := range d.Features.Accessors {
		d.ToRefer = append(d.ToRefer, kp) // struct
		for _, field := range d.Directives[kp].Accessors {
			d.ToRefer = append(d.ToRefer, kp.WithFieldPath(field)) // its field
		}
	}

	d.ToRefer = append(d.ToRefer, d.Features.Declare...)
	d.ToRefer = datas.Uniq(d.ToRefer)
}

func getAutogen(d *Directives) map[models.FlattenKeypath]models.TypeName {
	targets := map[models.FlattenKeypath]bool{}
	for _, kp := range d.Keypaths {
		targets[kp] = false
	}
	for _, kp := range d.Features.Export {
		targets[kp] = true // overwrite exported
	}
	autogen := namings.GenerateTypenames(targets)
	return autogen
}

func getProvided(d *Directives) map[models.FlattenKeypath]models.TypeName {
	provided := map[models.FlattenKeypath]models.TypeName{}
	for _, kp := range d.Features.Declare {
		provided[kp] = d.Directives[kp].Declare
	}
	return provided
}

func (d *Directives) typenameElection() error {
	autogen := getAutogen(d)
	provided := getProvided(d)
	for _, kp := range d.ToRefer {
		if tn, ok := provided[kp]; ok {
			d.Elected[kp] = tn
			continue
		}
		if id, ok := d.Exprs[kp].(*ast.Ident); ok {
			d.Elected[kp] = models.TypeName(id.Name)
			continue
		}
		if autogen, ok := autogen[kp]; ok {
			d.Elected[kp] = autogen
			continue
		}
		return fmt.Errorf("can't elect a typename for keypath: %s", kp)
	}
	d.Instances = datas.Revmap(d.Elected)
	return nil
}

func (d *Directives) implementTypeDeclarations() {
	neededToBeDeclared := []models.FlattenKeypath{}
	neededToBeDeclared = append(neededToBeDeclared, d.Features.Parent...)
	neededToBeDeclared = append(neededToBeDeclared, d.Features.Embed...)
	// declare referred types except string, int, etc.
	for _, kp := range d.ToRefer {
		if _, ok := d.Exprs[kp].(*ast.Ident); !ok {
			neededToBeDeclared = append(neededToBeDeclared, kp)
		}
	}
	neededToBeDeclared = datas.Uniq(neededToBeDeclared)

	uniq := map[models.TypeName]ast.Expr{}
	for _, kp := range neededToBeDeclared {
		uniq[d.Elected[kp]] = d.Exprs[kp]
	}
	// for _, tn := range d.FeaturesForTypenames.Named {
	// 	uniq[tn] = d.KeypathTypeExprs[d.TypenameUsers[tn][0]]
	// }

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
	for tn, kps := range d.Instances {
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
