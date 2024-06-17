package directives

import (
	"fmt"
	"go/ast"
	"slices"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

func (d *Directives) checkTypenameRequirements() {
	for _, kp := range d.FeaturesForKeypaths.Parent {
		d.NeededToBeDeclared = append(d.NeededToBeDeclared, kp)          // itself to declare
		d.NeededToBeReferred = append(d.NeededToBeReferred, kp.Parent()) // parent to refer
	}

	for _, kp := range d.FeaturesForKeypaths.Accessors {
		d.NeededToBeReferred = append(d.NeededToBeReferred, kp) // struct
		for _, field := range d.DirectivesForKeypaths[kp].Accessors {
			d.NeededToBeReferred = append(d.NeededToBeReferred, kp.WithFieldPath(field)) // its field
		}
	}
}

func (d *Directives) typenameElection() error {
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
