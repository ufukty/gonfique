package directives

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/compares"
	"github.com/ufukty/gonfique/internal/models"
)

func (d *Directives) populateNamedTypeExprs() error {
	typenames := map[models.TypeName][]models.FlattenKeypath{}
	for wckp, dirs := range *d.b.Df {
		if dirs.Named != "" {
			tn := dirs.Named
			for _, kp := range d.Expansions[wckp] {
				if _, ok := typenames[tn]; !ok {
					typenames[tn] = []models.FlattenKeypath{}
				}
				typenames[tn] = append(typenames[tn], kp)
			}
		}
	}

	for tn, kps := range typenames {
		for i := 1; i < len(kps); i++ {
			if !compares.Compare(d.TypeExprs[kps[0]], d.TypeExprs[kps[i]]) {
				return fmt.Errorf("can't use same type %q for %q and %q", tn, kps[0], kps[i])
			}
			for _, kp := range kps {
				d.TypenamesElected[kp] = tn
			}
			d.NamedTypeExprs[tn] = d.TypeExprs[kps[0]]
		}
	}

	return nil
}
