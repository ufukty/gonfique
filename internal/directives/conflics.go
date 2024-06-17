package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/compares"
)

// func CheckConflictingDirectives(b *bundle.Bundle) error {
// 	kps := map[ast.Node][]models.WildcardKeypath{}
// 	for kp, matches := range b.Expansions {
// 		for _, match := range matches {
// 			if _, ok := kps[match]; !ok {
// 				kps[match] = []models.WildcardKeypath{}
// 			}
// 			kps[match] = append(kps[match], kp)
// 		}
// 	}

// 	keypaths := (b.Expansions)
// }

func (d *Directives) compareMergedTypenameUsers() error {
	for tn, kps := range d.TypenameUsers {
		for i := 1; i < len(kps); i++ {
			if !compares.Compare(d.TypeExprs[kps[0]], d.TypeExprs[kps[i]]) {
				return fmt.Errorf("%q and %q doesn't share the same schema, but required to share same type %q", kps[0], kps[i], tn)
			}
		}
	}
	return nil
}
