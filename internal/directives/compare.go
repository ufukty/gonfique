package directives

import (
	"strings"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/transform"
)

func caseInsensitiveCompareTypenames(a, b config.Typename) int {
	if strings.ToLower(string(a)) < strings.ToLower(string(b)) {
		return -1
	} else {
		return +1
	}
}

func caseInsensitiveCompareFieldnames(a, b transform.FieldName) int {
	if strings.ToLower(string(a)) < strings.ToLower(string(b)) {
		return -1
	} else {
		return +1
	}
}
