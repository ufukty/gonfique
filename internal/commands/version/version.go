package version

import (
	"fmt"

	"go.ufukty.com/gonfique/v2/internal/version"
)

func Run() error {
	v, err := version.OfBuild()
	if err != nil {
		return fmt.Errorf("digging build details: %w", err)
	}
	fmt.Println(v)
	return nil
}
