package version

import "fmt"

var Version = "test version"

func Run() error {
	fmt.Println(Version)
	return nil
}
