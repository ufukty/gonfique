package sort

import (
	"cmp"
	"fmt"
	"go/ast"
	"slices"
)

var basics = map[string]any{
	"bool":       nil,
	"byte":       nil,
	"complex128": nil,
	"complex64":  nil,
	"float32":    nil,
	"float64":    nil,
	"int":        nil,
	"int16":      nil,
	"int32":      nil,
	"int64":      nil,
	"int8":       nil,
	"rune":       nil,
	"string":     nil,
	"uint":       nil,
	"uint16":     nil,
	"uint32":     nil,
	"uint64":     nil,
	"uint8":      nil,
	"uintptr":    nil,
}

func isBasicType(n ast.Node) bool {
	i, ok := n.(*ast.Ident)
	if !ok {
		return false
	}
	_, ok = basics[i.Name]
	return ok
}

// returns if [dependent] mentions [dependency] somewhere in its type expression
func doesDependOn(dependent ast.Node, dependency *ast.Ident) bool {
	if isBasicType(dependent) { // can come from struct field types or array/map component types
		return false
	}

	switch dependent := dependent.(type) {
	// termination
	case *ast.Ident:
		return dependency.Name == dependent.Name
	case *ast.StarExpr:
		return dependent.X != nil && doesDependOn(dependent.X, dependency)
	case *ast.SelectorExpr:
		return false // caused by 'replace' directive

	// both for structs and functions (recv & arg list)
	case *ast.Field:
		if dependent.Type != nil && doesDependOn(dependent.Type, dependency) {
			return true
		}
	case *ast.FieldList:
		if dependent.List != nil {
			for _, f := range dependent.List {
				if doesDependOn(f.Type, dependency) {
					return true
				}
			}
		}

	case *ast.ArrayType:
		return dependent.Elt != nil && doesDependOn(dependent.Elt, dependency)
	case *ast.MapType:
		return (dependent.Key != nil && doesDependOn(dependent.Key, dependency)) ||
			(dependent.Value != nil && doesDependOn(dependent.Value, dependency))
	case *ast.StructType:
		if dependent.Fields != nil && doesDependOn(dependent.Fields, dependency) {
			return true
		}
	case *ast.TypeSpec:
		if dependent.Type != nil && doesDependOn(dependent.Type, dependency) {
			return true
		}
	case *ast.GenDecl:
		if dependent.Specs != nil {
			for _, spec := range dependent.Specs {
				if doesDependOn(spec, dependency) {
					return true
				}
			}
		}

	case *ast.FuncType:
		if dependent.Params != nil && doesDependOn(dependent.Params, dependency) {
			return true
		}
		if dependent.Results != nil && doesDependOn(dependent.Results, dependency) {
			return true
		}
	case *ast.FuncDecl:
		if dependent.Recv != nil && doesDependOn(dependent.Recv, dependency) {
			return true
		}
		if dependent.Type != nil && doesDependOn(dependent.Type, dependency) {
			return true
		}
	default:
		panic(fmt.Sprintf("unsupported dependent type (%T) with value (%v)", dependent, dependent))
	}
	return false
}

func symbolname(d ast.Decl) *ast.Ident {
	switch d := d.(type) {
	case *ast.GenDecl:
		if d.Specs != nil {
			return d.Specs[0].(*ast.TypeSpec).Name
		}
	case *ast.FuncDecl:
		return d.Name
	}
	return nil
}

func symbol(a, b ast.Decl) int {
	return cmp.Compare(symbolname(a).Name, symbolname(b).Name)
}

// map init key
func mik(m map[ast.Decl][]ast.Decl, k ast.Decl) {
	if _, ok := m[k]; !ok {
		m[k] = []ast.Decl{}
	}
}

func depgraph(decls []ast.Decl) (f, r map[ast.Decl][]ast.Decl) {
	vertices := map[ast.Decl][]ast.Decl{}
	rvertices := map[ast.Decl][]ast.Decl{}
	for _, d1 := range decls {
		for _, d2 := range decls {
			if d1 != d2 {
				if doesDependOn(d1, symbolname(d2)) {
					mik(vertices, d1)
					vertices[d1] = append(vertices[d1], d2)

					mik(rvertices, d2)
					rvertices[d2] = append(rvertices[d2], d1)

				} else if doesDependOn(d2, symbolname(d1)) {
					mik(vertices, d2)
					vertices[d2] = append(vertices[d2], d1)

					mik(rvertices, d1)
					rvertices[d1] = append(rvertices[d1], d2)
				}
			}
		}
	}
	return vertices, rvertices
}

func roots(nodes []ast.Decl, vertices map[ast.Decl][]ast.Decl) []ast.Decl {
	isRoot := map[ast.Decl]bool{}
	for _, n := range nodes {
		isRoot[n] = true
	}
	for _, k := range vertices {
		for _, n := range k {
			isRoot[n] = false
		}
	}
	roots := []ast.Decl{}
	for d, s := range isRoot {
		if s {
			roots = append(roots, d)
		}
	}
	return roots
}

func has[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func reverse[C comparable](cmp func(a, b C) int) func(a, b C) int {
	return func(a, b C) int {
		return cmp(b, a)
	}
}

type resolver struct {
	vertices, rvertices map[ast.Decl][]ast.Decl
	placed              map[ast.Decl]bool
	reversed            []ast.Decl
}

func (v *resolver) visit(c ast.Decl) {
	if !has(v.placed, c) {
		delay := false
		if has(v.rvertices, c) {
			for _, d := range v.rvertices[c] {
				if !has(v.placed, d) {
					delay = true
				}
			}
		}
		if !delay {
			v.reversed = append(v.reversed, c)
			v.placed[c] = true
		}
	}

	if has(v.vertices, c) {
		for _, a := range v.vertices[c] {
			v.visit(a)
		}
	}
}

func Decls(decls []ast.Decl) []ast.Decl {
	vertices, rvertices := depgraph(decls)
	for _, vs := range vertices {
		slices.SortFunc(vs, reverse(symbol))
	}
	roots := roots(decls, vertices)
	d := resolver{
		vertices:  vertices,
		rvertices: rvertices,
		placed:    map[ast.Decl]bool{},
		reversed:  []ast.Decl{},
	}
	for _, root := range roots {
		d.visit(root)
	}
	slices.Reverse(d.reversed)
	return d.reversed
}
