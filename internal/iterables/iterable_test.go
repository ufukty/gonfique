package iterables

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/internal/coder"
	"github.com/ufukty/gonfique/internal/files"
	"github.com/ufukty/gonfique/internal/organizer"
	"github.com/ufukty/gonfique/internal/testutils"
)

func TestIterators(t *testing.T) {
	tcs := []string{"api", "small"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			f, err := files.ReadConfigFile(filepath.Join("testdata", tc, "config.yml"), "Config")
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			testloc, err := testutils.PrepareTestCase(tc, []string{"go.mod", "go.sum", "config_test.go", "config.yml"})
			if err != nil {
				t.Error(fmt.Errorf("preparing testcase to test: :%w", err))
			}

			organizer.Organize(f)
			err = ImplementIterators(f)
			if err != nil {
				t.Fatal(fmt.Errorf("generating iterators for all-same-type-field structs: %w", err))
			}

			if err := coder.Write(f, filepath.Join(testloc, "config.go"), "config"); err != nil {
				t.Fatal(fmt.Errorf("creating config.go file: %w", err))
			}

			cmd := exec.Command("/usr/local/go/bin/go", "test",
				"-timeout", "10s",
				"-run", "^TestIterators$",
				"test",
				"-v", "-count=1",
			)
			cmd.Dir = testloc
			cmd.Stdout = os.Stderr
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				t.Fatal(fmt.Errorf("running go-test: %w", err))
			}

		})

	}
}
