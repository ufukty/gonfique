package conflicts

import (
	"fmt"
	"strings"

	"github.com/ufukty/gonfique/internal/files/config"
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

func format[C comparable](vs valueSources[C]) string {
	msg := ""
	for value, cps := range vs {
		msg += fmt.Sprintf("value: %v\n", value)
		for _, cp := range cps {
			msg += fmt.Sprintf("%sset by: %s\n", indent(2), cp)
		}
	}
	return msg
}

// it returns an error of conflicting values on same target set by different paths
// it uses 'value' argument to get the value from correct directive
func assertSingleValue[C comparable](cps []config.Path, value func(cp config.Path) C) error {
	var zero C
	vs := valueSources[C]{}
	for _, cp := range cps {
		if v := value(cp); v != zero {
			vs.Add(v, cp)
		}
	}
	if len(vs) > 1 {
		return fmt.Errorf("%s", format(vs))
	}
	return nil
}

func Check(c map[config.Path]config.Directives, cps []config.Path) error {
	checks := map[string]error{
		"declare": assertSingleValue(cps, func(cp config.Path) config.Typename { return c[cp].Declare }),
		"export":  assertSingleValue(cps, func(cp config.Path) bool { return c[cp].Export }),
		"dict":    assertSingleValue(cps, func(cp config.Path) config.Dict { return c[cp].Dict }),
		"replace": assertSingleValue(cps, func(cp config.Path) string { return c[cp].Replace }),
	}

	conflicts := []string{}
	for directive, err := range checks {
		if err != nil {
			msg := fmt.Sprintf("directive: %s\n%s", directive, indentlines(err.Error(), 2))
			conflicts = append(conflicts, msg)
		}
	}
	if len(conflicts) > 0 {
		return fmt.Errorf("%s", strings.Join(conflicts, "\n"))
	}
	return nil
}
