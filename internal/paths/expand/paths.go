package expand

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/expand/matcher"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"github.com/ufukty/gonfique/internal/transform"
)

// TODO: update argument list after rewriting [matcher.matcher.FindHolders] for strings
func Paths(ti *transform.Info, c *config.File, holders map[resolve.Path]ast.Node) (map[config.Path][]resolve.Path, error) {
	paths := datas.Invmap(holders)
	expansions := map[config.Path][]resolve.Path{}

	m := matcher.New(ti)
	for p := range c.Paths {
		hs, err := m.FindHolders(p)
		if err != nil {
			return nil, fmt.Errorf("matching the rule: %w", err)
		}
		if len(hs) == 0 {
			fmt.Printf("No match for path: %s\n", p)
			continue
		}
		expansions[p] = []resolve.Path{}
		for _, h := range hs {
			expansions[p] = append(expansions[p], paths[h])
		}
	}

	return expansions, nil
}
