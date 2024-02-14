package pkg

import (
	"go/ast"
	"log"
	"reflect"
	"slices"
	"strings"
)

func degraded(keys []string) []string {
	keys2 := slices.Clone(keys)
	keys2[0] = "*" // replaces "**"
	return keys2
}

type node struct {
	i *ast.Ident
	v reflect.Value // to let owner able to put another node to its place in AST
}

func match(n ast.Node, keys []string) []ast.Node {
	if len(keys) == 0 {
		return []ast.Node{}
	}

	var t ast.Expr
	switch n := n.(type) {
	case *ast.TypeSpec:
		t = n.Type
	case *ast.Field:
		t = n.Type
	case *ast.Ident:
		log.Fatalln("to implement 2")
	case *ast.ArrayType:
		t = n.Elt
	default:
		log.Fatalln("unhandled type", reflect.TypeOf(n))
	}

	matches := []ast.Node{}

	switch key := keys[0]; key {
	case "**": // works like Levenshtein (DP)
		if st, ok := t.(*ast.StructType); ok {
			for _, f := range st.Fields.List {
				matches = append(matches, match(f, keys)...)
				matches = append(matches, match(f, degraded(keys))...)
			}
		}

	case "*":
		if st, ok := t.(*ast.StructType); ok {
			for _, f := range st.Fields.List {
				matches = append(matches, match(f, keys[1:])...)
			}
		}

	case "[]": // other selectors like [birthday], [photo,email] are reserved for later use
		if at, ok := t.(*ast.ArrayType); ok {
			if len(keys) == 1 { // should be leaf
				matches = append(matches, at.Elt)
			} else {
				matches = append(matches, match(at, keys[1:])...)
			}
		}

	default:
		if st, ok := t.(*ast.StructType); ok {
			for _, f := range st.Fields.List {
				ckey, err := stripKeyname(f.Tag.Value)
				if err != nil {
					log.Fatalf("could not get the key name out of field tag for %s", f.Tag.Value)
				}
				if ckey == key {
					if len(keys) == 1 { // should be leaf
						matches = append(matches, f)
					} else {
						matches = append(matches, match(f, keys[1:])...)
					}
				}
			}
		}
	}
	return matches
}

// accepts processed form of Config type AST which:
//   - should not have multiple names per ast.Field
//   - array types should be defined by combining compatible item fields
func Match(cfg *ast.TypeSpec, keypath string) []ast.Node {
	keys := strings.Split(keypath, ".")
	if len(keys) == 0 {
		return []ast.Node{}
	} else if l := keys[len(keys)-1]; l == "*" || l == "**" {
		return []ast.Node{}
	}

	cds := match(cfg, keys)

	return cds
}
