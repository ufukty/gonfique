package directives

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

type effectingWckps struct {
	named      []models.WildcardKeypath
	parent     []models.WildcardKeypath
	typeassign []models.WildcardKeypath
}

func (ewckp effectingWckps) Print() {
	if len(ewckp.named) > 0 {
		fmt.Printf("    named: %v\n", ewckp.named)
	}
	if len(ewckp.parent) > 0 {
		fmt.Printf("    parent: %v\n", ewckp.parent)
	}
	if len(ewckp.typeassign) > 0 {
		fmt.Printf("    type: %v\n", ewckp.typeassign)
	}
}

func caseInsensitiveCompareKeypaths(a, b models.FlattenKeypath) int {
	if strings.ToLower(string(a)) < strings.ToLower(string(b)) {
		return -1
	} else {
		return +1
	}
}

type effectCollection map[models.FlattenKeypath]effectingWckps

func (ec effectCollection) Print() {
	fmt.Println("rules grouped in effecting keypaths:")
	sorted := maps.Keys(ec)
	slices.SortFunc(sorted, caseInsensitiveCompareKeypaths)
	for _, kp := range sorted {
		dirs := ec[kp]
		fmt.Printf("  %s\n", kp)
		dirs.Print()
	}
}

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

	directives := effectCollection{}
	for kp, wckps := range effects {
		kpdirs := effectingWckps{}
		for _, wckp := range wckps {
			wckpdirs := (*b.Df)[wckp]
			if wckpdirs.Named != "" {
				kpdirs.named = append(kpdirs.named, wckp)
			}
			if wckpdirs.Parent != "" {
				kpdirs.parent = append(kpdirs.parent, wckp)
			}
			if wckpdirs.Type != "" {
				kpdirs.typeassign = append(kpdirs.typeassign, wckp)
			}
			directives[kp] = kpdirs
		}
	}

	directives.Print()

}
