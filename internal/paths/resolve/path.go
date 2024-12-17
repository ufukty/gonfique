package resolve

import "strings"

type Path string

func (p Path) Segments() []string {
	return strings.Split(string(p), ".")
}
