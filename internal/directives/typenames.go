package directives

import (
	"fmt"
	"go/ast"
	"slices"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

type featureusers struct {
	Named     []models.FlattenKeypath
	Parent    []models.FlattenKeypath
	Type      []models.FlattenKeypath
	Import    []models.FlattenKeypath
	Embed     []models.FlattenKeypath
	Accessors []models.FlattenKeypath
}

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

	// election
	for _, kp := range slices.Concat(d.NeededToBeReferred, d.NeededToBeDeclared) {
		tn, ok := d.TypenamesProvided[kp]
		if ok {
			d.TypenamesElected[kp] = tn
			continue
		}
		id, ok := d.TypeExprs[kp].(*ast.Ident)
		if ok {
			d.TypenamesElected[kp] = models.TypeName(id.Name)
			continue
		}
		if autogen, ok := d.TypenamesAutogen[kp]; ok {
			d.TypenamesElected[kp] = autogen
			continue
		}
		return fmt.Errorf("can't elect a typename for keypath: %s", kp)
	}
	d.TypenameUsers = datas.Revmap(d.TypenamesElected)

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
