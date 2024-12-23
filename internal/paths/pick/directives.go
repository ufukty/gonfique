package pick

import (
	"github.com/ufukty/gonfique/internal/files/config"
)

// it returns the first directive value
// for one target
// that is different than the zero value
func value[C comparable](paths []config.Path, value func(cp config.Path) C) (C, bool) {
	var zero C
	for _, cp := range paths {
		if v := value(cp); v != zero {
			return v, true
		}
	}
	return zero, false
}

func Dict(cps []config.Path, paths map[config.Path]config.Directives) (config.Dict, bool) {
	return value(cps, func(cp config.Path) config.Dict {
		if paths[cp].Dict == config.Map {
			return config.Map
		}
		return ""
	})
}

func Replace(cps []config.Path, paths map[config.Path]config.Directives) (string, bool) {
	return value(cps, func(cp config.Path) string {
		return paths[cp].Replace
	})
}

func Declare(cps []config.Path, paths map[config.Path]config.Directives) (config.Typename, bool) {
	return value(cps, func(cp config.Path) config.Typename {
		return paths[cp].Declare
	})
}

func Export(cps []config.Path, paths map[config.Path]config.Directives) (bool, bool) {
	return value(cps, func(cp config.Path) bool {
		return paths[cp].Export && paths[cp].Declare == "" && paths[cp].Replace == ""
	})
}
