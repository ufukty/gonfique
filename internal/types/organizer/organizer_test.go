package organizer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"go.ufukty.com/gonfique/internal/coder"
	"go.ufukty.com/gonfique/internal/files/input"
	"go.ufukty.com/gonfique/internal/testutils"
	"go.ufukty.com/gonfique/internal/transform"
)

func TestOrganizer(t *testing.T) {
	tcs := []string{"api"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			b := bundle.New("Config")

			f, enc, err := input.Read(filepath.Join("testdata", tc, "config.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			testloc, err := testutils.PrepareTestCase(tc, []string{"go.mod", "go.sum", "config_test.go", "config.yml"})
			if err != nil {
				t.Error(fmt.Errorf("preparing testcase to test: :%w", err))
			}

			transform.Transform(b)
			Organize(b)

			if err := coder.Write(b, filepath.Join(testloc, "config.go"), "config"); err != nil {
				t.Fatal(fmt.Errorf("creating config.go file: %w", err))
			}

			cmd := exec.Command("/usr/local/go/bin/go", "test",
				"-timeout", "10s",
				"-run", "^TestOrganize$",
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
