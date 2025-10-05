package embeds

// import (
// 	"go/ast"

// 	"go.ufukty.com/gonfique/internal/paths/models"
// )

// type embeddingDetails struct {
// 	EmbeddedTypeName       config.Typename
// 	EmbeddedTypeImportPath string
// 	EmbeddedTypeImportAs   string
// 	EmbeddedTypeFieldList  *ast.FieldList
// }

// func getEmbeddingDetails(users map[config.Typename][]models.FlattenKeypath) (map[config.Typename]embeddingDetails, error) {
// 	details := map[config.Typename]embeddingDetails{}

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
