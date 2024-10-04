package directives

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/compares"
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
				conflicts = append(conflicts, fmt.Sprintf("  declare type for incompatible targets: (%s, %s) => %s", kps[0], kps[i], tn))
			}
		}
	}

	for kp := range d.DirectivesForKeypaths {
		for pkp := kp.Parent(); pkp != ""; pkp = pkp.Parent() {
			if d.DirectivesForKeypaths[pkp].Replace.Typename != "" {
				conflicts = append(conflicts, fmt.Sprintf("  directive for unmanaged subtree: %s (type of %s is replaced with %s)", kp, pkp, d.DirectivesForKeypaths[pkp].Replace))
			}
		}
	}

	for _, user := range d.FeaturesForKeypaths.Parent {
		if len(user.Segments()) <= 1 {
			conflicts = append(conflicts, fmt.Sprintf("  parent ref on top node: %s", user))
		}
		if _, ok := d.KeypathTypeExprs[user].(*ast.StructType); !ok {
			conflicts = append(conflicts, fmt.Sprintf("  non-dict target for parent directive: %s", user))
		}
	}

	if len(conflicts) > 0 {
		slices.Sort(conflicts)
		return fmt.Errorf("found %d conflicts:\n%s", len(conflicts), strings.Join(conflicts, "\n"))
	}
	return nil
}
