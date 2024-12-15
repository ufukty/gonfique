package paths

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/conflicts"
	"github.com/ufukty/gonfique/internal/paths/declare"
	"github.com/ufukty/gonfique/internal/paths/expand"
	"github.com/ufukty/gonfique/internal/paths/export"
	"github.com/ufukty/gonfique/internal/paths/pick"
	"github.com/ufukty/gonfique/internal/paths/replace"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"github.com/ufukty/gonfique/internal/transform"
	"golang.org/x/exp/maps"
)

type picks struct {
	declare map[resolve.Path]config.Typename
	export  map[resolve.Path]bool
	replace map[resolve.Path]string
	// dict    map[resolve.Path]config.Dict
}

type aux struct {
	Imports       []string
	Declare, Auto []*ast.GenDecl
}

// TODO: dict
// DONE: replace
// DONE: declare
// DONE: export
func Process(ti *transform.Info, c *config.File, verbose bool) (*aux, error) {
	holders := resolve.Holders(ti)
	expansions, err := expand.Paths(ti, c, holders)
	if err != nil {
		return nil, fmt.Errorf("expanding paths: %w", err)
	}
	rev := datas.RevSliceMap(expansions)
	err = conflicts.Check(rev, c)
	if err != nil {
		return nil, fmt.Errorf("conflicting directives: %w", err)
	}
	ps := picks{
		declare: pick.Values(rev, func(cp config.Path) config.Typename { return c.Paths[cp].Declare }),
		export: pick.Values(rev, func(cp config.Path) bool {
			return c.Paths[cp].Export && c.Paths[cp].Declare == "" && c.Paths[cp].Replace == ""
		}),
		replace: pick.Values(rev, func(cp config.Path) string { return c.Paths[cp].Replace }),
		// dict:    pick.Values(rev, func(cp config.Path) config.Dict { return c.Paths[cp].Dict }),
	}
	imports, err := replace.Expressions(ps.replace, holders)
	if err != nil {
		return nil, fmt.Errorf("replacing: %w", err)
	}
	decls, err := declare.Declare(ps.declare, holders)
	if err != nil {
		return nil, fmt.Errorf("declaring: %w", err)
	}
	reserved := datas.Uniq(maps.Values(ps.declare))
	auto, err := export.Types(maps.Keys(ps.export), reserved, holders)
	if err != nil {
		return nil, fmt.Errorf("export: %w", err)
	}
	a := &aux{
		Imports: imports,
		Declare: decls,
		Auto:    auto,
	}
	return a, nil
}
