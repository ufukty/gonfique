package tree

import (
	"fmt"
	"strings"
)

func indent(s string, last bool) string {
	if !last {
		return strings.Join(strings.Split(s, "\n"), "\n│  ")
	} else {
		return strings.Join(strings.Split(s, "\n"), "\n   ")
	}
}

// use in recursion
func List(heading string, items []string) string {
	msg := heading
	for i, item := range items {
		last := i == len(items)-1
		if !last {
			msg += fmt.Sprintf("\n├─ %s", indent(item, false))

		} else {
			msg += fmt.Sprintf("\n╰─ %s", indent(item, true))

		}
	}

	return msg
}
