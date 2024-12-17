package match

import (
	"slices"

	"github.com/ufukty/gonfique/internal/files/config"
)

// returns true if [c]onfig path matches the [r]esolved path
func matches(c []string, r []string) bool {
	if len(c) == 0 && len(r) == 0 {
		return true
	}
	if len(c) == 0 || len(r) == 0 {
		return false
	}
	switch c[0] {
	case "**":
		return matches(c[1:], r) || matches(slices.Insert(c[1:], 0, "*"), r)
	case "*":
		return !slices.Contains([]string{"[]", "[key]", "[value]"}, r[0]) && matches(c[1:], r[1:])
	case "[]", "[key]", "[value]":
		fallthrough
	default:
		return c[0] == r[0] && matches(c[1:], r[1:])
	}
}

func Matches(cps map[config.Path]config.PathConfig, r []string) []config.Path {
	ms := []config.Path{}
	for cp := range cps {
		if matches(cp.Segments(), r) {
			ms = append(ms, cp)
		}
	}
	return ms
}
