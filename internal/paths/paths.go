package paths

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/conflicts"
	"github.com/ufukty/gonfique/internal/paths/declare"
	"github.com/ufukty/gonfique/internal/paths/export"
	"github.com/ufukty/gonfique/internal/paths/match"
	"github.com/ufukty/gonfique/internal/paths/pick"
	"github.com/ufukty/gonfique/internal/paths/replace"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"github.com/ufukty/gonfique/internal/transform"
)

type products struct {
	Imports       []string
	Declare, Auto []*ast.GenDecl
}

type args struct {
	node   ast.Node
	holder ast.Node
	path   []string
}

func with[E any](s []E, v E) []E {
	return append(slices.Clone(s), v)
}

func Process(ti *transform.Info, c *config.File, verbose bool) (*products, error) {
	declare := declare.New()
	export := export.New()
	replace := replace.New()

	// bfs (auto package needs it)
	queue := []args{{ti.Type, nil, []string{"<Config>"}}}
	later := []args{} // after finishing traversal in the current tree, start for declared types
	for len(queue) > 0 {
		node, holder, path := queue[0].node, queue[0].holder, queue[0].path
		recursion := true
		if holder != nil {
			cps := match.Matches(c.Paths, path)
			err := conflicts.Check(c.Paths, cps)
			if err != nil {
				return nil, fmt.Errorf("checking conflicts: %w", err)
			}

			if _, ok := pick.Dict(cps, c.Paths); ok {
				panic("to implement") // TODO:
			}

			if repl, ok := pick.Replace(cps, c.Paths); ok {
				if err := replace.Expression(repl, holder, path[len(path)-1]); err != nil {
					return nil, fmt.Errorf("replacing: %w", err)
				}
				recursion = false
			}

			rp := resolve.Path(strings.Join(path, "."))

			if decl, ok := pick.Declare(cps, c.Paths); ok {
				ts, err := declare.Declare(holder, path[len(path)-1], decl, rp)
				if err != nil {
					return nil, fmt.Errorf("declaring: %w", err)
				}
				if ts != nil {
					later = append(later, args{
						ts.Type, ts, []string{fmt.Sprintf("<%s>", decl)},
					})
				}
				recursion = false // manually perform traversal later once per declared type (not as many as its users)
			}

			if _, ok := pick.Export(cps, c.Paths); ok {
				if err := export.Type(rp, declare.Typenames(), holder, path[len(path)-1]); err != nil {
					return nil, fmt.Errorf("exporting: %w", err)
				}
			}
		}

		if recursion {
			switch n := node.(type) {
			case *ast.StructType:
				if n.Fields != nil && n.Fields.List != nil {
					for _, f := range n.Fields.List {
						if f != nil && f.Type != nil {
							// TODO: sort by field names to stabilize export typename generation
							queue = append(queue, args{f.Type, f, with(path, ti.Keys[f])})
						}
					}
				}

			case *ast.MapType:
				queue = append(queue, args{n.Key, n, with(path, "[key]")})
				queue = append(queue, args{n.Value, n, with(path, "[value]")})

			case *ast.ArrayType:
				queue = append(queue, args{n.Elt, n, with(path, "[]")})
			}
		}

		queue = queue[1:]
		if len(queue) == 0 && len(later) > 0 {
			queue, later = append(queue, later[0]), later[1:]
		}
	}

	err := declare.Conflicts()
	if err != nil {
		return nil, fmt.Errorf("checking conflicts on declare directives:\n%w", err)
	}

	return &products{
		Imports: replace.Imports,
		Declare: declare.Decls,
		Auto:    export.Decls,
	}, nil
}
