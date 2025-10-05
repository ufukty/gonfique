package rules

import (
	"regexp"
	"strings"

	"go.ufukty.com/gonfique/internal/files/config"
)

func TypeTargeting(c *config.File) map[config.Path]config.Typename {
	typename := regexp.MustCompile(`^<(\w+)>$`)
	types := map[config.Path]config.Typename{}
	for cp := range c.Rules {
		ss := cp.Segments()
		if len(ss) != 1 {
			continue
		}
		if !typename.MatchString(ss[0]) {
			continue
		}
		s := strings.TrimSuffix(strings.TrimPrefix(ss[0], "<"), ">")
		types[cp] = config.Typename(s)
	}
	return types
}

func Filter(c *config.File, filter map[config.Path]config.Typename) map[config.Typename]config.Directives {
	f := map[config.Typename]config.Directives{}
	for path, tn := range filter {
		f[tn] = c.Rules[path]
	}
	return f
}
