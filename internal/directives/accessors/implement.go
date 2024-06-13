package accessors

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

func caseInsensitiveCompareTypenames(a, b models.TypeName) int {
	if strings.ToLower(string(a)) < strings.ToLower(string(b)) {
		return -1
	} else {
		return +1
	}
}

func caseInsensitiveCompareFieldnames(a, b models.FieldName) int {
	if strings.ToLower(string(a)) < strings.ToLower(string(b)) {
		return -1
	} else {
		return +1
	}
}

func Implement(b *bundle.Bundle) error {
	if b.Df == nil {
		return fmt.Errorf("directive file is not populated")
	} else if b.ElectedTypenames == nil {
		return fmt.Errorf("elected type names are missing")
	}
	b.Accessors = []*ast.FuncDecl{}

	fieldsfortypes := map[models.TypeName]map[models.FieldName]models.TypeName{}
	for wckp, directives := range *b.Df {
		if directives.Accessors != nil {
			for _, kp := range b.Expansions[wckp] {
				tn := b.ElectedTypenames[kp]
				if _, ok := fieldsfortypes[tn]; !ok {
					fieldsfortypes[tn] = map[models.FieldName]models.TypeName{}
				}
				for _, fp := range directives.Accessors {
					fkp := kp.WithFieldPath(fp)
					ftn := b.ElectedTypenames[fkp]
					fn := b.Fieldnames[b.Holders[fkp]]
					fieldsfortypes[tn][fn] = ftn
				}
			}
		}
	}

	sorted := maps.Keys(fieldsfortypes)
	slices.SortFunc(sorted, caseInsensitiveCompareTypenames)
	for _, tn := range sorted {
		fields := fieldsfortypes[tn]
		sortedfields := maps.Keys(fields)
		slices.SortFunc(sortedfields, caseInsensitiveCompareFieldnames)
		for _, fn := range sortedfields {
			ftn := fields[fn]
			b.Accessors = append(b.Accessors,
				generateGetter(tn, fn, ftn),
				generateSetter(tn, fn, ftn),
			)
		}
	}

	return nil
}
