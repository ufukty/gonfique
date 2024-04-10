package organizer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/pkg/files"
	"github.com/ufukty/gonfique/pkg/testutils"
)

func TestOrganizer(t *testing.T) {
	tcs := []string{"tc6-organize"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			cts, err := files.ReadConfigYaml(filepath.Join("testdata", tc, "config.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			testloc, err := os.MkdirTemp(os.TempDir(), "*")
			if err != nil {
				t.Fatal(fmt.Errorf("creating temporary directory to test the created schema: %w", err))
			}
			fmt.Println("using tmp dir:", testloc)

			filenames := []string{"go.mod", "go.sum", "config_test.go", "config.yml"}
			for _, file := range filenames {
				src := filepath.Join("testdata", tc, file)
				dst := filepath.Join(testloc, file)
				if err := testutils.CopyFile(src, dst); err != nil {
					t.Fatal(fmt.Errorf("copying %q to %q: %w", file, dst, err))
				}
			}

			if err := files.WriteConfigGo(filepath.Join(testloc, "config.go"), cts, nil, Organize(cts), nil, "config"); err != nil {
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
