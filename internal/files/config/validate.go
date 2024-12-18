package config

import (
	"fmt"
	"regexp"

	"github.com/ufukty/gonfique/internal/files/config/tree"
)

func validatePathsSection(paths map[Path]PathConfig) string {
	typenames := regexp.MustCompile(`<\w+>`)
	msg1 := []string{}
	for cp := range paths {
		msg2 := []string{}
		for i, s := range cp.Segments() {
			if typenames.MatchString(s) != (i == 0) {
				msg2 = append(msg2, "path should contain a typename as the first term and not later")
			}
		}
		if len(msg2) > 0 {
			msg1 = append(msg1, tree.List(string(cp), msg2))
		}
	}
	if len(msg1) > 0 {
		return tree.List("path section", msg1)
	}
	return ""
}

func Validate(f *File) error {
	msg := []string{}

	if err := validatePathsSection(f.Paths); err != "" {
		msg = append(msg, err)
	}

	if len(msg) > 0 {
		return fmt.Errorf("%s", tree.List("file", msg))
	}
	return nil
}
