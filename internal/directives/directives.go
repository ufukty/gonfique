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

	Keypaths       map[ast.Node]models.FlattenKeypath // holder -> keypath (resolver)
	Holders        map[models.FlattenKeypath]ast.Node // keypath -> Field, ArrayType (inverse Keypaths)
	TypeExprs      map[models.FlattenKeypath]ast.Expr
	NamedTypeExprs map[models.TypeName]ast.Expr
	Expansions     map[models.WildcardKeypath][]models.FlattenKeypath // matches

	NeededToBeReferred []models.FlattenKeypath
	NeededToBeDeclared []models.FlattenKeypath
	ElectedTypenames   map[models.FlattenKeypath]models.TypeName
	TypenameUsers      map[models.TypeName][]models.FlattenKeypath
}

func New(b *bundle.Bundle) *Directives {
	return &Directives{
		b: b,

		NeededToBeReferred: []models.FlattenKeypath{},
		NeededToBeDeclared: []models.FlattenKeypath{},

		ElectedTypenames: map[models.FlattenKeypath]models.TypeName{},
		TypenameUsers:    map[models.TypeName][]models.FlattenKeypath{},

		Expansions:     map[models.WildcardKeypath][]models.FlattenKeypath{},
		TypeExprs:      map[models.FlattenKeypath]ast.Expr{},
		NamedTypeExprs: map[models.TypeName]ast.Expr{},
	}
}

func (d *Directives) Apply(verbose bool) error {
	d.populateKeypathsAndHolders()

	if err := d.populateExprs(); err != nil {
		return fmt.Errorf("collecting type expressions for each keypaths: %w", err)
	}
	if err := d.expandKeypathsInDirectives(); err != nil {
		return fmt.Errorf("expanding keypaths: %w", err)
	}
	d.linearizeDirectives()

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
	if err := d.AddAccessorFuncDecls(); err != nil {
		return fmt.Errorf("implementing accessor methods: %w", err)
	}
	if err := d.addParentRefs(); err != nil {
		return fmt.Errorf("adding parent refs as fields: %w", err)
	}

	return nil
}
