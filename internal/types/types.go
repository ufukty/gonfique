package types

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/transform"
	"github.com/ufukty/gonfique/internal/types/accessors"
	"github.com/ufukty/gonfique/internal/types/iterator"
	"github.com/ufukty/gonfique/internal/types/rules"
)

type aux struct {
	Accessors map[config.Typename][]*ast.FuncDecl
	Iterator  map[config.Typename]*ast.FuncDecl
}

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

// TODO: apply parent directive
// TODO: apply embed directive
// DONE: implement accessors
// DONE: implement iterator
func Apply(ti *transform.Info, c *config.File, decls map[config.Typename]*ast.GenDecl) (*aux, error) {
	accessors := accessors.New()
	iterator := iterator.New()

	tts := rules.TypeTargeting(c)
	rs := rules.Filter(c, tts)

	for tn, dirs := range rs {
		if !has(decls, tn) {
			return nil, fmt.Errorf("declaration for typename (%s) not found. has any value directed to declare this type?", tn)
		}
		fmt.Println(tn)

		if len(dirs.Accessors) > 0 {
			if err := accessors.Implement(ti, tn, decls[tn], dirs.Accessors); err != nil {
				return nil, fmt.Errorf("accessors: %w", err)
			}
		}

		// if dirs.Embed != "" {

		// }

		if dirs.Iterator {
			err := iterator.Implement(ti, tn, decls[tn])
			if err != nil {
				return nil, fmt.Errorf("iterator: %w", err)
			}
		}

		// if dirs.Parent == "" {

		// }
	}

	return &aux{Accessors: accessors.Decls, Iterator: iterator.Decls}, nil
}
