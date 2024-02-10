package pkg

import (
	"fmt"
	"go/ast"
	"log"
	"reflect"
	"slices"
	"strings"
)

var ErrNoResult = fmt.Errorf("no result")

type candidate struct {
	i *ast.Ident
	v reflect.Value // to replace its place in AST
}

// segms <- keypath segments
func match(t ast.Expr, keys []string, anc []string) []candidate {
	if len(keys) == 0 {
		return []candidate{}
	}
	replaced := []candidate{}
	switch t := t.(type) {
	case *ast.StructType:
		switch key := keys[0]; key {
		case "**":
			log.Fatalln("to implement")
		case "*":
			for _, f := range t.Fields.List {
				replaced = append(replaced,
					match(f, keys[1:], append(slices.Clone(anc), f.Names[0].Name))...,
				)
			}
		default:
			for _, f := range t.Fields.List {
				kn, err := stripKeyname(f.Tag.Value)
				if err != nil {
					log.Fatalf("could not get the key name out of field tag for %s", f.Tag.Value)
				}
				if kn == key {

				}
				replaced = append(replaced, candidate{}) // f.Type
			}

		}

	case *ast.ArrayType:
		if keys[0] != "[*]" {

		}
		log.Fatalln("to implement")

	default:
		if len(keys) != 1 { // no match
			return []candidate{}
		}

	}
	return replaced
}

func MatchAndReplace(cfg *ast.TypeSpec, keypath string) []string {
	keys := strings.Split(keypath, ".")
	if len(keys) == 0 {
		return []string{}
	} else if l := keys[len(keys)-1]; l == "*" || l == "[*]" {
		return []string{}
	}

	cds := match(cfg.Type, keys, []string{})

	for _, cd := range cds {
		fmt.Println(cd)
	}

	return nil
}
