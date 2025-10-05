package paths

import (
	"fmt"
	"go/ast"
	"slices"

	"go.ufukty.com/gonfique/internal/datas/sortby"
	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/paths/conflicts"
	"go.ufukty.com/gonfique/internal/paths/declare"
	"go.ufukty.com/gonfique/internal/paths/dict"
	"go.ufukty.com/gonfique/internal/paths/export"
	"go.ufukty.com/gonfique/internal/paths/mapper"
	"go.ufukty.com/gonfique/internal/paths/mapper/absolute"
	"go.ufukty.com/gonfique/internal/paths/mapper/resolve"
	"go.ufukty.com/gonfique/internal/paths/match"
	"go.ufukty.com/gonfique/internal/paths/pick"
	"go.ufukty.com/gonfique/internal/paths/replace"
	"go.ufukty.com/gonfique/internal/transform"
	"go.ufukty.com/gonfique/internal/tree/bucket"
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

	mps, err := mapper.Bfs(ti, c, paths, export)
	if err != nil {
		return nil, fmt.Errorf("traversing the type expression for first time to list the node paths: %w", err)
	}

	if verbose {
		b := bucket.New("dumping every node match with resolved paths")
		for rp, users := range sortby.KeyFunc(mps.Nodes, resolve.DependencyFirst) {
			b := b.Sub(string(rp))
			slices.SortFunc(users, absolute.DependencyFirst)
			for _, path := range users {
				b.Add(string(path))
			}
		}
		fmt.Println(b)
	}

	for rp, users := range sortby.KeyFunc(mps.Nodes, resolve.DependencyFirst) {
		cps := match.Matches(paths, rp.Terms()) // TODO: store the results of search in bfs?
		err := conflicts.Check(c.Rules, cps)
		if err != nil {
			return nil, fmt.Errorf("checking conflicts: %w", err)
		}

		var (
			_, okDict       = pick.Dict(cps, c.Rules)
			_, okExport     = pick.Export(cps, c.Rules)
			decl, okDeclare = pick.Declare(cps, c.Rules)
			repl, okReplace = pick.Replace(cps, c.Rules)
		)

		for _, node := range users {
			if okDict {
				_, err := dict.ConvertToMap(mps.Holders[node], ti)
				if err != nil {
					return nil, fmt.Errorf("converting dict to map: %w", err)
				}
			}
			if okReplace {
				if err := replace.Expression(mps.Holders[node], repl); err != nil {
					return nil, fmt.Errorf("replacing: %w", err)
				}
			}
			if okDeclare {
				_, err := declare.Declare(mps.Holders[node], decl, rp)
				if err != nil {
					return nil, fmt.Errorf("declaring: %w", err)
				}
			} else if okExport {
				if err := export.Type(mps.Holders[node], rp, declare.Typenames()); err != nil {
					return nil, fmt.Errorf("exporting: %w", err)
				}
			}
		}
	}

	// err = declare.Conflicts()
	// if err != nil {
	// 	return nil, fmt.Errorf("checking conflicts on declare directives:\n%w", err)
	// }

	return &products{
		Imports: replace.Imports,
		Declare: declare.Decls,
		Auto:    export.Decls,
	}, nil
}
