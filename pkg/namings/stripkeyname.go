package namings

import (
	"fmt"
	"regexp"
)

var yamlKeyStripper = regexp.MustCompile("`yaml:\"(.*)\"`")

func StripKeyname(tag string) (string, error) {
	matches := yamlKeyStripper.FindStringSubmatch(tag)
	if len(matches) < 2 {
		return "", fmt.Errorf("key not found")
	}
	return matches[1], nil
}
