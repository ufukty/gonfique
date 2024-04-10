package pkg

import (
	"fmt"
	"go/ast"
	"reflect"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

func PrintPathways(v reflect.Value, anc []string) {
	if v.IsZero() {
		return
	}

	switch t := v.Type(); t.Kind() {
	case reflect.Interface:
		PrintPathways(v.Elem(), anc)
		return

	case reflect.Map:
		iter := v.MapRange()
		m := map[string]reflect.Value{}
		for iter.Next() {
			m[iter.Key().String()] = iter.Value()
		}
		ks := maps.Keys(m)
		slices.Sort(ks)
		for _, k := range ks {
			PrintPathways(m[k], append(slices.Clone(anc), k))
		}

	default:
		fmt.Printf("%s: %s\n", strings.Join(anc, "."), v.String())
	}
}

func nodeString(n ast.Node) string {
	switch n := n.(type) {
	case *ast.Field:
		return fmt.Sprintf("%s (%p)", n.Names[0].Name, n)
	case *ast.TypeSpec:
		return fmt.Sprintf("%s (%p)", n.Name.Name, n)
	}
	return fmt.Sprintf("anonymous (%p)", n)
}

func nodeSliceString(ns []ast.Node) string {
	s := []string{}
	for _, n := range ns {
		s = append(s, nodeString(n))
	}
	return strings.Join(s, ", ")
}
