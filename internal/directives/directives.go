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

	keypaths   map[ast.Node]models.FlattenKeypath // holder (Field, ArrayType etc.) -> keypath (resolver)
	holders    map[models.FlattenKeypath]ast.Node // inverse Keypaths
	exprs      map[models.FlattenKeypath]ast.Expr
	expansions map[models.WildcardKeypath][]models.FlattenKeypath
	sources    map[models.FlattenKeypath]parameterSources
	directives map[models.FlattenKeypath]directivefile.Directives // flatten Sources
	features   featuresForKeypaths                                // convenience
	toRefer    []models.FlattenKeypath
	elected    map[models.FlattenKeypath]models.TypeName
	instances  map[models.TypeName][]models.FlattenKeypath // inverse Elected
}

func New(b *bundle.Bundle) *Directives {
	return &Directives{
		b:          b,
		expansions: map[models.WildcardKeypath][]models.FlattenKeypath{},
		features:   featuresForKeypaths{},
		toRefer:    []models.FlattenKeypath{},
		exprs:      map[models.FlattenKeypath]ast.Expr{},
		elected:    map[models.FlattenKeypath]models.TypeName{},
		instances:  map[models.TypeName][]models.FlattenKeypath{},
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
