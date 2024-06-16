package directives

import (
	"fmt"
	"strings"

	"github.com/ufukty/gonfique/internal/compares"
	"github.com/ufukty/gonfique/internal/models"
)

func caseInsensitiveCompareTypenames(a, b models.TypeName) int {
	if strings.ToLower(string(a)) < strings.ToLower(string(b)) {
		return -1
	} else {
		return +1
	}
}

func caseInsensitiveCompareFieldnames(a, b models.FieldName) int {
	if strings.ToLower(string(a)) < strings.ToLower(string(b)) {
		return -1
	} else {
		return +1
	}
}

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
