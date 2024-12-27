package resolve

import (
	"cmp"
	"strings"
)

type Path string

func (p Path) Segments() []string {
	return strings.Split(string(p), ".")
}

func DependencyFirst(a, b Path) int {
	if strings.Contains(string(a), string(b)) {
		return -1
	} else if strings.Contains(string(b), string(a)) {
		return 1
	}
	return cmp.Compare(string(a), string(b))
}
