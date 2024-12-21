package match

import (
	"fmt"
	"strings"
	"testing"
)

func TestMatches(t *testing.T) {
	type tc struct {
		name     string
		pattern  string
		positive []string
		negative []string
	}

	tests := []tc{
		{
			"ExactMatch",
			"a.b.c",
			[]string{"a.b.c"},
			[]string{"", "a", "a.b.e", "b.a.c"},
		},
		{
			"TypeSegmentStart",
			"<Config>.a.b.c",
			[]string{"<Config>.a.b.c"},
			[]string{"<NotConfig>.a.b.c", "<>.a.b.c", "a.b.c"},
		},
		{
			"BothPathsEmpty",
			"",
			[]string{""},
			[]string{"a"},
		},
		{
			"Wildcard",
			"a.*.c",
			[]string{"a.b.c"},
			[]string{"a.b.d"},
		},
		{
			"DoubleWildcardAtStart",
			"**.d",
			[]string{"d", "a.d", "a.b.c.d"},
			[]string{"", "a.b.e", "d.e", "a.d.e"},
		},
		{
			"DoubleWildcardInMiddle",
			"a.**.d",
			[]string{"a.d", "a.b.d", "a.b.c.d", "a.d.d"},
			[]string{"", "a.b.e"},
		},
		{
			"DoubleWildcardAtEnd",
			"a.**",
			[]string{"a", "a.b", "a.[]", "a.[key]", "a.[value]", "a.b.c"},
			[]string{"", "b", "b.a", "[].a"},
		},
		{
			"TwiceDoubleWildcard",
			"a.**.e.**",
			[]string{"a.e", "a.b.e.f", "a.b.[].d.e", "a.b.[].d.e.f", "a.b.[].[value].e.f"},
			[]string{"e.a"},
		},
		{
			"Components",
			"a.[key].c",
			[]string{"a.[key].c"},
			[]string{"a.[value].c", "a.c"},
		},
		{
			"WildcardShouldNotMatchComponents",
			"a.*.c",
			[]string{},
			[]string{"a.[].c", "a.[key].c", "a.[value].c"},
		},
		{
			"ApiProject",
			"<Config>.gateways.**.endpoints.*",
			[]string{"<Config>.gateways.public.services.document.endpoints.list"},
			[]string{
				"<Config>.gateways.public.services.document.endpoints",
				"<Config>.gateways.public.services.document.endpoints.[]",
				"<Config>.gateways.public.services.document.endpoints.[key]",
				"<Config>.gateways.public.services.document.endpoints.[value]",
			},
		},
	}

	for _, tt := range tests {
		for _, tc := range tt.positive {
			t.Run(fmt.Sprintf("%s(%s)", tt.name, tc), func(t *testing.T) {
				if !matches(strings.Split(tt.pattern, "."), strings.Split(tc, ".")) {
					t.Errorf("should've match")
				}
			})
		}
		for _, tc := range tt.negative {
			t.Run(fmt.Sprintf("%s(%s)", tt.name, tc), func(t *testing.T) {
				if matches(strings.Split(tt.pattern, "."), strings.Split(tc, ".")) {
					t.Errorf("should've not match")
				}
			})
		}
	}
}
