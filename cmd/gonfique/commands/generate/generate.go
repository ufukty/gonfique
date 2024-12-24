package generate

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/files/coder"
	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/files/config/meta"
	"github.com/ufukty/gonfique/internal/files/config/validate"
	"github.com/ufukty/gonfique/internal/files/input"
	"github.com/ufukty/gonfique/internal/paths"
	"github.com/ufukty/gonfique/internal/transform"
	"github.com/ufukty/gonfique/internal/types"
)

func simple(in, out string) error {
	i, enc, err := input.Read(in)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	c := coder.Coder{
		Meta:     meta.Default,
		Encoding: enc,
		Config:   transform.Transform(i, enc).Type,
	}
	err = c.Write(out)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

func concat[K comparable, V any](ms ...map[K]V) map[K]V {
	t := 0
	for _, m := range ms {
		t += len(m)
	}
	n := make(map[K]V, t)
	for _, m := range ms {
		for k, v := range m {
			n[k] = v
		}
	}
	return n
}

func withconfig(in, conf, out string, verbose bool) error {
	i, enc, err := input.Read(in)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	ti := transform.Transform(i, enc)

	c, err := config.Read(conf)
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	err = validate.File(c)
	if err != nil {
		return fmt.Errorf("validating config file: %w", err)
	}
	aux, err := paths.Process(ti, c, verbose)
	if err != nil {
		return fmt.Errorf("applying value targeting rules: %w", err)
	}

	auxT, err := types.Apply(ti, c, concat(aux.Auto, aux.Declare))
	if err != nil {
		return fmt.Errorf("applying type targeting rules: %w", err)
	}

	coder := coder.Coder{
		Meta:      meta.Default,
		Encoding:  enc,
		Config:    ti.Type,
		Imports:   aux.Imports,
		Named:     aux.Declare,
		Auto:      aux.Auto,
		Accessors: auxT.Accessors,
	}
	err = coder.Write(out)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}
