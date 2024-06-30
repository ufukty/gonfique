package directives

import (
	"fmt"
	"go/ast"

	"github.com/ufukty/gonfique/internal/models"
)

func (d *Directives) populateTypeExprs() {
	for tn, kps := range d.TypenameUsers {
		modelkp := kps[0]
		modelty := d.KeypathTypeExprs[modelkp]
		d.TypeExprs[tn] = modelty
	}
}
type typenameParentDetails struct {
	Fieldname  models.FieldName
	ParentType models.TypeName
}

type typenameEmbedDetails struct {
	EmbeddedTypeName       models.TypeName
	EmbeddedTypeImportPath string
	EmbeddedTypeImportAs   string
	EmbeddedTypeFieldList  *ast.FieldList
}

type typenameAccessorsDetails struct {
	FieldsAndTypes map[models.FieldName]models.TypeName
}

type parametersForTypenames struct {
	// Declare
	// Replace
	// Export
	Accessors map[models.TypeName]typenameAccessorsDetails
	Embed     map[models.TypeName]typenameEmbedDetails
	Parent    map[models.TypeName]typenameParentDetails
}

// TODO: move error checking into post-type conflict checking; improve them for more helpful logs
func (d *Directives) mergeDirectiveParametersForTypes() error {
	f := parametersForTypenames{
		Accessors: map[models.TypeName]typenameAccessorsDetails{},
		Embed:     map[models.TypeName]typenameEmbedDetails{},
		Parent:    map[models.TypeName]typenameParentDetails{},
	}

	for tn, kps := range d.TypenameUsers {
		init := true
		f.Accessors[tn] = typenameAccessorsDetails{
			FieldsAndTypes: map[models.FieldName]models.TypeName{},
		}
		for _, kp := range kps {
			for _, fn := range d.DirectivesForKeypaths[kp].Accessors {
				fkp := kp.WithFieldPath(fn)
				ftn := d.TypenamesElected[fkp]
				fn := d.b.Fieldnames[d.Holders[fkp]]
				current := f.Accessors[tn]
				if !init {
					if current.FieldsAndTypes[fn] != ftn {
						return fmt.Errorf("typename %q is directed to have accessors on the field %q  which its type resolving to different types", tn, fn)
					}
				}
				f.Accessors[tn].FieldsAndTypes[fn] = ftn
				init = true
			}
		}
	}

	// FIXME: use "types" or "ast" package to inspect field list of specified embedding struct
	for tn, kps := range d.TypenameUsers {
		f.Embed[tn] = typenameEmbedDetails{
			EmbeddedTypeImportPath: "",
			EmbeddedTypeName:       "",
			EmbeddedTypeFieldList:  nil,
		}
		for _, kp := range kps {
			dirs := d.DirectivesForKeypaths[kp]
			if dirs.Embed.Typename != "" {
				dirs := d.DirectivesForKeypaths[kp].Embed
				current := f.Embed[tn]
				// if dirs.Typename != current.EmbeddedTypeName {
				//   return fmt.Errorf("typename %q is directed to embed different types %q and %q", tn, dirs.Typename, current.EmbeddedTypeName)
				// }
				// if dirs.ImportPath != current.EmbeddedTypeImportPath {
				//   return fmt.Errorf("typename %q is directed to embed types from different imports %q and %q", tn, dirs.ImportPath, current.EmbeddedTypeImportPath)
				// }
				// if dirs.ImportAs != current.EmbeddedTypeImportAs {
				//   return fmt.Errorf("typename %q is directed to embed types from imports with different aliases %q and %q", tn, dirs.ImportAs, current.EmbeddedTypeImportPath)
				// }
				current.EmbeddedTypeName = dirs.Typename
				current.EmbeddedTypeImportPath = dirs.ImportPath
				current.EmbeddedTypeImportAs = dirs.ImportAs
				f.Embed[tn] = current
			}
		}
	}

	for tn, kps := range d.TypenameUsers {
		for _, kp := range kps {
			dirs := d.DirectivesForKeypaths[kp]
			if dirs.Parent != "" {
				current, ok := f.Parent[tn]
				if !ok {
					f.Parent[tn] = typenameParentDetails{
						Fieldname:  "",
						ParentType: "",
					}
				}
				// if dirs.Parent.Fieldname != current.Fieldname {
				//   return fmt.Errorf("typename %q is ...", tn)
				// }
				// if dirs.Parent.Accessors != current.Accessors {
				//   return fmt.Errorf("typename %q is ...", tn)
				// }
				// if dirs.Parent.Level != current.Level {
				//   return fmt.Errorf("typename %q is ...", tn)
				// }
				current.Fieldname = dirs.Parent
				current.ParentType = d.TypenamesElected[kp.Parent()] // TODO: compare with current value
				f.Parent[tn] = current
			}
		}
	}

	d.ParametersForTypenames = f
	return nil
}
