package version

import (
	"fmt"

	"go.ufukty.com/gonfique/internal/version"
)

func Run() error {
	fmt.Println(version.Version)
	return nil
}
