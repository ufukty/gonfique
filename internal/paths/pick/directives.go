package pick

import (
	"github.com/ufukty/gonfique/internal/files/config"
)

// it returns the first directive value
// for one target
func Value[C comparable](paths []config.Path, value func(cp config.Path) C) (C, bool) {
	var zero C
	for _, cp := range paths {
		if v := value(cp); v != zero {
			return v, true
		}
	}
	return zero, false
}

type set struct {
	Declare *config.Typename
	Export  *bool
	Replace *string
	Dict    *config.Dict
}

// it returns the first directive value
// for every directive
// for one target
func Set(paths map[config.Path]config.PathConfig, cps []config.Path) set {
	s := set{}
	decl, ok := Value(cps, func(cp config.Path) config.Typename {
		return paths[cp].Declare
	})
	if ok {
		s.Declare = &decl
	}
	expo, ok := Value(cps, func(cp config.Path) bool {
		return paths[cp].Export && paths[cp].Declare == "" && paths[cp].Replace == ""
	})
	if ok {
		s.Export = &expo
	}
	repl, ok := Value(cps, func(cp config.Path) string {
		return paths[cp].Replace
	})
	if ok {
		s.Replace = &repl
	}
	dict, ok := Value(cps, func(cp config.Path) config.Dict {
		if v := paths[cp].Dict; v != config.Struct {
			return v
		}
		return ""
	})
	if ok {
		s.Dict = &dict
	}
	return s
}
