package version

import (
	"fmt"

	"github.com/ufukty/gonfique/internal/version"
)

func Run() error {
	fmt.Println(version.Version)
	return nil
}
