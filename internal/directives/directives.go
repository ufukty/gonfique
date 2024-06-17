package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/models"
)

type Directives struct {
	b *bundle.Bundle

	DirectivesForKeypaths map[models.FlattenKeypath]directivefile.Directives
	Expansions            map[models.WildcardKeypath][]models.FlattenKeypath // matches
	FeaturesForKeypaths   perfeature[[]models.FlattenKeypath]
	FeaturesForTypenames  perfeature[[]models.TypeName]
	Holders               map[models.FlattenKeypath]ast.Node // keypath -> Field, ArrayType (inverse Keypaths)
	Keypaths              map[ast.Node]models.FlattenKeypath // holder -> keypath (resolver)
	NeededToBeDeclared    []models.FlattenKeypath
	NeededToBeReferred    []models.FlattenKeypath
	TypeExprs             map[models.FlattenKeypath]ast.Expr
	TypenamesAutogen      map[models.FlattenKeypath]models.TypeName
	TypenamesElected      map[models.FlattenKeypath]models.TypeName
	TypenamesProvided     map[models.FlattenKeypath]models.TypeName
	TypenameUsers         map[models.TypeName][]models.FlattenKeypath
}

func New(b *bundle.Bundle) *Directives {
	return &Directives{
		b: b,

		Expansions:           map[models.WildcardKeypath][]models.FlattenKeypath{},
		FeaturesForKeypaths:  perfeature[[]models.FlattenKeypath]{},
		FeaturesForTypenames: perfeature[[]models.TypeName]{},
		NeededToBeDeclared:   []models.FlattenKeypath{},
		NeededToBeReferred:   []models.FlattenKeypath{},
		TypeExprs:            map[models.FlattenKeypath]ast.Expr{},
		TypenamesElected:     map[models.FlattenKeypath]models.TypeName{},
		TypenameUsers:        map[models.TypeName][]models.FlattenKeypath{},
	}
}

// DONE: pre-type conflicts
// TODO: target group type merging
// FIXME: type assign
// FIXME: post-type conflicts
// TODO: type manipulation (parent, embed)
func (d *Directives) Apply(verbose bool) error {
	d.populateKeypathsAndHolders()
	if err := d.populateExprs(); err != nil {
		return fmt.Errorf("listing types: %w", err)
	}
	if err := d.expandKeypaths(); err != nil {
		return fmt.Errorf("expanding keypaths: %w", err)
	}
	d.populateDirectivesForKeypaths()
	d.populateTypenamesAutogen()
	d.populateTypenamesProvided()
	d.populateFeaturesForKeypaths()
	if err := d.preTypeConflicts(); err != nil {
		return fmt.Errorf("pre-type conflict checking: %w", err)
	}
	d.checkTypenameRequirements()
	if err := d.typenameElection(); err != nil {
		return fmt.Errorf("typename election: %w", err)
	}
	d.populateFeaturesForTypenames()
	if err := d.postTypeConflicts(); err != nil {
		return fmt.Errorf("post-type conflict checking: %w", err)
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
