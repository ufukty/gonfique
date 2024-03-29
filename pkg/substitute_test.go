package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestSubstitute(t *testing.T) {
	tcs := []string{"tc3", "tc4"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			cts, err := ReadConfigYaml(filepath.Join("testdata", tc, "config.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			testloc, err := os.MkdirTemp(os.TempDir(), "*")
			if err != nil {
				t.Fatal(fmt.Errorf("creating temporary directory to test the created schema: %w", err))
			}
			fmt.Println("using tmp dir:", testloc)

			files := []string{"go.mod", "go.sum", "config_test.go", "config.yml", "use.go"}
			for _, file := range files {
				src := filepath.Join("testdata", tc, file)
				dst := filepath.Join(testloc, file)
				if err := copyFile(src, dst); err != nil {
					t.Fatal(fmt.Errorf("copying %q to %q: %w", file, dst, err))
				}
			}

			etss, err := ReadTypes(filepath.Join("testdata", tc, "use.go"))
			if err != nil {
				t.Fatal(fmt.Errorf("reading types to use in substitution: %w", err))
			}

			Substitute(cts, etss)

			if err := WriteConfigGo(filepath.Join(testloc, "config.go"), cts, nil, nil, "config"); err != nil {
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
