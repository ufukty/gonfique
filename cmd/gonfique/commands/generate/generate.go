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

func withconfig(in, conf, out string, verbose bool) error {
	i, enc, err := input.Read(in)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	ti := transform.Transform(i, enc)

	f, err := config.Read(conf)
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	err = validate.File(f)
	if err != nil {
		return fmt.Errorf("validating config file: %w", err)
	}
	aux, err := paths.Process(&ti, f, verbose)
	if err != nil {
		return fmt.Errorf("paths: %w", err)
	}

	c := coder.Coder{
		Meta:     meta.Default,
		Encoding: enc,
		Config:   ti.Type,
		Imports:  aux.Imports,
		Named:    aux.Declare,
		Auto:     aux.Auto,
	}
	err = c.Write(out)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}
