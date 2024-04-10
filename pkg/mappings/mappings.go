package mappings

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"

	"github.com/ufukty/gonfique/pkg/compares"
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

func Mappings(cts *ast.TypeSpec, mappings map[Pathway]TypeName) []*ast.GenDecl {
	miss := map[*ast.Ident][]matchitem{}
	for pw, tn := range mappings {
		matches := MatchTypeDefinitionHolder(cts, pw)
		if len(matches) == 0 {
			fmt.Printf("Pattern %q (->%s) didn't match any region\n", pw, tn)
		}
		miss[ast.NewIdent(tn)] = matches
	}

	products := map[*ast.Ident]ast.Expr{}
	for i, mis := range miss {
		for _, mi := range mis {
			switch holder := mi.holder.(type) {
			case *ast.Field:
				if t, ok := products[i]; ok {
					if !compares.Compare(t, holder.Type) {
						log.Fatalf("conflicting schemas found for type name %q\n", i.Name)
					}
				} else {
					products[i] = holder.Type
				}
				holder.Type = i
			case *ast.ArrayType:
				if t, ok := products[i]; ok {
					if !compares.Compare(t, holder.Elt) {
						log.Fatalf("conflicting schemas found for type name %q\n", i.Name)
					}
				} else {
					products[i] = holder.Elt
				}
				holder.Elt = i
			}
		}
	}

	gds := []*ast.GenDecl{}
	for i, t := range products {
		gds = append(gds, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{
				Name: i,
				Type: t,
			}},
		})
	}
	return gds
}
