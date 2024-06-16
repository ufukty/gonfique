package directives

import (
	"fmt"
	"slices"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

func (d *Directives) debug() {
	usages := datas.Revmap(d.ElectedTypenames)

	fmt.Println("elected types:")
	for tn, kps := range usages {
		fmt.Printf("  %s:\n", tn)
		slices.Sort(kps)
		for _, kp := range kps {
			fmt.Printf("    %s\n", kp)
		}
	}

	effects := map[models.FlattenKeypath][]models.WildcardKeypath{}

	for wckp := range *d.b.Df {
		for _, kp := range d.Expansions[wckp] {
			if _, ok := effects[kp]; !ok {
				effects[kp] = []models.WildcardKeypath{}
			}
			effects[kp] = append(effects[kp], wckp)
		}
	}

}
