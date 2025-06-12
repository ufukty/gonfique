package absolute

import (
	"cmp"
	"strings"
)

type Path string

func (p Path) Terms() []string {
	if p == "" {
		return []string{}
	}
	return strings.Split(string(p), ".")
}

func (p Path) Termination() string {
	ss := p.Terms()
	return ss[len(ss)-1]
}

func (p Path) Sub(term string) Path {
	return Path(strings.Join(append(p.Terms(), term), "."))
}

func DependencyFirst(a, b Path) int {
	if strings.Contains(string(a), string(b)) {
		return -1
	} else if strings.Contains(string(b), string(a)) {
		return 1
	}
	return cmp.Compare(string(a), string(b))
}
