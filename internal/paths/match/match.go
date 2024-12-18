package match

import (
	"slices"

	"github.com/ufukty/gonfique/internal/files/config"
)

func consume(c []string) []string {
	return c[1:]
}

func downgrade(s []string) []string {
	s2 := slices.Clone(s)
	s2[0] = "*"
	return s2
}

var keywords = []string{"[]", "[key]", "[value]"}

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
		return matches(consume(c), r) || matches(downgrade(c), r) || matches(c, r[1:])
	case "*":
		return !slices.Contains(keywords, r[0]) && matches(c[1:], r[1:])
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
