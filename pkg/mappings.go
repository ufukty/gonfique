package pkg

import (
	"fmt"
	"go/ast"
	"go/token"
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

func Mappings(cts *ast.TypeSpec, mappings map[Pathway]TypeName) *ast.GenDecl {
	mss := map[*ast.Ident][]ast.Node{}
	for pw, tn := range mappings {
		i := ast.NewIdent(tn)
		fmt.Printf("%-30s %s\n", pw, tn)
		mss[i] = Match(cts, pw)
	}

	produced := map[*ast.Ident]ast.Expr{}
	for i, ms := range mss {
		for _, n := range ms {
			fmt.Println("> match")
			switch n := n.(type) {
			case *ast.Field:
				if t, ok := produced[i]; ok && !compare(t, n.Type) {
					log.Fatalf("conflicting schemas found for type name %q\n", i.Name)
				}
				produced[i] = n.Type
				n.Type = i
			case *ast.ArrayType:
				if t, ok := produced[i]; ok && !compare(t, n.Elt) {
					log.Fatalf("conflicting schemas found for type name %q\n", i.Name)
				}
				produced[i] = n.Elt
				n.Elt = i
			}
		}
	}

	fmt.Println("|||")
	gd := &ast.GenDecl{
		Tok:   token.TYPE,
		Specs: []ast.Spec{},
	}
	for i, t := range produced {
		gd.Specs = append(gd.Specs, &ast.TypeSpec{
			Name: i,
			Type: t,
		})
	}
	return gd
}
