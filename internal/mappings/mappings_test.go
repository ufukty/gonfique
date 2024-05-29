package mappings

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/internal/files"
	"github.com/ufukty/gonfique/internal/testutils"
)

func TestMappings(t *testing.T) {

	tcs := []string{"k8s", "api"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {

			f, err := files.ReadConfigFile(filepath.Join("testdata", tc, "config.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			ms, err := ReadMappings(filepath.Join("testdata", tc, "mappings.gonfique.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("reading mappings from user-provided file: %w", err))
			}

			// apply mappings before "organize" & "iterate"
			ApplyMappings(f, ms)

			testloc, err := testutils.PrepareTestCase(tc, []string{"go.mod", "go.sum", "config_test.go", "config.yml", "extend.go"})
			if err != nil {
				t.Error(fmt.Errorf("preparing testcase to test: :%w", err))
			}

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
