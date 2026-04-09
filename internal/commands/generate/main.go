package generate

import (
	"flag"
	"fmt"
	"strings"

	"go.ufukty.com/gonfique/internal/generates"
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
		return fmt.Errorf("missing args: %s", strings.Join(ms, ", "))
	}
	return nil
}

func Run() error {
	args := getArgs()
	if err := missing(args); err != nil {
		return fmt.Errorf("checking args: %w", err)
	}
	return generates.FromPaths(args.In, args.Config, args.Out, args.Verbose)
}
