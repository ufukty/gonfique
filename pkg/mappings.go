package pkg

import (
	"fmt"
	"go/ast"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Pathway = string
type TypeName = string

func ReadMappings(src string) (map[Pathway]TypeName, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	ms := map[Pathway]TypeName{}
	err = yaml.NewDecoder(f).Decode(&ms)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}

	return ms, nil
}

func Mappings(cts *ast.TypeSpec, mappings map[Pathway]TypeName) map[*ast.Ident]ast.Expr {
	produced := map[*ast.Ident]ast.Expr{}
	for pw, tn := range mappings {
		i := ast.NewIdent(tn)
		for _, n := range Match(cts, pw) {
			switch n := n.(type) {
			case *ast.Field:
				if t, ok := produced[i]; ok && !compare(t, n.Type) {
					log.Fatalf("conflicting schemas found for type name %q\n", tn)
				}
				produced[i] = n.Type
				n.Type = i
			case *ast.ArrayType:
				if t, ok := produced[i]; ok && !compare(t, n.Elt) {
					log.Fatalf("conflicting schemas found for type name %q\n", tn)
				}
				produced[i] = n.Elt
				n.Elt = i
			}
		}
	}
	return produced
}
