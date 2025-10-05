package validate

import (
	"fmt"
	"regexp"

	"go.ufukty.com/gonfique/internal/datas/sortby"
	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/tree/bucket"
)

var typenames = regexp.MustCompile(`<\w+>`)

type target string

const (
	typeTargeting  target = "type"
	valueTargeting target = "value"
)

func path(cp config.Path, b *bucket.Bucket) target {
	ss := cp.Segments()
	for i, s := range ss {
		if i != 0 && typenames.MatchString(s) {
			b.Add(fmt.Sprintf("type segment after start: %q", s))
		}
	}
	if len(ss) == 1 && typenames.MatchString(ss[0]) {
		return typeTargeting
	}
	return valueTargeting
}

func directives(pc config.Directives, t target, b *bucket.Bucket) {
	if err := pc.Dict.Validate(); err != nil {
		b.Add(fmt.Sprintf("checking 'dict' value: %s", err))
	}
	if pc.IsZero() {
		b.Add("directives are missing")
	}
	switch t {
	case valueTargeting:
		b = b.Sub("value targeting rules can't contain directives for types")
		if pc.Iterator {
			b.Add("iterator")
		}
		if pc.Embed != "" {
			b.Add("embed")
		}
		if len(pc.Accessors) > 0 {
			b.Add("accessors")
		}
		if pc.Parent != "" {
			b.Add("parent")
		}
	case typeTargeting:
		b = b.Sub("type targeting rules can't contain directives for values")
		if pc.Declare != "" {
			b.Add("declare")
		}
		if pc.Export {
			b.Add("export")
		}
		if pc.Replace != "" {
			b.Add("replace")
		}
		if pc.Dict != "" {
			b.Add("dict")
		}
	}
}

func rules(rules map[config.Path]config.Directives, b *bucket.Bucket) {
	for cp, pc := range sortby.Key(rules) {
		s := b.Sub(string(cp))
		t := path(cp, s.Sub("path"))
		directives(pc, t, s.Sub("directives"))
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
