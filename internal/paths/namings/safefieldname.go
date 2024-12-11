package namings

import (
	"regexp"
	"strings"
	"unicode"
)

var smallcaps = regexp.MustCompile("[a-z]+")

// Best effort on creating Go var/field names out of YAML keys idiomatically
func SafeFieldName(keyname string) string {
	if smallcaps.Find([]byte(keyname)) == nil {
		keyname = strings.ToLower(keyname)
	}
	n := ""
	newSegment := false
	for i, r := range keyname {
		if i == 0 || newSegment {
			n += strings.ToUpper(string(r))
			newSegment = false
		} else if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			newSegment = true
		} else {
			n += string(r)
		}
	}
	return n
}

func SafeTypeName(keyname string, exported bool) string {
	if smallcaps.Find([]byte(keyname)) == nil {
		keyname = strings.ToLower(keyname)
	}
	n := ""
	newSegment := false
	for i, r := range keyname {
		if exported && i == 0 {
			n += strings.ToUpper(string(r))
		} else if newSegment {
			n += strings.ToUpper(string(r))
			newSegment = false
		} else if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			newSegment = true
		} else {
			n += string(r)
		}
	}
	return n
}
