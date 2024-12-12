package conflicts

import (
	"fmt"
	"strings"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/resolve"
)

// use when parameters are already comparable (unlike slices, such as in accessors)
type withcompare[K comparable] map[K][]config.Path

// this is for convenience opposed to sliceParameterDetailsSources.AddSource()
func (m withcompare[K]) AddSource(parameters K, source config.Path) {
	if _, ok := m[parameters]; !ok {
		m[parameters] = []config.Path{}
	}
	m[parameters] = append(m[parameters], source)
}

type params struct {
	Declare withcompare[config.Typename]
	Export  withcompare[bool]
	Replace withcompare[string]
	Dict    withcompare[config.Dict]
}

func Check(expansions map[config.Path][]resolve.Path, c *config.File) error {
	sources := map[resolve.Path]params{}
	for p, cps := range datas.RevSliceMap(expansions) {
		ss := params{
			Declare: withcompare[config.Typename]{},
			Export:  withcompare[bool]{},
			Replace: withcompare[string]{},
			Dict:    withcompare[config.Dict]{},
		}
		for _, cp := range cps {
			ds := c.Paths[cp]
			ss.Declare.AddSource(ds.Declare, cp)
			ss.Export.AddSource(ds.Export, cp)
			ss.Replace.AddSource(ds.Replace, cp)
			ss.Dict.AddSource(ds.Dict, cp)
		}
		// remove default values
		delete(ss.Declare, config.Typename(""))
		delete(ss.Export, false)
		delete(ss.Replace, "")
		delete(ss.Dict, "")
		sources[p] = ss
	}

	conflicts := []string{}
	for kp, ss := range sources {
		if len(ss.Declare) > 1 {
			msg := fmt.Sprintf("%s: conflicting 'declare' parameters:", kp)
			for val, wckps := range ss.Declare {
				msg += fmt.Sprintf("\n  %v => %q", wckps, val)
			}
			conflicts = append(conflicts, msg)
		}
		if len(ss.Dict) > 1 {
			msg := fmt.Sprintf("%s: conflicting 'dict' parameters:", kp)
			for val, wckps := range ss.Dict {
				msg += fmt.Sprintf("\n  %v => %v", wckps, val)
			}
			conflicts = append(conflicts, msg)
		}
		if len(ss.Replace) > 1 {
			msg := fmt.Sprintf("%s: conflicting 'replace' parameters:", kp)
			for val, wckps := range ss.Replace {
				msg += fmt.Sprintf("\n  %v => %v", wckps, val)
			}
			conflicts = append(conflicts, msg)
		}
	}
	if len(conflicts) > 0 {
		return fmt.Errorf(strings.Join(conflicts, "\n"))
	}
	return nil
}
