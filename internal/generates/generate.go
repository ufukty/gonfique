package generates

import (
	"fmt"
	"io"
	"os"

	"go.ufukty.com/gonfique/internal/files/coder"
	"go.ufukty.com/gonfique/internal/files/config"
	"go.ufukty.com/gonfique/internal/files/config/meta"
	"go.ufukty.com/gonfique/internal/files/config/validate"
	"go.ufukty.com/gonfique/internal/files/input"
	"go.ufukty.com/gonfique/internal/files/input/encoders"
	"go.ufukty.com/gonfique/internal/paths"
	"go.ufukty.com/gonfique/internal/transform"
	"go.ufukty.com/gonfique/internal/types"
)

func simple(in io.Reader, enc encoders.Encoding, out io.Writer) error {
	i, err := input.Read(in, enc)
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

func withConfig(in io.Reader, enc encoders.Encoding, conf io.Reader, out io.Writer, verbose bool) error {
	i, err := input.Read(in, enc)
	if err != nil {
		return fmt.Errorf("input file: %w", err)
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
		Iterators: auxT.Iterator,
	}

	err = coder.Write(out)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

func FromReaders(in io.Reader, enc encoders.Encoding, conf io.Reader, out io.Writer, verbose bool) error {
	if conf == nil {
		return simple(in, enc, out)
	}
	return withConfig(in, enc, conf, out, verbose)
}

func FromPaths(in, conf, out string, verbose bool) error {
	enc, err := encoders.FromExtension(in)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	i, err := os.Open(in)
	if err != nil {
		return fmt.Errorf("input: %w", err)
	}
	defer i.Close()

	o, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}
	defer o.Close()

	if conf == "" {
		return FromReaders(i, enc, nil, o, verbose)
	}

	c, err := os.Open(conf)
	if err != nil {
		return fmt.Errorf("open config file: %w", err)
	}
	defer c.Close()

	return FromReaders(i, enc, c, o, verbose)
}
