package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ufukty/gonfique/pkg"
)

type Args struct {
	In  string
	Out string
	Pkg string
	Use string
	Org bool
}

func getArgs() Args {
	args := Args{}
	flag.StringVar(&args.In, "in", "", "input file path (yml or yaml)")
	flag.StringVar(&args.Out, "out", "", "output file path (go)")
	flag.StringVar(&args.Pkg, "pkg", "", "package name that will be inserted into the generated file")
	flag.StringVar(&args.Use, "use", "", "(optional) use type definitions found in <file>")
	flag.BoolVar(&args.Org, "organize", false, "(optional) defines the types of struct fields that are also structs separately instead inline, with auto generated UNSTABLE names.")
	flag.Parse()
	return args
}

func checkMissingArgs(args Args) error {
	ms := []string{}
	if args.In == "" {
		ms = append(ms, "-in")
	}
	if args.Out == "" {
		ms = append(ms, "-out")
	}
	if args.Pkg == "" {
		ms = append(ms, "-pkg")
	}
	if len(ms) > 0 {
		return fmt.Errorf("some arguments are missing: %s", strings.Join(ms, ", "))
	}
	return nil
}

func perform() error {
	args := getArgs()
	if err := checkMissingArgs(args); err != nil {
		return fmt.Errorf("checking args: %w", err)
	}

	cfgts, err := pkg.ReadConfigYaml(args.In)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}

	if args.Use != "" {
		tss, err := pkg.ReadTypes(args.Use)
		if err != nil {
			return fmt.Errorf("reading -use file %q: %w", args.Use, err)
		}
		pkg.Substitute(cfgts, tss)
	}

	if !args.Org {
		if err := pkg.WriteConfigGo(args.Out, cfgts, nil, nil, args.Pkg); err != nil {
			return fmt.Errorf("creating %q: %w", args.Out, err)
		}
	} else {
		isolated := pkg.Organize(cfgts)
		iterators, err := pkg.Iterators(cfgts, isolated)
		if err != nil {
			return fmt.Errorf("creating iterators: %w", err)
		}
		if err := pkg.WriteConfigGo(args.Out, cfgts, isolated, iterators, args.Pkg); err != nil {
			return fmt.Errorf("creating %q: %w", args.Out, err)
		}
	}

	return nil
}

func main() {
	if err := perform(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
