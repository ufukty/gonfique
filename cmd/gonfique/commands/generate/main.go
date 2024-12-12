package generate

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ufukty/gonfique/internal/coder"
	"github.com/ufukty/gonfique/internal/files/input"
	"github.com/ufukty/gonfique/internal/transform"
)

type Args struct {
	In      string
	Out     string
	Config  string
	Verbose bool
}

func getArgs() Args {
	args := Args{}
	flag.StringVar(&args.In, "in", "", "path")
	flag.StringVar(&args.Out, "out", "", "path")
	flag.StringVar(&args.Config, "config", "", "(optional) path to a Gonfique config")
	flag.BoolVar(&args.Verbose, "verbose", false, "(optional) print information regarding to processes.")
	flag.Parse()
	return args
}

func missing(args Args) error {
	ms := []string{}
	if args.In == "" {
		ms = append(ms, "-in")
	}
	if args.Out == "" {
		ms = append(ms, "-out")
	}
	if len(ms) > 0 {
		return fmt.Errorf("some arguments are missing: %s", strings.Join(ms, ", "))
	}
	return nil
}

func Run() error {
	args := getArgs()
	if err := missing(args); err != nil {
		return fmt.Errorf("checking args: %w", err)
	}

	i, e, err := input.Read(args.In)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}
	t := transform.Transform(i, e)

	// directives
	// substitude
	// mappings
	// organizer
	// iterables

	if err := coder.Write(t, args.Out, args.Pkg); err != nil {
		return fmt.Errorf("creating %q: %w", args.Out, err)
	}

	return nil
}
