package directives

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/datas/collects"
	"github.com/ufukty/gonfique/internal/namings"
	"github.com/ufukty/gonfique/internal/paths/models"
)

func (d *Directives) checkKeypathsToReferTheirType() {
	for _, kp := range d.features.Parent {
		d.toRefer = append(d.toRefer, kp)          // itself to declare
		d.toRefer = append(d.toRefer, kp.Parent()) // parent to refer
	}

	for _, kp := range d.features.Accessors {
		d.toRefer = append(d.toRefer, kp) // struct
		for _, field := range d.directives[kp].Accessors {
			d.toRefer = append(d.toRefer, kp.WithFieldPath(field)) // its field
		}
	}

	d.toRefer = append(d.toRefer, d.features.Declare...)
	d.toRefer = datas.Uniq(d.toRefer)
}

func getAutogen(d *Directives) map[models.FlattenKeypath]models.TypeName {
	targets := map[models.FlattenKeypath]bool{}
	for _, kp := range d.keypaths {
		targets[kp] = false
	}
	for _, kp := range d.features.Export {
		targets[kp] = true // overwrite exported
	}
	autogen := namings.GenerateTypenames(targets)
	return autogen
}

func getProvided(d *Directives) map[models.FlattenKeypath]models.TypeName {
	provided := map[models.FlattenKeypath]models.TypeName{}
	for _, kp := range d.features.Declare {
		provided[kp] = d.directives[kp].Declare
	}
	return provided
}

func (d *Directives) typenameElection() error {
	autogen := getAutogen(d)
	provided := getProvided(d)
	for _, kp := range d.toRefer {
		if tn, ok := provided[kp]; ok {
			d.elected[kp] = tn
			continue
		}
		if id, ok := d.exprs[kp].(*ast.Ident); ok {
			d.elected[kp] = models.TypeName(id.Name)
			continue
		}
		if autogen, ok := autogen[kp]; ok {
			d.elected[kp] = autogen
			continue
		}
		return fmt.Errorf("can't elect a typename for keypath: %s", kp)
	}
	d.instances = datas.Revmap(d.elected)
	return nil
}

func (d *Directives) implementTypeDeclarations() {
	declare := collects.Set[models.FlattenKeypath]{}
	declare.Add(d.features.Parent...)
	declare.Add(d.features.Embed...)
	// declare referred types except string, int, etc.
	for _, kp := range d.toRefer {
		if _, ok := d.exprs[kp].(*ast.Ident); !ok {
			declare.Add(kp)
		}
	}

	d.molds = map[models.TypeName]ast.Expr{}
	for kp := range declare.Iter() {
		d.molds[d.elected[kp]] = d.exprs[kp]
	}

	for tn, expr := range d.molds {
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
	for tn, kps := range d.instances {
		for _, kp := range kps {
			holder := d.holders[kp]
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
