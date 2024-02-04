package pkg

import (
	"regexp"
	"strings"
	"unicode"
)

// func startCapsAtEachSegment(s string, pattern string) string {
// 	s2 := ""
// 	for _, w := range strings.Split(s, pattern) {
// 		w2 := ""
// 		for i, r := range w {
// 			if i == 0 {
// 				w2 += strings.ToUpper(string(r))
// 			} else {
// 				w2 += strings.ToLower(string(r))
// 			}
// 		}
// 		s2 += w2
// 	}
// 	return s2
// }

var smallcaps = regexp.MustCompile("[a-z]+")

// Best effort on creating Go var/field names out of YAML keys idiomatically
func safeFieldName(keyname string) string {
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
