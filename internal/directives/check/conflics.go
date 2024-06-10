package check

import (
	"fmt"
	"go/ast"
	"reflect"
	"slices"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

func neededTypesForAccessorsDirective(b *bundle.Bundle) ([]models.FlattenKeypath, error) {
	needed := []models.FlattenKeypath{}
	for wckp, drs := range *b.Df {
		if drs.Accessors != nil {
			matches, ok := b.Expansions[wckp]
			if !ok {
				return nil, fmt.Errorf("expansions are not found for wildcard containing keypath: %s", wckp)
			}
			for _, match := range matches {
				if _, ok := match.(*ast.Ident); ok {
					continue
				}
				kp, ok := b.Keypaths[match]
				if !ok {
					return nil, fmt.Errorf("flatten keypath is not found for wildcard containing keypath: %s", kp)
				}
				needed = append(needed, kp) // struct
				for _, field := range drs.Accessors {
					fkp := kp.WithFieldPath(field)
					if _, ok := b.TypeExprs[fkp].(*ast.Ident); ok {
						continue
					}
					needed = append(needed, fkp) // its field
				}
			}
		}
	}
	return needed, nil
}

func neededTypesForParentDirective(b *bundle.Bundle) ([]models.FlattenKeypath, error) {
	needed := []models.FlattenKeypath{}
	for wckp, drs := range *b.Df {
		if drs.Parent != "" {
			matches, ok := b.Expansions[wckp]
			if !ok {
				return nil, fmt.Errorf("expansions are not found for wildcard containing keypath: %s", wckp)
			}
			for _, match := range matches {
				kp, ok := b.Keypaths[match]
				if !ok {
					return nil, fmt.Errorf("flatten keypath is not found for wildcard containing keypath: %s", kp)
				}
				needed = append(needed, kp.Parent())
			}
		}
	}
	return needed, nil
}

// both the struct and field types at each directive needs to be declared as named (not inline)
func MarkNeededNamedTypes(b *bundle.Bundle) error {
	accessors, err := neededTypesForAccessorsDirective(b)
	if err != nil {
		return fmt.Errorf("accessors: %w", err)
	}
	parent, err := neededTypesForParentDirective(b)
	if err != nil {
		return fmt.Errorf("parent: %w", err)
	}
	b.NeedsToBeNamed = datas.Uniq(slices.Concat(
		accessors,
		parent,
	))
	return nil
}

// func CheckConflictingDirectives(b *bundle.Bundle) error {
// 	kps := map[ast.Node][]models.WildcardKeypath{}
// 	for kp, matches := range b.Expansions {
// 		for _, match := range matches {
// 			if _, ok := kps[match]; !ok {
// 				kps[match] = []models.WildcardKeypath{}
// 			}
// 			kps[match] = append(kps[match], kp)
// 		}
// 	}

// 	keypaths := (b.Expansions)
// }

func PopulateExprs(b *bundle.Bundle) error {
	for kp, n := range b.Holders {
		switch n := n.(type) {
		case *ast.Field:
			b.TypeExprs[kp] = n.Type
		case *ast.ArrayType:
			b.TypeExprs[kp] = n.Elt
		default:
			return fmt.Errorf("unrecognized holder type: %s", reflect.TypeOf(n).String())
		}
	}
	return nil
}
