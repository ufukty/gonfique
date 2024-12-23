package transform

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/compares"
)

func fieldsByTags(fl *ast.FieldList) map[string]*ast.Field {
	tgs := map[string]*ast.Field{}
	for _, f := range fl.List {
		tgs[f.Tag.Value] = f
	}
	return tgs
}

func areMergeable(a, b *ast.FieldList) error {
	bfs := fieldsByTags(b)
	conflicts := []string{}
	for _, af := range a.List {
		if bf, ok := bfs[af.Tag.Value]; ok {
			if !compares.Compare(af.Type, bf.Type) {
				conflicts = append(conflicts, af.Names[0].Name) // FIXME: ".Name" is the transformed version of the user-provided key
			}
		}
	}
	if len(conflicts) > 0 {
		return fmt.Errorf("%s", strings.Join(conflicts, ", "))
	}
	return nil
}

func sort(fl *ast.FieldList) {
	slices.SortFunc(fl.List, func(a, b *ast.Field) int {
		if a.Names[0].Name < b.Names[0].Name {
			return -1
		} else if a.Names[0].Name > b.Names[0].Name {
			return 1
		} else {
			return 0
		}
	})
}
