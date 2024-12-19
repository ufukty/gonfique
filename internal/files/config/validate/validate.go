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
			b.Add(fmt.Sprintf("non-starting type segment: %s", s))
		}
	}
}

func pathDirectives(pc config.PathConfig, b *bucket.Bucket) {
	if err := pc.Dict.Validate(); err != nil {
		b.Add(fmt.Sprintf("checking 'dict' value: %s", err))
	}
}

func paths(paths map[config.Path]config.PathConfig, b *bucket.Bucket) {
	for cp, pc := range paths {
		s := b.Sub(string(cp))
		path(cp, s.Sub("path"))
		pathDirectives(pc, s.Sub("directives"))
	}
}

func File(f *config.File) error {
	b := bucket.New("file")
	paths(f.Paths, b.Sub("paths"))

	if !b.IsEmpty() {
		return fmt.Errorf("%s", b)
	}
	return nil
}
