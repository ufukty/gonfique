package main

import (
	"fmt"
	"os"

	"github.com/ufukty/gonfique/cmd/main/commands/generate"
	"github.com/ufukty/gonfique/cmd/main/commands/version"
)

func dispatch() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("gon")
	}
	switch os.Args[1] {
	case "version":
		return version.Run()
	default:
		return generate.Run()
	}
}

func main() {
	if err := dispatch(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
