package paths

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas/sortby"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/conflicts"
	"github.com/ufukty/gonfique/internal/paths/declare"
	"github.com/ufukty/gonfique/internal/paths/dict"
	"github.com/ufukty/gonfique/internal/paths/export"
	"github.com/ufukty/gonfique/internal/paths/match"
	"github.com/ufukty/gonfique/internal/paths/pick"
	"github.com/ufukty/gonfique/internal/paths/replace"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"github.com/ufukty/gonfique/internal/transform"
	"github.com/ufukty/gonfique/internal/tree/bucket"
	"golang.org/x/exp/maps"
)

type products struct {
	Imports       []string
	Declare, Auto map[config.Typename]*ast.GenDecl
}

func Process(ti *transform.Info, c *config.File, verbose bool) (*products, error) {
	declare := declare.New()
	export := export.New([]config.Typename{}) // FIXME:
	replace := replace.New()
	paths := maps.Keys(c.Rules)

	hs, _, err := bfs(ti, c, paths, export)
	if err != nil {
		return nil, fmt.Errorf("traversing the type expression for first time to list the node paths: %w", err)
	}

	if verbose {
		b := bucket.New("dumping every node match with resolved paths")
		for rp, users := range sortby.KeyFunc(hs, resolve.DependencyFirst) {
			b := b.Sub(string(rp))
			for path := range sortby.KeyFunc(users, resolve.DependencyFirst) {
				b.Add(string(path))
			}
		}
		fmt.Println(b)
	}

	for rp, users := range sortby.KeyFunc(hs, resolve.DependencyFirst) {
		for _, user := range users {
			cps := match.Matches(paths, rp.Segments()) // TODO: store the results of search in bfs?
			err := conflicts.Check(c.Rules, cps)
			if err != nil {
				return nil, fmt.Errorf("checking conflicts: %w", err)
			}

			exportable := true

			if _, ok := pick.Dict(cps, c.Rules); ok {
				_, err := dict.ConvertToMap(user, ti)
				if err != nil {
					return nil, fmt.Errorf("converting dict to map: %w", err)
				}
			}

			if repl, ok := pick.Replace(cps, c.Rules); ok {
				if err := replace.Expression(user, repl); err != nil {
					return nil, fmt.Errorf("replacing: %w", err)
				}
			}

			if decl, ok := pick.Declare(cps, c.Rules); ok {
				_, err := declare.Declare(user, decl, rp)
				if err != nil {
					return nil, fmt.Errorf("declaring: %w", err)
				}
				// if ts != nil {
				// 	later = append(later, args{ts.Type, ts, []string{fmt.Sprintf("<%s>", decl)}})
				// }
				exportable = false
			}

			if _, ok := pick.Export(cps, c.Rules); ok && exportable {
				if err := export.Type(user, rp, declare.Typenames()); err != nil {
					return nil, fmt.Errorf("exporting: %w", err)
				}
			}
		}
	}

	err = declare.Conflicts()
	if err != nil {
		return nil, fmt.Errorf("checking conflicts on declare directives:\n%w", err)
	}

	return &products{
		Imports: replace.Imports,
		Declare: declare.Decls,
		Auto:    export.Decls,
	}, nil
}
