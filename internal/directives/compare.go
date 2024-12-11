package directives

import (
	"strings"

	"github.com/ufukty/gonfique/internal/paths/models"
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
