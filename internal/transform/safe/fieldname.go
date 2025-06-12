package safe

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/ufukty/gonfique/internal/files/config"
)

var smallcaps = regexp.MustCompile("[a-z]+")

// Best effort on creating Go var/field names out of YAML keys idiomatically
func FieldName(key string) config.Fieldname {
	if !smallcaps.MatchString(key) {
		key = strings.ToLower(key)
	}
	n := ""
	newSegment := false
	for i, r := range key {
		if i == 0 || newSegment {
			n += strings.ToUpper(string(r))
			newSegment = false
		} else if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			newSegment = true
		} else {
			n += string(r)
		}
	}
	return config.Fieldname(n)
}
