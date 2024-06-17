package directives

import (
	"fmt"
	"go/ast"
	"slices"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
	"golang.org/x/exp/maps"
)

func (d *Directives) typenames() error {
	// collect
	if err := d.typenameRequirementsForAccessors(); err != nil {
		return fmt.Errorf("checking requirements for accessors: %w", err)
	}
	if err := d.typenameRequirementsForParent(); err != nil {
		return fmt.Errorf("checking requirements for parent refs: %w", err)
	}

	d.NeededToBeReferred = datas.Uniq(d.NeededToBeReferred)
	d.NeededToBeDeclared = datas.Uniq(d.NeededToBeDeclared)

	autogeneratedTypenames := namings.GenerateTypenames(maps.Values(d.Keypaths))
	providedTypenames := map[models.FlattenKeypath]models.TypeName{}
	for wckp, dirs := range *d.b.Df {
		if dirs.Named != "" {
			kps := d.Expansions[wckp]
			for _, kp := range kps {
				providedTypenames[kp] = dirs.Named
			}
		}
	}

	// election
	for _, kp := range slices.Concat(d.NeededToBeReferred, d.NeededToBeDeclared) {
		tn, ok := providedTypenames[kp]
		if ok {
			d.ElectedTypenames[kp] = tn
			continue
		}
		id, ok := d.TypeExprs[kp].(*ast.Ident)
		if ok {
			d.ElectedTypenames[kp] = models.TypeName(id.Name)
			continue
		}
		if autogen, ok := autogeneratedTypenames[kp]; ok {
			d.ElectedTypenames[kp] = autogen
			continue
		}
		return fmt.Errorf("can't elect a typename for keypath: %s", kp)
	}
	d.TypenameUsers = datas.Revmap(d.ElectedTypenames)

	// declare referred types except string, int, etc.
	for _, kp := range d.NeededToBeReferred {
		if _, ok := d.TypeExprs[kp].(*ast.Ident); !ok {
			d.NeededToBeDeclared = append(d.NeededToBeDeclared, kp)
		}
	}

	d.NeededToBeReferred = datas.Uniq(d.NeededToBeReferred)
	d.NeededToBeDeclared = datas.Uniq(d.NeededToBeDeclared)

	return nil
}
