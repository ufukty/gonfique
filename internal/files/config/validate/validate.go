package validate

import (
	"fmt"
	"regexp"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/tree/bucket"
)

var typenames = regexp.MustCompile(`<\w+>`)

func path(cp config.Path, b *bucket.Bucket) {
	for i, s := range cp.Segments() {
		if i != 0 && typenames.MatchString(s) {
			b.Add(fmt.Sprintf("type segment after start: %q", s))
		}
	}
}

func directives(pc config.Directives, b *bucket.Bucket) {
	if err := pc.Dict.Validate(); err != nil {
		b.Add(fmt.Sprintf("checking 'dict' value: %s", err))
	}
	if pc.IsZero() {
		b.Add("directives are missing")
	}
}

func rules(rules map[config.Path]config.Directives, b *bucket.Bucket) {
	for cp, pc := range rules {
		s := b.Sub(string(cp))
		path(cp, s.Sub("path"))
		directives(pc, s.Sub("directives"))
	}
}

func File(f *config.File) error {
	b := bucket.New("file")
	rules(f.Rules, b.Sub("rules"))

	if !b.IsEmpty() {
		return fmt.Errorf("%s", b)
	}
	return nil
}
