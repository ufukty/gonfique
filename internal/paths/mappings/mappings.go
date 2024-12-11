package mappings

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/compares"
	"github.com/ufukty/gonfique/internal/matcher"
	"github.com/ufukty/gonfique/internal/paths/models"
)

func ApplyMappings(b *bundle.Bundle, mappings map[models.WildcardKeypath]models.TypeName) error {
	matchlists := map[*ast.Ident][]ast.Node{}
	for kp, tn := range mappings {
		matches, err := matcher.FindTypeDefHoldersForKeypath(b.CfgType, kp, b.OriginalKeys)
		if err != nil {
			return fmt.Errorf("matching the rule: %w", err)
		}
		if len(matches) == 0 {
			fmt.Printf("Pattern %q (->%s) didn't match any region\n", kp, tn)
		}
		matchlists[tn.Ident()] = matches
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
		b.Named = append(b.Named, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{
				Name: i,
				Type: t,
			}},
		})
	}

	return nil
}
