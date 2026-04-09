package help

import (
	_ "embed"
	"fmt"
)

//go:embed main.txt
var main string

func Run() error {
	fmt.Println(main)
	return nil
}
