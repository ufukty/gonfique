package directives

// import (
// 	"go/ast"

// 	"github.com/ufukty/gonfique/internal/paths/models"
// )

// type embeddingDetails struct {
// 	EmbeddedTypeName       models.TypeName
// 	EmbeddedTypeImportPath string
// 	EmbeddedTypeImportAs   string
// 	EmbeddedTypeFieldList  *ast.FieldList
// }

// func (d *Directives) getEmbeddingDetails(users map[models.TypeName][]models.FlattenKeypath) (map[models.TypeName]embeddingDetails, error) {
// 	details := map[models.TypeName]embeddingDetails{}

// 	// FIXME: use "types" or "ast" package to inspect field list of specified embedding struct
// 	for tn, kps := range users {
// 		details[tn] = embeddingDetails{
// 			EmbeddedTypeImportPath: "",
// 			EmbeddedTypeName:       "",
// 			EmbeddedTypeFieldList:  nil,
// 		}
// 		for _, kp := range kps {
// 			dirs := d.DirectivesForKeypaths[kp]
// 			if dirs.Embed.Typename != "" {
// 				dirs := d.DirectivesForKeypaths[kp].Embed
// 				current := details[tn]
// 				current.EmbeddedTypeName = dirs.Typename
// 				current.EmbeddedTypeImportPath = dirs.ImportPath
// 				current.EmbeddedTypeImportAs = dirs.ImportAs
// 				details[tn] = current
// 			}
// 		}
// 	}

// 	return details, nil
// }
