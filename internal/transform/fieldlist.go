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
		return fmt.Errorf(strings.Join(conflicts, ", "))
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

// merges 2 FieldList assumed do not contain same field with different types
func combine(a, b *ast.FieldList) *ast.FieldList {
	m := &ast.FieldList{}
	addedIdents := map[string]bool{}
	for _, f := range a.List {
		m.List = append(m.List, f)
		addedIdents[f.Names[0].Name] = true
	}
	for _, f := range b.List {
		if _, ok := addedIdents[f.Names[0].Name]; !ok {
			m.List = append(m.List, f)
			addedIdents[f.Names[0].Name] = true
		}
	}
	return m
}
