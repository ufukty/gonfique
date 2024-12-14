package pick

import (
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/resolve"
)

// it returns the directive value for each target
func Values[C comparable](rev map[resolve.Path][]config.Path, value func(cp config.Path) C) map[resolve.Path]C {
	var zero C
	picks := map[resolve.Path]C{}
	for rp, cps := range rev {
		for _, cp := range cps {
			if v := value(cp); v != zero {
				picks[rp] = v
				break
			}
		}
	}
	return picks
}
