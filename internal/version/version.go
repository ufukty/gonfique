package version

import (
	"fmt"
	"runtime/debug"
)

func OfBuild() (string, error) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "", fmt.Errorf("build info is not available")
	}
	return bi.Main.Version, nil
}
