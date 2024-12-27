package paths

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/holders"
	"github.com/ufukty/gonfique/internal/paths/conflicts"
	"github.com/ufukty/gonfique/internal/paths/export"
	"github.com/ufukty/gonfique/internal/paths/match"
	"github.com/ufukty/gonfique/internal/paths/pick"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"github.com/ufukty/gonfique/internal/transform"
	"golang.org/x/exp/maps"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func with[E any](s []E, v E) []E {
	return append(slices.Clone(s), v)
}

func mki2[K1, K2 comparable, V any](m map[K1]map[K2]V, k1 K1, k2 K2) {
	if _, ok := m[k1]; !ok {
		m[k1] = make(map[K2]V)
	}
	if _, ok := m[k1][k2]; !ok {
		m[k1][k2] = *new(V)
	}
}

func bfs(ti *transform.Info, c *config.File, paths []config.Path, ea *export.Agent) (map[resolve.Path]map[resolve.Path]holders.Node, []config.Path, error) {
	hs := map[resolve.Path]map[resolve.Path]holders.Node{}
	dictmap := map[resolve.Path]config.Dict{} // (struct/map) convertion
	used := map[config.Path]any{}

	type args struct {
		node, holder ast.Node
		path         []string
		mpath        []string // logging
	}
	queue := []args{{ti.Type, nil, []string{}, []string{}}}
	later := []args{} // after finishing traversal in the current tree, start for declared types

	iter := 0
	for len(queue) > 0 {
		cue := queue[0]
		queue = queue[1:]
		node, holder, path, mpath := cue.node, cue.holder, cue.path, cue.mpath

		rp := resolve.Path(strings.Join(path, "."))

		recursion := true
		if holder != nil { // not the root
			mp := resolve.Path(strings.Join(mpath, "."))
			mki2(hs, rp, mp)
			hs[rp][mp] = holders.Node{holder, path[len(path)-1]}

			cps := match.Matches(paths, path)
			for _, cp := range cps {
				used[cp] = nil
			}
			err := conflicts.Check(c.Rules, cps)
			if err != nil {
				return nil, nil, fmt.Errorf("checking conflicts: %w", err)
			}
			if _, ok := pick.Replace(cps, c.Rules); ok {
				recursion = false
			}

			if tn, ok := pick.Declare(cps, c.Rules); ok {
				term := fmt.Sprintf("<%s>", tn)
				rp = resolve.Path(term)
				path = []string{term}
			} else if _, ok := pick.Export(cps, c.Rules); ok {
				err := ea.Reserve(rp)
				if err != nil {
					return nil, nil, fmt.Errorf("reserve typename for value: %w", err)
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
					return nil, nil, fmt.Errorf("unexpected uninitialized field list")
				}
				for _, f := range n.Fields.List {
					if f.Type == nil {
						return nil, nil, fmt.Errorf("unexpected uninitialized field type")
					}
					// TODO: sort by field names to stabilize export typename generation
					if has(dictmap, rp) {
						queue = append(queue, args{f.Type, f, with(path, "[value]"), with(mpath, ti.Keys[f])})
					} else {
						queue = append(queue, args{f.Type, f, with(path, ti.Keys[f]), with(mpath, ti.Keys[f])})
					}
				}

			case *ast.ArrayType:
				queue = append(queue, args{n.Elt, n, with(path, "[]"), with(mpath, "[]")})

			case *ast.Ident:
				break

			default:
				return nil, nil, fmt.Errorf("implementation error, unexpected type (%T)", n)

			}
		}

		if len(queue) == 0 && len(later) > 0 {
			queue, later = append(queue, later[0]), later[1:]
		}
		iter++
		if iter == 2000 {
			return nil, nil, fmt.Errorf("iteration limit exceeded (report the issue)")
		}
	}
	return hs, maps.Keys(used), nil
}
