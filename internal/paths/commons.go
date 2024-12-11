package paths

import (
	"fmt"
	"strings"
)

type FlattenKeypath string

func (kp FlattenKeypath) Segments() []string {
	return strings.Split(string(kp), ".")
}

func (kp FlattenKeypath) WithFieldPath(f FieldPath) FlattenKeypath {
	return FlattenKeypath(fmt.Sprintf("%s.%s", kp, f))
}

func (kp FlattenKeypath) Parent() FlattenKeypath {
	ss := kp.Segments()
	l := max(len(ss)-1, 0)
	return FlattenKeypath(strings.Join(ss[:l], "."))
}

type FieldPath string
