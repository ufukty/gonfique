package match

import (
	"strings"
	"testing"
)

func TestMatches(t *testing.T) {
	type tc struct {
		name     string
		config   string
		resolved string
		expected bool
	}

	tests := []tc{
		// Test exact match
		{"Exact match", "a.b.c", "a.b.c", true},

		// Test empty paths
		{"Both paths empty", "", "", true},
		{"Config empty, resolved non-empty", "", "a", false},
		{"Resolved empty, config non-empty", "a", "", false},

		// Test single wildcard "*"
		{"Wildcard match", "a.*.c", "a.b.c", true},
		{"Wildcard mismatch", "a.*.c", "a.b.d", false},

		// Test double wildcard "**"
		{"Double wildcard match", "a.**.d", "a.b.c.d", true},
		{"Double wildcard at end", "a.**", "a.b.c", true},
		{"Double wildcard at end without match", "a.**", "a", false},
		{"Double wildcard mismatch", "a.**.e", "a.b.c.d", false},
		{"Twice double wildcard mismatch", "a.**.e.**", "a.b.c.d.e.f", true},
		{"Twice double wildcard mismatch", "a.**.e.**", "a.b.c.d.e.f", true},
		{"Double wildcard to pass keywords without last match", "a.**.e.**", "a.b.[].d.e", false},
		{"Double wildcard to pass keywords 2", "a.**.e.**", "a.b.[].d.e.f", true},
		{"Double wildcard to pass keywords 3", "a.**.e.**", "a.b.[].[value].e.f", true},

		// Test specific keywords
		{"Keyed match", "a.[key].c", "a.[key].c", true},
		{"Keyed mismatch", "a.[key].c", "a.[value].c", false},

		// Test wildcard exclusions
		{"Wildcard excludes []", "a.*.c", "a.[].c", false},
		{"Wildcard excludes [key]", "a.*.c", "a.[key].c", false},
		{"Wildcard excludes [value]", "a.*.c", "a.[value].c", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := matches(strings.Split(tt.config, "."), strings.Split(tt.resolved, "."))
			if result != tt.expected {
				t.Errorf("Test '%s' failed: matches(%v, %v) = %v; want %v", tt.name, tt.config, tt.resolved, result, tt.expected)
			}
		})
	}
}
