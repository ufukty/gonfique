package generate

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/coder"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/files"
	"github.com/ufukty/gonfique/internal/iterables"
	"github.com/ufukty/gonfique/internal/mappings"
	"github.com/ufukty/gonfique/internal/organizer"
	"github.com/ufukty/gonfique/internal/substitude"
	"github.com/ufukty/gonfique/internal/transform"
)

type Args struct {
	In         string
	Out        string
	Pkg        string
	TypeName   string
	Use        string
	Org        bool
	Mappings   string
	Directives string
}

func getArgs() Args {
	args := Args{}
	flag.StringVar(&args.In, "in", "", "input file path (yml or yaml)")
	flag.StringVar(&args.Out, "out", "", "output file path (go)")
	flag.StringVar(&args.Pkg, "pkg", "", "package name that will be inserted into the generated file")
	flag.StringVar(&args.TypeName, "type-name", "Config", "will be used to name generated type")
	flag.StringVar(&args.Use, "use", "", "(optional) use type definitions found in <file>")
	flag.StringVar(&args.Mappings, "mappings", "", "(optional) use typenames found in the <file>. see examples for mapping file structure")
	flag.StringVar(&args.Directives, "directives", "", "(optional) use a directives file")
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

func checkConflictingFeatures(args Args) error {
	if args.Directives != "" && args.Mappings != "" {
		return fmt.Errorf("using 'directives' and 'mappings' together is not allowed")
	}
	return nil
}

func Run() error {
	args := getArgs()
	if err := checkMissingArgs(args); err != nil {
		return fmt.Errorf("checking args: %w", err)
	}

	if err := checkConflictingFeatures(args); err != nil {
		return fmt.Errorf("checking conflicting features: %w", err)
	}

	b := bundle.New(args.TypeName)

	err := files.ReadConfigFile(b, args.In)
	if err != nil {
		return fmt.Errorf("reading input file: %w", err)
	}

	transform.Transform(b)

	if args.Directives != "" {
		b.Df, err = directivefile.ReadDirectiveFile(args.Directives)
		if err != nil {
			return fmt.Errorf("reading directives file: %w", err)
		}
		err = directives.Apply(b)
		if err != nil {
			return fmt.Errorf("applying directives: %w", err)
		}
	}

	if args.Use != "" {
		tss, err := substitude.ReadTypes(args.Use)
		if err != nil {
			return fmt.Errorf("reading -use file %q: %w", args.Use, err)
		}
		substitude.UserProvided(b, tss)
	}

	if args.Mappings != "" {
		rules, err := files.ReadMappings(args.Mappings)
		if err != nil {
			return fmt.Errorf("reading -mappings file %q: %w", args.Mappings, err)
		}
		mappings.ApplyMappings(b, rules)
	}

	if args.Org {
		organizer.Organize(b)
	}

	if b.Isolated != nil || b.Named != nil {
		err := iterables.ImplementIterators(b)
		if err != nil {
			return fmt.Errorf("creating iterators: %w", err)
		}
	}

	if err := coder.Write(b, args.Out, args.Pkg); err != nil {
		return fmt.Errorf("creating %q: %w", args.Out, err)
	}

	return nil
}
