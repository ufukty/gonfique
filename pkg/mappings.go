package pkg

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
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

func Mappings(cts *ast.TypeSpec, mappings map[Pathway]TypeName) *ast.GenDecl {
	idents := map[TypeName]*ast.Ident{}
	for _, tn := range mappings {
		idents[tn] = ast.NewIdent(tn)
	}

	mis := map[*matchitem]*ast.Ident{}
	for pw, tn := range mappings {
		for _, m := range MatchTypeDefinitionHolder(cts, pw) {
			mis[&m] = idents[tn]
		}
	}

	miskeys := maps.Keys(mis)
	slices.SortFunc(miskeys, func(l, r *matchitem) int {
		if containsPathway(l.pathway, r.pathway) {
			return -1
		} else {
			return +1
		}
	})

	products := map[*ast.Ident]ast.Expr{}
	for _, mi := range miskeys {
		fmt.Println(strings.Join(mi.pathway, "."))
		i := mis[mi]
		switch holder := mi.holder.(type) {
		case *ast.Field:
			if t, ok := products[i]; ok && !compare(t, holder.Type) {
				log.Fatalf("conflicting schemas found for type name %q\n", i.Name)
			}
			products[i] = holder.Type
			holder.Type = i
		case *ast.ArrayType:
			if t, ok := products[i]; ok && !compare(t, holder.Elt) {
				log.Fatalf("conflicting schemas found for type name %q\n", i.Name)
			}
			products[i] = holder.Elt
			holder.Elt = i
		}
	}

	gd := &ast.GenDecl{
		Tok:   token.TYPE,
		Specs: []ast.Spec{},
	}
	for i, t := range products {
		gd.Specs = append(gd.Specs, &ast.TypeSpec{
			Name: i,
			Type: t,
		})
	}
	return gd
}
