package version

import "fmt"

var Version = "v0"

func Run() error {
	fmt.Println(Version)
	return nil
}
