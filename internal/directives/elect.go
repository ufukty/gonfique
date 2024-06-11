package directives

import (
	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
)

func populateProvidedTypeNames(b *bundle.Bundle) {
	b.ProvidedTypenames = map[models.FlattenKeypath]models.TypeName{}
	for wckp, dirs := range *b.Df {
		if dirs.Named != "" {
			matches := b.Expansions[wckp]
			for _, match := range matches {
				kp := b.Keypaths[match]
				b.ProvidedTypenames[kp] = dirs.Named
			}
		}
	}
}

func electTypeNames(b *bundle.Bundle) {
	b.ElectedTypenames = map[models.FlattenKeypath]models.TypeName{}
	for _, kp := range b.NeedsToBeNamed {
		if tn, ok := b.ProvidedTypenames[kp]; ok {
			b.ElectedTypenames[kp] = tn
		} else if tn, ok := b.GeneratedTypenames[kp]; ok {
			b.ElectedTypenames[kp] = tn
		}
	}
}
