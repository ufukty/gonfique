package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/models"
)

type Directives struct {
	b          *bundle.Bundle
	Directives map[models.FlattenKeypath]directivefile.Directives

	Expansions         map[models.WildcardKeypath][]models.FlattenKeypath // matches
	FeatureUsers       featureusers
	Holders            map[models.FlattenKeypath]ast.Node // keypath -> Field, ArrayType (inverse Keypaths)
	Keypaths           map[ast.Node]models.FlattenKeypath // holder -> keypath (resolver)
	NamedTypeExprs     map[models.TypeName]ast.Expr
	NeededToBeDeclared []models.FlattenKeypath
	NeededToBeReferred []models.FlattenKeypath
	TypeExprs          map[models.FlattenKeypath]ast.Expr
	TypenamesAutogen   map[models.FlattenKeypath]models.TypeName
	TypenamesElected   map[models.FlattenKeypath]models.TypeName
	TypenamesProvided  map[models.FlattenKeypath]models.TypeName
	TypenameUsers      map[models.TypeName][]models.FlattenKeypath
}

func New(b *bundle.Bundle) *Directives {
	return &Directives{
		b: b,

		Expansions:         map[models.WildcardKeypath][]models.FlattenKeypath{},
		FeatureUsers:       featureusers{},
		NamedTypeExprs:     map[models.TypeName]ast.Expr{},
		NeededToBeDeclared: []models.FlattenKeypath{},
		NeededToBeReferred: []models.FlattenKeypath{},
		TypeExprs:          map[models.FlattenKeypath]ast.Expr{},
		TypenamesElected:   map[models.FlattenKeypath]models.TypeName{},
		TypenameUsers:      map[models.TypeName][]models.FlattenKeypath{},
	}
}

func (d *Directives) Apply(verbose bool) error {
	d.populateKeypathsAndHolders()

	if err := d.populateExprs(); err != nil {
		return fmt.Errorf("collecting type expressions for each keypaths: %w", err)
	}
	if err := d.expandKeypaths(); err != nil {
		return fmt.Errorf("expanding keypaths: %w", err)
	}
	d.linearizeDirectives()
	d.populateTypenamesAutogen()
	d.populateTypenamesProvided()
	d.populateFeatureUsers()

	if err := d.typenames(); err != nil {
		return fmt.Errorf("listing, declaring typenames and swapping definitions: %w", err)
	}
	if err := d.populateNamedTypeExprs(); err != nil {
		return fmt.Errorf("checking for named directive: %w", err)
	}
	if err := d.checkConflictsForParentRefs(); err != nil {
		return fmt.Errorf("checking conflicts for adding parent refs: %w", err)
	}
	if verbose {
		d.debug()
	}
	d.implementTypeDeclarations()
	if err := d.replaceTypeExpressionsWithIdents(); err != nil {
		return fmt.Errorf("declaring named types: %w", err)
	}
	if err := d.addAccessorFuncDecls(); err != nil {
		return fmt.Errorf("implementing accessor methods: %w", err)
	}
	if err := d.addParentRefs(); err != nil {
		return fmt.Errorf("adding parent refs as fields: %w", err)
	}

	return nil
}
