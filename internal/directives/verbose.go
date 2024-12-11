package directives

import (
	"fmt"
	"slices"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/models"
)

func (d *Directives) debug() {
	usages := datas.Revmap(d.elected)

	fmt.Println("elected types:")
	for tn, kps := range usages {
		fmt.Printf("  %s:\n", tn)
		slices.Sort(kps)
		for _, kp := range kps {
			fmt.Printf("    %s\n", kp)
		}
	}

	effects := map[models.FlattenKeypath][]config.Path{}

	for wckp := range *d.b.Df {
		for _, kp := range d.expansions[wckp] {
			if _, ok := effects[kp]; !ok {
				effects[kp] = []config.Path{}
			}
			effects[kp] = append(effects[kp], wckp)
		}
	}

}
