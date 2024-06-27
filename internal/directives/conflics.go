package directives

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/compares"
	"github.com/ufukty/gonfique/internal/datas"
	"github.com/ufukty/gonfique/internal/models"
)

// TODO: Does it makes sense to add checks for accessors, export etc.
func (d *Directives) checkConflictingSources() error {
	conflicts := []string{}
	for kp, directivesources := range d.ParameterSources {
		if len(directivesources.Declare) > 1 {
			msg := fmt.Sprintf("%s: conflicting declare parameters:", kp)
			for val, wckps := range directivesources.Declare {
				msg += fmt.Sprintf("\n  %v => %q", wckps, val)
			}
			conflicts = append(conflicts, msg)
		}
		if len(directivesources.Parent) > 1 {
			msg := fmt.Sprintf("%s: conflicting parent parameters:", kp)
			for val, wckps := range directivesources.Parent {
				msg += fmt.Sprintf("\n  %v => %v", wckps, val)
			}
			conflicts = append(conflicts, msg)
		}
		if len(directivesources.Replace) > 1 {
			msg := fmt.Sprintf("%s: conflicting replace parameters:", kp)
			for val, wckps := range directivesources.Replace {
				msg += fmt.Sprintf("\n  %v => %v", wckps, val)
			}
			conflicts = append(conflicts, msg)
		}
	}
	if len(conflicts) > 0 {
		return fmt.Errorf(strings.Join(conflicts, "\n"))
	}
	return nil
}

func (d *Directives) checkPreTypeConflicts() error {
	conflicts := []string{}

	for tn, kps := range d.TypenameUsers {
		for i := 1; i < len(kps); i++ {
			if !compares.Compare(d.KeypathTypeExprs[kps[0]], d.KeypathTypeExprs[kps[i]]) {
				conflicts = append(conflicts, fmt.Sprintf("%s: typename is used for 2 targets with conflicting schemas: %s, %s", tn, kps[0], kps[i]))
			}
		}
	}

	for kp := range d.DirectivesForKeypaths {
		for pkp := kp.Parent(); pkp != ""; pkp = pkp.Parent() {
			if d.DirectivesForKeypaths[pkp].Replace.Typename != "" {
				conflicts = append(conflicts, fmt.Sprintf("%q: directive for unmanaged subtree (caused by type of %s is manually assigned to %s)", kp, pkp, d.DirectivesForKeypaths[pkp].Replace))
			}
		}
	}

	for _, user := range d.FeaturesForKeypaths.Parent {
		if _, ok := d.KeypathTypeExprs[user].(*ast.StructType); !ok {
			conflicts = append(conflicts, fmt.Sprintf("%s: non-dict target for parent directive", user))
		}
	}

	if len(conflicts) > 0 {
		slices.Sort(conflicts)
		return fmt.Errorf("found %d conflicts:\n%s", len(conflicts), strings.Join(conflicts, "\n"))
	}
	return nil
}

func (d *Directives) checkPostTypeConflicts() error {
	conflicts := []string{}

	for tn := range d.ParametersForTypenames.Parent {
		parents := []models.TypeName{}
		for _, user := range d.TypenameUsers[tn] {
			parents = append(parents, d.TypenamesElected[user.Parent()])
		}
		simplified := datas.Uniq(parents)
		if len(simplified) > 1 {
			return fmt.Errorf("users of type %q have parents with different types: %v", tn, simplified)
		}
	}

	for _, user := range d.FeaturesForKeypaths.Parent {
		if user == "" {
			return fmt.Errorf("")
		}
	}

	if len(conflicts) > 0 {
		slices.Sort(conflicts)
		return fmt.Errorf("found %d conflicts:\n%s", len(conflicts), strings.Join(conflicts, "\n"))
	}
	return nil
}
