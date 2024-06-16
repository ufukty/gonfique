package directives

import (
	"fmt"
	"slices"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

func debug(b *bundle.Bundle) {
	usages := datas.Revmap(b.ElectedTypenames)

	fmt.Println("elected types:")
	for tn, kps := range usages {
		fmt.Printf("  %s:\n", tn)
		slices.Sort(kps)
		for _, kp := range kps {
			fmt.Printf("    %s\n", kp)
		}
	}

	effects := map[models.FlattenKeypath][]models.WildcardKeypath{}

	for wckp := range *b.Df {
		for _, kp := range b.Expansions[wckp] {
			if _, ok := effects[kp]; !ok {
				effects[kp] = []models.WildcardKeypath{}
			}
			effects[kp] = append(effects[kp], wckp)
		}
	}

}
