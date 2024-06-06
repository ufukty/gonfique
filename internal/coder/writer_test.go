package coder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/files"
	"github.com/ufukty/gonfique/internal/testutils"
)

func TestCreatation(t *testing.T) {
	tcs := []string{"api", "generic"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			cfgcontent, encoding, err := files.ReadConfigFile(filepath.Join("testdata", tc, "config.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			b := bundle.New(cfgcontent, encoding, "Config")

			if err := Write(b, os.DevNull, "config"); err != nil {
				t.Fatal(fmt.Errorf("creating config.go file: %w", err))
			}
		})
	}
}

func Test_CreateAndUseForYaml(t *testing.T) {
	tcs := []string{"api", "generic", "k8s", "duration", "null"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			cfgcontent, encoding, err := files.ReadConfigFile(filepath.Join("testdata", tc, "config.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			b := bundle.New(cfgcontent, encoding, "Config")

			testloc, err := testutils.PrepareTestCase(tc, []string{"go.mod", "go.sum", "config_test.go", "config.yml"})
			if err != nil {
				t.Error(fmt.Errorf("preparing testcase to test: :%w", err))
			}

			if err := Write(b, filepath.Join(testloc, "config.go"), "config"); err != nil {
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

func Test_CreateAndUseForJson(t *testing.T) {
	tcs := []string{"null-json"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			cfgcontent, encoding, err := files.ReadConfigFile(filepath.Join("testdata", tc, "config.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			b := bundle.New(cfgcontent, encoding, "Config")

			testloc, err := testutils.PrepareTestCase(tc, []string{"go.mod", "go.sum", "config_test.go", "config.json"})
			if err != nil {
				t.Error(fmt.Errorf("preparing testcase to test: :%w", err))
			}

			if err := Write(b, filepath.Join(testloc, "config.go"), "config"); err != nil {
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
