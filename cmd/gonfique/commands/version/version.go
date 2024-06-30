package version

import "fmt"

var Version = ""

func Run() error {
	fmt.Println(Version)
	return nil
}
