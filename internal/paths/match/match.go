package match

import (
	"regexp"
	"slices"

	"github.com/ufukty/gonfique/internal/files/config"
)

var components = []string{"[]", "[key]", "[value]"}

func isComponent(s string) bool {
	return slices.Contains(components, s)
}

var typename = regexp.MustCompile(`<\w+>`)

func isTypename(s string) bool {
	return typename.MatchString(s)
}

// returns true if [c]onfig path matches the [r]esolved path
func matches(c []string, r []string) bool {
	if len(c) == 0 {
		return len(r) == 0
	}
	if len(c) != 0 && len(r) == 0 {
		return len(c) == 1 && c[0] == "**"
	}
	if isTypename(c[0]) != isTypename(r[0]) {
		return false
	}
	switch c[0] {
	case "**":
		return matches(c[1:], r) || matches(c, r[1:]) || matches(c[1:], r[1:])
	case "*":
		return !isComponent(r[0]) && matches(c[1:], r[1:])
	default:
		return c[0] == r[0] && matches(c[1:], r[1:])
	}
}

func Matches(cps map[config.Path]config.Directives, r []string) []config.Path {
	ms := []config.Path{}
	for cp := range cps {
		if matches(cp.Segments(), r) {
			ms = append(ms, cp)
		}
	}
	return ms
}
