package generate

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/internal/testutils"
)

func TestWithConfig(t *testing.T) {
	files := []string{
		"config_test.go",
		"extend.go",
		"go.mod",
		"go.sum",
		"input.yml",
	}
	tcs := []string{
		"api", "k8s",
	}
	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			testloc, err := testutils.PrepareTestCase(tc, files)
			if err != nil {
				t.Fatal(fmt.Errorf("prep, test folder: %w", err))
			}

			err = withconfig(
				filepath.Join(testloc, "input.yml"),
				filepath.Join("testdata", tc, "config.yml"),
				filepath.Join(testloc, "config.go"),
				true,
			)
			if err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}

			cmd := exec.Command("/usr/local/go/bin/go", "test",
				"-timeout", "10s",
				"-run", "^TestConfig$",
				"test",
				"-v", "-count=1",
			)
			cmd.Dir = testloc
			cmd.Stdout = os.Stderr
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				t.Fatal(fmt.Errorf("assert, go-test: %w", err))
			}
		})
	}
}
