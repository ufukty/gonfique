package namings

import (
	"slices"

	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

func groupKeypathsByDepth(kps []models.FlattenKeypath) map[int][]models.FlattenKeypath {
	groups := map[int][]models.FlattenKeypath{}
	for _, kp := range kps {
		depth := len(kp.Segments())
		if _, ok := groups[depth]; !ok {
			groups[depth] = []models.FlattenKeypath{}
		}
		groups[depth] = append(groups[depth], kp)
	}
	return groups
}

func orderKeypaths(kps []models.FlattenKeypath) []models.FlattenKeypath {
	// 1. group by depth
	// 2. order each group alphabetically
	ordered := []models.FlattenKeypath{}
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
	l := len(segments)
	if l == 1 && segments[0] == "[]" {
		return "" // come back next round with 2 segments
	}
	tn := ""
	for i, s := range segments {
		if s != "[]" {
			tn += SafeTypeName(s, i != 0)
		} else if i == 0 {
			tn += "item"
		} else {
			tn += "Item"
		}
	}
	return models.TypeName(tn)

}

// FIXME: consider [] containing keypaths
func GenerateTypenames(keypaths []models.FlattenKeypath) map[models.FlattenKeypath]models.TypeName {
	ordered := orderKeypaths(keypaths)
	tns := map[models.FlattenKeypath]models.TypeName{}
	reserved := map[models.TypeName]bool{
		"": true, // defect
	}
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
