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

	DirectivesForKeypaths  map[models.FlattenKeypath]directivefile.Directives
	Expansions             map[models.WildcardKeypath][]models.FlattenKeypath // matches
	FeaturesForKeypaths    featuresForKeypaths
	ParametersForTypenames parametersForTypenames // in-effect directive parameters
	ParameterSources       map[models.FlattenKeypath]parameterSources
	Holders                map[models.FlattenKeypath]ast.Node // keypath -> Field, ArrayType (inverse Keypaths)
	Keypaths               map[ast.Node]models.FlattenKeypath // holder -> keypath (resolver)
	NeededToBeDeclared     []models.FlattenKeypath
	NeededToBeReferred     []models.FlattenKeypath
	KeypathTypeExprs       map[models.FlattenKeypath]ast.Expr
	TypeExprs              map[models.TypeName]ast.Expr
	TypenamesElected       map[models.FlattenKeypath]models.TypeName
	TypenameUsers          map[models.TypeName][]models.FlattenKeypath // inverse TypenamesElected
}

func New(b *bundle.Bundle) *Directives {
	return &Directives{
		b:                      b,
		Expansions:             map[models.WildcardKeypath][]models.FlattenKeypath{},
		FeaturesForKeypaths:    featuresForKeypaths{},
		ParametersForTypenames: parametersForTypenames{},
		NeededToBeDeclared:     []models.FlattenKeypath{},
		NeededToBeReferred:     []models.FlattenKeypath{},
		KeypathTypeExprs:       map[models.FlattenKeypath]ast.Expr{},
		TypeExprs:              map[models.TypeName]ast.Expr{},
		TypenamesElected:       map[models.FlattenKeypath]models.TypeName{},
		TypenameUsers:          map[models.TypeName][]models.FlattenKeypath{},
	}
}

// DONE: check conflicting rules & directives
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
	d.parameterSourceClassification()
	if err := d.checkConflictingSources(); err != nil {
		return fmt.Errorf("comparing directives applied to same targets: %w", err)
	}
	d.populateDirectivesForKeypaths()
	d.populateFeaturesForKeypaths()
	if err := d.checkPreTypeConflicts(); err != nil {
		return fmt.Errorf("pre-type conflict checking: %w", err)
	}
	d.checkKeypathsToReferTheirType()
	if err := d.typenameElection(); err != nil {
		return fmt.Errorf("typename election: %w", err)
	}
	d.populateTypeExprs()
	if err := d.mergeDirectiveParametersForTypes(); err != nil {
		return fmt.Errorf("merging directives defined on multiple paths of same type: %w", err)
	}
	if err := d.checkPostTypeConflicts(); err != nil {
		return fmt.Errorf("post-type conflict checking: %w", err)
	}
	if verbose {
		d.debug()
	}
	d.checkKeypathsToModifyTheirType()
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
