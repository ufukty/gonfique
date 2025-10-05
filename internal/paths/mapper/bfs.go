package mapper

import (
	"fmt"
	"go/ast"

	"go.ufukty.com/gonfique/internal/datas/inits"
	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/holders"
	"go.ufukty.com/gonfique/internal/paths/conflicts"
	"go.ufukty.com/gonfique/internal/paths/export"
	"go.ufukty.com/gonfique/internal/paths/mapper/absolute"
	"go.ufukty.com/gonfique/internal/paths/mapper/resolve"
	"go.ufukty.com/gonfique/internal/paths/match"
	"go.ufukty.com/gonfique/internal/paths/pick"
	"go.ufukty.com/gonfique/internal/transform"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

type products struct {
	Nodes   map[resolve.Path][]absolute.Path
	Holders map[absolute.Path]holders.Node
}

func Bfs(ti *transform.Info, c *config.File, paths []config.Path, ea *export.Agent) (*products, error) {
	p := &products{
		Nodes:   map[resolve.Path][]absolute.Path{},
		Holders: map[absolute.Path]holders.Node{},
	}
	dictmap := map[resolve.Path]config.Dict{} // (struct/map) convertion

	type args struct {
		node, holder ast.Node
		path         resolve.Path
		abspath      absolute.Path // logging
	}
	queue := []args{{ti.Type, nil, "", ""}}
	later := []args{} // after finishing traversal in the current tree, start for declared types

	iter := 0
	for len(queue) > 0 {
		cue := queue[0]
		queue = queue[1:]
		node, holder, rp, mp := cue.node, cue.holder, cue.path, cue.abspath

		recursion := true
		if holder != nil { // not the root
			inits.Key(p.Nodes, rp)
			p.Nodes[rp] = append(p.Nodes[rp], mp)
			p.Holders[mp] = holders.Node{holder, rp.Termination()}

			cps := match.Matches(paths, rp.Terms())
			err := conflicts.Check(c.Rules, cps)
			if err != nil {
				return nil, fmt.Errorf("checking conflicts: %w", err)
			}
			if _, ok := pick.Replace(cps, c.Rules); ok {
				recursion = false
			}

			if tn, ok := pick.Declare(cps, c.Rules); ok {
				rp = resolve.Path(fmt.Sprintf("<%s>", tn))
			} else if _, ok := pick.Export(cps, c.Rules); ok {
				err := ea.Reserve(rp)
				if err != nil {
					return nil, fmt.Errorf("reserve typename for value: %w", err)
				}
			}

			if _, ok := pick.Dict(cps, c.Rules); ok {
				dictmap[rp] = config.Map
			}
		}

		if recursion {
			switch n := node.(type) {
			case *ast.StructType:
				if n.Fields == nil || n.Fields.List == nil {
					return nil, fmt.Errorf("unexpected uninitialized field list")
				}
				for _, f := range n.Fields.List {
					if f.Type == nil {
						return nil, fmt.Errorf("unexpected uninitialized field type")
					}
					// TODO: sort by field names to stabilize export typename generation
					if has(dictmap, rp) {
						queue = append(queue, args{f.Type, f, rp.Sub("[value]"), mp.Sub(ti.Keys[f])})
					} else {
						queue = append(queue, args{f.Type, f, rp.Sub(ti.Keys[f]), mp.Sub(ti.Keys[f])})
					}
				}

			case *ast.ArrayType:
				queue = append(queue, args{n.Elt, n, rp.Sub("[]"), mp.Sub("[]")})

			case *ast.Ident:
				break

			default:
				return nil, fmt.Errorf("implementation error, unexpected type (%T)", n)

			}
		}

		if len(queue) == 0 && len(later) > 0 {
			queue, later = append(queue, later[0]), later[1:]
		}
		iter++
		if iter == 2000 {
			return nil, fmt.Errorf("iteration limit exceeded (report the issue)")
		}
	}
	return p, nil
}
