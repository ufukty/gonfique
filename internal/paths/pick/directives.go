package pick

import (
	"maps"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/resolve"
)

type valueSources[C comparable] map[C][]config.Path

func (vs valueSources[C]) Add(value C, source config.Path) {
	if _, ok := vs[value]; !ok {
		vs[value] = []config.Path{}
	}
	vs[value] = append(vs[value], source)
}

func (vs valueSources[C]) First() C {
	for k := range maps.Keys(vs) {
		return k
	}
	var zero C
	return zero
}

// it returns all values of one directive for each target
func Values[C comparable](rev map[resolve.Path][]config.Path, value func(cp config.Path) C) map[resolve.Path]C {
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
	picks := map[resolve.Path]C{}
	for rp, vs := range vss {
		picks[rp] = vs.First()
	}
	return picks
}
