package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/models"
)

type featuresForKeypaths struct {
	Accessors []models.FlattenKeypath
	Embed     []models.FlattenKeypath
	Export    []models.FlattenKeypath
	Import    []models.FlattenKeypath
	Declare   []models.FlattenKeypath
	Parent    []models.FlattenKeypath
	Replace   []models.FlattenKeypath
}

type Directives struct {
	b *bundle.Bundle

	Keypaths   map[ast.Node]models.FlattenKeypath // holder (Field, ArrayType etc.) -> keypath (resolver)
	Holders    map[models.FlattenKeypath]ast.Node // inverse Keypaths
	Exprs      map[models.FlattenKeypath]ast.Expr
	Expansions map[models.WildcardKeypath][]models.FlattenKeypath
	Sources    map[models.FlattenKeypath]parameterSources
	Directives map[models.FlattenKeypath]directivefile.Directives // flatten Sources
	Features   featuresForKeypaths                                // convenience
	ToRefer    []models.FlattenKeypath
	Elected    map[models.FlattenKeypath]models.TypeName
	Instances  map[models.TypeName][]models.FlattenKeypath // inverse Elected
}

func New(b *bundle.Bundle) *Directives {
	return &Directives{
		b:          b,
		Expansions: map[models.WildcardKeypath][]models.FlattenKeypath{},
		Features:   featuresForKeypaths{},
		ToRefer:    []models.FlattenKeypath{},
		Exprs:      map[models.FlattenKeypath]ast.Expr{},
		Elected:    map[models.FlattenKeypath]models.TypeName{},
		Instances:  map[models.TypeName][]models.FlattenKeypath{},
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
	d.populateDirectivesAndFeaturesForKeypaths()
	if err := d.checkPreTypeConflicts(); err != nil {
		return fmt.Errorf("pre-type conflict checking: %w", err)
	}
	d.checkKeypathsToReferTheirType()
	if err := d.typenameElection(); err != nil {
		return fmt.Errorf("typename election: %w", err)
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
	if err := d.implementParentRefs(); err != nil {
		return fmt.Errorf("implementing fields for parent refs: %w", err)
	}
	return nil
}
