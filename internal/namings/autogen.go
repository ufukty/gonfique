package namings

import (
	"slices"

	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

func groupKeypathsByDepth(kps []models.Keypath) map[int][]models.Keypath {
	groups := map[int][]models.Keypath{}
	for _, kp := range kps {
		depth := len(kp.Segments())
		if _, ok := groups[depth]; !ok {
			groups[depth] = []models.Keypath{}
		}
		groups[depth] = append(groups[depth], kp)
	}
	return groups
}

func orderKeypaths(kps []models.Keypath) []models.Keypath {
	// 1. group by depth
	// 2. order each group alphabetically
	ordered := []models.Keypath{}
	grouped := groupKeypathsByDepth(kps)
	depths := maps.Keys(grouped)
	slices.Sort(depths)
	for _, depth := range depths {
		slices.Sort(grouped[depth])
		ordered = append(ordered, grouped[depth]...)
	}
	return ordered
}

func typenameForSegments(segments []string) models.TypeName {
	tn := ""
	for i, s := range segments {
		if s == "[]" {
			continue
		}
		tn += SafeTypeName(s, i != 0)
	}
	return models.TypeName(tn)
}

func GenerateTypenames(keypaths []models.Keypath) map[models.Keypath]models.TypeName {
	ordered := orderKeypaths(keypaths)
	tns := map[models.Keypath]models.TypeName{}
	reserved := map[models.TypeName]bool{}
	for _, kp := range ordered {
		segments := kp.Segments()
		for i := len(segments) - 1; i >= 0; i-- {
			tn := typenameForSegments(segments[i:])
			if _, found := reserved[tn]; !found {
				reserved[tn] = true
				tns[kp] = tn
				break
			}
		}
	}
	return tns
}
