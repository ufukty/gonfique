package auto

import (
	"regexp"
	"slices"
	"strings"
	"unicode"

	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/paths/mapper/resolve"
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

var specials = []string{"[]", "[key]", "[value]"}

func typenameForSegments(segments []string) config.Typename {
	l := len(segments)
	if l == 1 && slices.Contains(specials, segments[0]) {
		return "" // come back next round with 2 segments
	}
	tn := ""
	for _, s := range segments {
		switch s {
		case "[key]":
			tn += "Key"
		case "[value]":
			tn += "Value"
		case "[]":
			tn += "Item"
		default:
			tn += safeTypeName(s)
		}
	}
	return config.Typename(tn)

}

var keywords = map[config.Typename]bool{
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

func isKeyword(tn config.Typename) bool {
	_, ok := keywords[tn]
	return ok
}

// FIXME: consider [] containing keypaths
// targets is map of keypaths and preference of exported typename
func Typename(rp resolve.Path, reserved []config.Typename) (config.Typename, bool) {
	segments := rp.Terms()
	for i := len(segments) - 1; i >= 0; i-- {
		tn := typenameForSegments(segments[i:])
		if !(isKeyword(tn) || slices.Contains(reserved, tn)) {
			return tn, true
		}
	}
	return "", false
}
