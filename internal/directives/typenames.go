package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/directives/accessors"
	"github.com/ufukty/gonfique/internal/directives/parent"
	"github.com/ufukty/gonfique/internal/models"
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

func electedTypenames(b *bundle.Bundle) {
	b.ProvidedTypenames = map[models.FlattenKeypath]models.TypeName{}
	for wckp, dirs := range *b.Df {
		if dirs.Named != "" {
			kps := b.Expansions[wckp]
			for _, kp := range kps {
				b.ProvidedTypenames[kp] = dirs.Named
			}
		}
	}
}
