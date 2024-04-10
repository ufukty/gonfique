package transform

import (
	"go/ast"
	"slices"
)

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
