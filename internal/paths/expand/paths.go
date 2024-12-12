package expand

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/directives/resolve"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/expand/matcher"
	"github.com/ufukty/gonfique/internal/transform"
)

func Paths(ti *transform.Info, c *config.File, paths map[ast.Node]resolve.Path) (map[config.Path][]resolve.Path, error) {
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
