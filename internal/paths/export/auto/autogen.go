package auto

import (
	"regexp"
	"slices"
	"strings"
	"unicode"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"golang.org/x/exp/maps"
)

var smallcaps = regexp.MustCompile("[a-z]+")

func safeTypeName(keyname string) string {
	if smallcaps.Find([]byte(keyname)) == nil {
		keyname = strings.ToLower(keyname)
	}
	n := ""
	newSegment := false
	for i, r := range keyname {
		if i == 0 {
			n += strings.ToUpper(string(r))
		} else if newSegment {
			n += strings.ToUpper(string(r))
			newSegment = false
		} else if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			newSegment = true
		} else {
			n += string(r)
		}
	}
	return n
}

func groupKeypathsByDepth(kps []resolve.Path) map[int][]resolve.Path {
	groups := map[int][]resolve.Path{}
	for _, kp := range kps {
		depth := len(kp.Segments())
		if _, ok := groups[depth]; !ok {
			groups[depth] = []resolve.Path{}
		}
		groups[depth] = append(groups[depth], kp)
	}
	return groups
}

func orderKeypaths(kps []resolve.Path) []resolve.Path {
	// 1. group by depth
	// 2. order each group alphabetically
	ordered := []resolve.Path{}
	grouped := groupKeypathsByDepth(kps)
	depths := maps.Keys(grouped)
	slices.Sort(depths)
	for _, depth := range depths {
		slices.Sort(grouped[depth])
		ordered = append(ordered, grouped[depth]...)
	}
	return ordered
}

func typenameForSegments(segments []string) config.Typename {
	l := len(segments)
	if l == 1 && segments[0] == "[]" {
		return "" // come back next round with 2 segments
	}
	tn := ""
	for i, s := range segments {
		if s != "[]" {
			tn += safeTypeName(s)
		} else if i == 0 {
			tn += "item"
		} else {
			tn += "Item"
		}
	}
	return config.Typename(tn)

}

// FIXME: consider [] containing keypaths
// targets is map of keypaths and preference of exported typename
func GenerateTypenames(targets []resolve.Path, reserved []config.Typename) map[resolve.Path]config.Typename {
	ordered := orderKeypaths(targets)
	tns := map[resolve.Path]config.Typename{}
	reserved2 := map[config.Typename]bool{
		"":            true, // defect
		"break":       true,
		"case":        true,
		"chan":        true,
		"const":       true,
		"continue":    true,
		"default":     true,
		"defer":       true,
		"else":        true,
		"fallthrough": true,
		"for":         true,
		"func":        true,
		"go":          true,
		"goto":        true,
		"if":          true,
		"import":      true,
		"interface":   true,
		"map":         true,
		"package":     true,
		"range":       true,
		"return":      true,
		"select":      true,
		"struct":      true,
		"switch":      true,
		"type":        true,
		"var":         true,
	}
	for _, r := range reserved {
		reserved2[r] = true
	}
	for _, kp := range ordered {
		segments := kp.Segments()
		for i := len(segments) - 1; i >= 0; i-- {
			tn := typenameForSegments(segments[i:])
			if _, found := reserved2[tn]; !found {
				reserved2[tn] = true
				tns[kp] = tn
				break
			}
		}
	}
	return tns
}