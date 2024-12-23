package combine

import (
	"go/ast"
)

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

// merges 2 FieldList assumed do not contain same field with different types
func FieldLists(fls ...*ast.FieldList) *ast.FieldList {
	combined := &ast.FieldList{}
	added := map[string]any{}
	for _, fl := range fls {
		if fl.List != nil {
			for _, f := range fl.List {
				if !has(added, f.Names[0].Name) {
					combined.List = append(combined.List, f)
					added[f.Names[0].Name] = nil
				}
			}
		}
	}
	return combined
}
