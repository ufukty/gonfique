package mappings

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"

	"github.com/ufukty/gonfique/internal/compares"
	"github.com/ufukty/gonfique/internal/files"
	"gopkg.in/yaml.v3"
)

type TypeName = string

func ReadMappings(src string) (map[files.Keypath]TypeName, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	ms := map[files.Keypath]TypeName{}
	err = yaml.NewDecoder(f).Decode(&ms)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}

	return ms, nil
}

func ApplyMappings(f *files.File, mappings map[files.Keypath]TypeName) error {
	matchlists := map[*ast.Ident][]ast.Node{}
	for kp, tn := range mappings {
		matches, err := matchTypeDefHolder(f.Cfg, kp, f.OriginalKeys)
		if err != nil {
			return fmt.Errorf("matching the rule: %w", err)
		}
		if len(matches) == 0 {
			fmt.Printf("Pattern %q (->%s) didn't match any region\n", kp, tn)
		}
		matchlists[ast.NewIdent(tn)] = matches
	}

	products := map[*ast.Ident]ast.Expr{}
	for i, matchlist := range matchlists {
		for _, match := range matchlist {

			switch holder := match.(type) {
			case *ast.Field:
				if t, ok := products[i]; ok {
					if !compares.Compare(t, holder.Type) {
						return fmt.Errorf("conflicting schemas found for type name %q", i.Name)
					}
				} else {
					products[i] = holder.Type
				}
				holder.Type = i
			case *ast.ArrayType:
				if t, ok := products[i]; ok {
					if !compares.Compare(t, holder.Elt) {
						return fmt.Errorf("conflicting schemas found for type name %q", i.Name)
					}
				} else {
					products[i] = holder.Elt
				}
				holder.Elt = i
			}
		}
	}

	for i, t := range products {
		f.Named = append(f.Named, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{
				Name: i,
				Type: t,
			}},
		})
	}

	return nil
}
