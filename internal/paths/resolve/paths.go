package resolve

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/ufukty/gonfique/internal/transform"
)

type Path string

type FieldPath string

func (p Path) Segments() []string {
	return strings.Split(string(p), ".")
}

func (p Path) WithFieldPath(f FieldPath) Path {
	return Path(fmt.Sprintf("%s.%s", p, f))
}

func (p Path) Parent() Path {
	ss := p.Segments()
	l := max(len(ss)-1, 0)
	return Path(strings.Join(ss[:l], "."))
}

type resolver struct {
	ti    *transform.Info
	paths map[ast.Node]Path
}

func (r *resolver) dfs(n ast.Node, holder ast.Node, path []string) {
	if holder != nil {
		r.paths[holder] = Path(strings.Join(path, "."))
	}

	switch n := n.(type) {
	case *ast.StructType:
		if n.Fields != nil && n.Fields.List != nil {
			for _, f := range n.Fields.List {
				if f != nil && f.Type != nil {
					r.dfs(f.Type, f, append(path, r.ti.Keys[f])) // FIXME: for `dict: dynamic`
				}
			}
		}

	case *ast.ArrayType:
		r.dfs(n.Elt, n, append(path, "[]"))
	}
}

func Paths(ti *transform.Info) map[ast.Node]Path {
	r := resolver{ti: ti, paths: map[ast.Node]Path{}}
	r.dfs(ti.Type, nil, []string{})
	return r.paths
}
