package collects

import (
	"fmt"

	"golang.org/x/exp/maps"
)

var ErrNoValues = fmt.Errorf("no values")

// Collects different values with their sources
type WithSources[Value comparable, Source ~string] map[Value][]Source

func (vs *WithSources[V, S]) Collect(value V, source S) {
	_, ok := (*vs)[value]
	if !ok {
		(*vs)[value] = []S{}
	}
	(*vs)[value] = append((*vs)[value], source)
}

// Returns "the value" if there is only one. Or returns an error which lists all contradicting values with sources
func (vs WithSources[V, S]) One() (V, error) {
	switch len(vs) {
	case 1:
		return maps.Keys(vs)[0], nil
	case 0:
		return *new(V), ErrNoValues
	default:
		err := "multiple values found (with sources):"
		for value, sources := range vs {
			var v any = value
			if stringer, ok := v.(fmt.Stringer); ok {
				err = fmt.Sprintf("%s\n+ %s (%s)", err, stringer, join(sources, ", "))
			} else {
				err = fmt.Sprintf("%s\n+ %p (%s)", err, &value, join(sources, ", "))
			}
		}
		return *new(V), fmt.Errorf("%s", err)
	}
}

func (vs WithSources[Value, Source]) Combine() []Value {
	return maps.Keys(vs)
}
