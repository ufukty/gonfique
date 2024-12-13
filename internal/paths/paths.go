package paths

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/conflicts"
	"github.com/ufukty/gonfique/internal/paths/expand"
	"github.com/ufukty/gonfique/internal/paths/pick"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"github.com/ufukty/gonfique/internal/transform"
)

func Process(ti *transform.Info, c *config.File, verbose bool) error {
	paths := resolve.Paths(ti)
	expansions, err := expand.Paths(ti, c, paths)
	if err != nil {
		return fmt.Errorf("expanding paths: %w", err)
	}
	rev := datas.RevSliceMap(expansions)
	err = conflicts.Check(rev, c)
	if err != nil {
		return fmt.Errorf("conflicting directives: %w", err)
	}
	pick.Values(rev, func(cp config.Path) config.Typename { return c.Paths[cp].Declare })

	exprs := map[resolve.Path]ast.Expr{}
	for p, n := range datas.Invmap(paths) {
		switch n := n.(type) {
		case *ast.Field:
			exprs[p] = n.Type
		case *ast.ArrayType:
			exprs[p] = n.Elt
		case *ast.MapType:
			panic("to implement")
		default:
			return fmt.Errorf("unknown holder type: %T", n)
		}
	}

	return nil
}
