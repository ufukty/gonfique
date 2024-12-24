package types

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/transform"
	"github.com/ufukty/gonfique/internal/types/accessors"
	"github.com/ufukty/gonfique/internal/types/rules"
)

type aux struct {
	Accessors map[config.Typename][]*ast.FuncDecl
}

// TODO: apply parent directive
// TODO: apply embed directive
// DONE: implement accessors
// TODO: implement iterator
func Apply(ti *transform.Info, c *config.File, decls map[config.Typename]*ast.GenDecl) (*aux, error) {
	accessors := accessors.New()

	tts := rules.TypeTargeting(c)
	rs := rules.Filter(c, tts)

	for tn, dirs := range rs {
		fmt.Println(tn)

		if len(dirs.Accessors) > 0 {
			if err := accessors.Implement(ti, tn, decls[tn], dirs.Accessors); err != nil {
				return nil, fmt.Errorf("accessors: %w", err)
			}
		}

		// if dirs.Embed != "" {

		// }

		// if dirs.Iterator {

		// }

		// if dirs.Parent == "" {

		// }
	}

	return &aux{Accessors: accessors.Decls}, nil
}
