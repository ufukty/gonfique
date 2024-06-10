package main

import (
	"fmt"
	"os"

	"github.com/ufukty/gonfique/cmd/gonfique/commands/generate"
	"github.com/ufukty/gonfique/cmd/gonfique/commands/help"
	"github.com/ufukty/gonfique/cmd/gonfique/commands/version"
)

type Run func() error

var commands = map[string]Run{
	"help":     help.Run,
	"version":  version.Run,
	"generate": generate.Run,
}

func dispatch() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("needs more arguments. run: gonfique help")
	}

	arg := os.Args[1]
	os.Args = os.Args[1:]
	command, ok := commands[arg]
	if !ok {
		return fmt.Errorf("command %q doesn't exist. run: gonfique help", arg)
	}
	err := command()
	if err != nil {
		return fmt.Errorf("%s: %w", arg, err)
	}
	return nil
}

func main() {
	if err := dispatch(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
