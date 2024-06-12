package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/directives/accessors"
	"github.com/ufukty/gonfique/internal/directives/parent"
	"github.com/ufukty/gonfique/internal/models"
	"github.com/ufukty/gonfique/internal/namings"
	"golang.org/x/exp/maps"
)

func checkTypenameRequirements(b *bundle.Bundle) error {
	if err := accessors.TypenameRequirements(b); err != nil {
		return fmt.Errorf("checking for accessors: %w", err)
	}
	if err := parent.TypenameRequirements(b); err != nil {
		return fmt.Errorf("checking for parent refs: %w", err)
	}
	b.NeededToBeReferred = datas.Uniq(b.NeededToBeReferred)
	b.NeededToBeDeclared = datas.Uniq(b.NeededToBeDeclared)
	return nil
}

func union(a, b []models.FlattenKeypath) []models.FlattenKeypath {
	ks := map[models.FlattenKeypath]bool{}
	for _, a := range a {
		ks[a] = true
	}
	for _, b := range b {
		ks[b] = true
	}
	return maps.Keys(ks)
}

func electTypenames(b *bundle.Bundle) {
	generatedTypenames := namings.GenerateTypenames(maps.Values(b.Keypaths))
	providedTypenames := map[models.FlattenKeypath]models.TypeName{}
	for wckp, dirs := range *b.Df {
		if dirs.Named != "" {
			kps := b.Expansions[wckp]
			for _, kp := range kps {
				providedTypenames[kp] = dirs.Named
			}
		}
	}

	for _, kp := range b.NeededToBeReferred {
		tn, ok := providedTypenames[kp.Parent()]
		if ok {
			b.ElectedTypenames[kp] = tn
			continue
		}
		id, ok := b.TypeExprs[kp].(*ast.Ident)
		if ok {
			b.ElectedTypenames[kp] = models.TypeName(id.Name)
			continue
		}
		b.ElectedTypenames[kp] = generatedTypenames[kp]
	}

	// for _, kp := range union(b.NeededToBeDeclared, b.NeededToBeReferred) {
	// 	toRefer := slices.Contains(b.NeededToBeReferred, kp)
	// 	toDeclare := slices.Contains(b.NeededToBeDeclared, kp)

	// 	if toRefer && toDeclare {

	// 	} else if toRefer {

	// 	} else if toDeclare {

	// 	}
	// }
}
