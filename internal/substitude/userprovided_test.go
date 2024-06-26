package substitude

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/internal/files"
	"github.com/ufukty/gonfique/internal/testutils"
)

func TestSubstitute(t *testing.T) {
	tcs := []string{"api", "apiv2"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			f, err := files.ReadConfigFile(filepath.Join("testdata", tc, "config.yml"), "Config")
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			testloc, err := testutils.PrepareTestCase(tc, []string{"go.mod", "go.sum", "config_test.go", "config.yml", "use.go"})
			if err != nil {
				t.Error(fmt.Errorf("preparing testcase to test: :%w", err))
			}

			etss, err := ReadTypes(filepath.Join("testdata", tc, "use.go"))
			if err != nil {
				t.Fatal(fmt.Errorf("reading types to use in substitution: %w", err))
			}

			UserProvided(f, etss)

			if err := f.Write(filepath.Join(testloc, "config.go"), "config"); err != nil {
				t.Fatal(fmt.Errorf("creating config.go file: %w", err))
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
				t.Fatal(fmt.Errorf("running go-test: %w", err))
			}

		})
	}
}
