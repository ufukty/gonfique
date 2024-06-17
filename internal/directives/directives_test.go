package directives

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/coder"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/directives/testdata/appendix"
	"github.com/ufukty/gonfique/internal/files"
	"github.com/ufukty/gonfique/internal/testutils"
	"github.com/ufukty/gonfique/internal/transform"
)

func TestAllKeypathsForHolders(t *testing.T) {
	b := &bundle.Bundle{
		CfgType:      appendix.ConfigType,
		OriginalKeys: appendix.Keys,
	}
	d := New(b)
	d.populateKeypathsAndHolders()
	got := d.Keypaths
	if len(got) != len(appendix.Keypaths) {
		t.Errorf("assert 1, length. want %d got %d", len(appendix.Keys), len(got))
	}
	for holder, wantkp := range appendix.Keypaths {
		if gotkp, ok := got[holder]; !ok {
			t.Errorf("assert 2, existence. want %q", wantkp)
		} else if gotkp != wantkp {
			t.Errorf("assert 3, mismatch. \nwant %q\ngot  %q", wantkp, gotkp)
		}
	}
}

func TestImplement(t *testing.T) {
	tcs := []string{"k8s", "api", "api-parent", "api-parent2"}

	for _, tc := range tcs {
		t.Run(tc, func(t *testing.T) {
			b := bundle.New("Config")

			err := files.ReadConfigFile(b, filepath.Join("testdata", tc, "config.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
			}

			transform.Transform(b)

			b.Df, err = directivefile.ReadDirectiveFile(filepath.Join("testdata", tc, "directives.yml"))
			if err != nil {
				t.Fatal(fmt.Errorf("reading directive file: %w", err))
			}

			d := New(b)
			if err = d.Apply(true); err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}

			testloc, err := testutils.PrepareTestCase(tc, []string{"go.mod", "go.sum", "config_test.go", "config.yml", "extend.go"})
			if err != nil {
				t.Error(fmt.Errorf("preparing testcase to test: :%w", err))
			}

			if err := coder.Write(b, filepath.Join(testloc, "config.go"), "config"); err != nil {
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
