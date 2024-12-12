package coder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/internal/files/config/meta"
	"github.com/ufukty/gonfique/internal/files/input"
	"github.com/ufukty/gonfique/internal/testutils"
	"github.com/ufukty/gonfique/internal/transform"
)

func TestCoder(t *testing.T) {
	type tc struct {
		encoding     input.Encoding
		folder, file string
	}
	tcs := []tc{
		{
			encoding: input.Yaml,
			folder:   "api",
			file:     "input.yml",
		},
		{
			encoding: input.Yaml,
			folder:   "generic",
			file:     "input.yml",
		},
		{
			encoding: input.Yaml,
			folder:   "k8s",
			file:     "input.yml",
		},
		{
			encoding: input.Yaml,
			folder:   "null",
			file:     "input.yml",
		},
		{
			encoding: input.Json,
			folder:   "null-json",
			file:     "input.json",
		},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%s_%s", tc.folder, tc.encoding), func(t *testing.T) {
			f, enc, err := input.Read(filepath.Join("testdata", tc.folder, tc.file))
			if err != nil {
				t.Fatal(fmt.Errorf("prep, read: %w", err))
			}

			testloc, err := testutils.PrepareTestCase(tc.folder, []string{"go.mod", "go.sum", "config_test.go", tc.file})
			if err != nil {
				t.Fatal(fmt.Errorf("prep, temp: :%w", err))
			}

			ti := transform.Transform(f, enc)
			c := Coder{
				Meta:     meta.Default,
				Encoding: tc.encoding,
				Config:   ti.Type,
			}
			if err := c.Write(filepath.Join(testloc, "config.go")); err != nil {
				t.Fatal(fmt.Errorf("act, write: %w", err))
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
