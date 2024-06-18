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
	for _, kp := range d.FeaturesForKeypaths.Parent {
		d.NeededToBeReferred = append(d.NeededToBeReferred, kp)          // itself to declare
		d.NeededToBeReferred = append(d.NeededToBeReferred, kp.Parent()) // parent to refer
	}

	for _, kp := range d.FeaturesForKeypaths.Accessors {
		d.NeededToBeReferred = append(d.NeededToBeReferred, kp) // struct
		for _, field := range d.DirectivesForKeypaths[kp].Accessors {
			d.NeededToBeReferred = append(d.NeededToBeReferred, kp.WithFieldPath(field)) // its field
		}
	}

	d.NeededToBeReferred = append(d.NeededToBeReferred, d.FeaturesForKeypaths.Named...)

	d.NeededToBeReferred = datas.Uniq(d.NeededToBeReferred)
}

func getAutogen(d *Directives) map[models.FlattenKeypath]models.TypeName {
	targets := map[models.FlattenKeypath]bool{}
	for _, kp := range d.Keypaths {
		targets[kp] = false
	}
	for _, kp := range d.FeaturesForKeypaths.Export {
		targets[kp] = true // overwrite exported
	}
	autogen := namings.GenerateTypenames(targets)
	return autogen
}

func getProvided(d *Directives) map[models.FlattenKeypath]models.TypeName {
	provided := map[models.FlattenKeypath]models.TypeName{}
	for _, kp := range d.FeaturesForKeypaths.Named {
		provided[kp] = d.DirectivesForKeypaths[kp].Named
	}
	return provided
}

func (d *Directives) typenameElection() error {
	autogen := getAutogen(d)
	provided := getProvided(d)
	for _, kp := range d.NeededToBeReferred {
		if tn, ok := provided[kp]; ok {
			d.TypenamesElected[kp] = tn
			continue
		}
		if id, ok := d.TypeExprs[kp].(*ast.Ident); ok {
			d.TypenamesElected[kp] = models.TypeName(id.Name)
			continue
		}
		if autogen, ok := autogen[kp]; ok {
			d.TypenamesElected[kp] = autogen
			continue
		}
		return fmt.Errorf("can't elect a typename for keypath: %s", kp)
	}
	d.TypenameUsers = datas.Revmap(d.TypenamesElected)
	return nil
}

func (d *Directives) checkKeypathsToModifyTheirType() {
	d.NeededToBeDeclared = append(d.NeededToBeDeclared, d.FeaturesForKeypaths.Parent...)
	d.NeededToBeDeclared = append(d.NeededToBeDeclared, d.FeaturesForKeypaths.Embed...)

	// declare referred types except string, int, etc.
	for _, kp := range d.NeededToBeReferred {
		if _, ok := d.TypeExprs[kp].(*ast.Ident); !ok {
			d.NeededToBeDeclared = append(d.NeededToBeDeclared, kp)
		}
	}

	d.NeededToBeDeclared = datas.Uniq(d.NeededToBeDeclared)
}

func (d *Directives) implementTypeDeclarations() {
	uniq := map[models.TypeName]ast.Expr{}
	for _, kp := range d.NeededToBeDeclared {
		uniq[d.TypenamesElected[kp]] = d.TypeExprs[kp]
	}
	for _, tn := range d.FeaturesForTypenames.Named {
		uniq[tn] = d.TypeExprs[d.TypenameUsers[tn][0]]
	}

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
