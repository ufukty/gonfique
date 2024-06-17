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

func (d *Directives) preTypeConflicts() error {
	conflicts := []string{}

	for tn, kps := range d.TypenameUsers {
		for i := 1; i < len(kps); i++ {
			if !compares.Compare(d.TypeExprs[kps[0]], d.TypeExprs[kps[i]]) {
				conflicts = append(conflicts, fmt.Sprintf("%s: typename is used for 2 targets with conflicting schemas: %s, %s", tn, kps[0], kps[i]))
			}
		}
	}

	for kp := range d.DirectivesForKeypaths {
		for pkp := kp.Parent(); pkp != ""; pkp = pkp.Parent() {
			if d.DirectivesForKeypaths[pkp].Type != "" {
				conflicts = append(conflicts, fmt.Sprintf("%q: directive for unmanaged subtree (caused by type of %s is manually assigned to %s)", kp, pkp, d.DirectivesForKeypaths[pkp].Type))
			}
		}
	}

	for _, user := range d.FeaturesForKeypaths.Parent {
		if _, ok := d.TypeExprs[user].(*ast.StructType); !ok {
			conflicts = append(conflicts, fmt.Sprintf("%s: non-dict target for parent directive", user))
		}
	}

	if len(conflicts) > 0 {
		slices.Sort(conflicts)
		return fmt.Errorf("found %d conflicts:\n%s", len(conflicts), strings.Join(conflicts, "\n"))
	}
	return nil
}

func (d *Directives) postTypeConflicts() error {
	conflicts := []string{}

	for _, tn := range d.FeaturesForTypenames.Named {
		kps := d.TypenameUsers[tn]
		for i := 1; i < len(kps); i++ {
			if !compares.Compare(d.TypeExprs[kps[0]], d.TypeExprs[kps[i]]) {
				return fmt.Errorf("can't use same type %q for %q and %q", tn, kps[0], kps[i])
			}
		}
	}

	for _, tn := range d.FeaturesForTypenames.Parent {
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
