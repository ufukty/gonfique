package conflicts

import (
	"fmt"
	"strings"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/resolve"
)

type valueSources[C comparable] map[C][]config.Path

func (m valueSources[K]) Add(parameters K, source config.Path) {
	if _, ok := m[parameters]; !ok {
		m[parameters] = []config.Path{}
	}
	m[parameters] = append(m[parameters], source)
}

func indent(n int) string {
	s := ""
	for range n {
		s += " "
	}
	return s
}

func indentlines(s string, n int) string {
	return indent(n) + strings.ReplaceAll(s, "\n", "\n"+indent(n))
}

// it returns an error of conflicting values on same target set by different paths
// it uses 'value' argument to get the value from correct directive
func assertSingleValuePerPath[C comparable](rev map[resolve.Path][]config.Path, value func(cp config.Path) C) error {
	var zero C
	vss := map[resolve.Path]valueSources[C]{}
	for rp, cps := range rev {
		vss[rp] = valueSources[C]{}
		for _, cp := range cps {
			if v := value(cp); v != zero {
				vss[rp].Add(v, cp)
			}
		}
	}

	msgs := []string{}
	for rp, vs := range vss {
		if len(vs) > 1 {
			msg := fmt.Sprintf("path: %s\n", rp)
			for value, cps := range vs {
				msg += fmt.Sprintf("%svalue: %v\n", indent(2), value)
				for _, cp := range cps {
					msg += fmt.Sprintf("%sset by: %s\n", indent(4), cp)
				}
			}
			msgs = append(msgs, msg)
		}
	}

	if len(msgs) > 0 {
		return fmt.Errorf(strings.Join(msgs, ""))
	}
	return nil
}

func Check(rev map[resolve.Path][]config.Path, c *config.File) error {
	checks := map[string]error{
		"declare": assertSingleValuePerPath(rev, func(cp config.Path) config.Typename { return c.Paths[cp].Declare }),
		"export":  assertSingleValuePerPath(rev, func(cp config.Path) bool { return c.Paths[cp].Export }),
		"dict":    assertSingleValuePerPath(rev, func(cp config.Path) config.Dict { return c.Paths[cp].Dict }),
		"replace": assertSingleValuePerPath(rev, func(cp config.Path) string { return c.Paths[cp].Replace }),
	}

	conflicts := []string{}
	for directive, err := range checks {
		if err != nil {
			conflicts = append(conflicts, fmt.Sprintf("directive: %s\n%s", directive, indentlines(err.Error(), 2)))
		}
	}
	if len(conflicts) > 0 {
		return fmt.Errorf(strings.Join(conflicts, "\n"))
	}
	return nil
}
