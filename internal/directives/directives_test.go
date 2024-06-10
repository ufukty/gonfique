package directives

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/coder"
	"github.com/ufukty/gonfique/internal/directives/accessors"
	"github.com/ufukty/gonfique/internal/directives/check"
	"github.com/ufukty/gonfique/internal/directives/directivefile"
	"github.com/ufukty/gonfique/internal/directives/expansion"
	"github.com/ufukty/gonfique/internal/directives/named"
	"github.com/ufukty/gonfique/internal/files"
	"github.com/ufukty/gonfique/internal/namings"
	"github.com/ufukty/gonfique/internal/resolver"
	"github.com/ufukty/gonfique/internal/testutils"
	"github.com/ufukty/gonfique/internal/transform"
	"golang.org/x/exp/maps"
)

func TestImplement(t *testing.T) {
	tcs := []string{"k8s", "api"}

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
			resolver.AllKeypathsForHolders(b)
			err = check.PopulateExprs(b)
			if err != nil {
				t.Fatal(fmt.Errorf("collecting type expressions for each keypaths: %w", err))
			}
			if err = expansion.ExpandKeypathsInDirectives(b); err != nil {
				t.Fatal(fmt.Errorf("expanding: %w", err))
			}
			check.MarkNeededNamedTypes(b)
			b.GeneratedTypenames = namings.GenerateTypenames(maps.Values(b.Keypaths))

			err = named.Implement(b)
			if err != nil {
				t.Fatal(fmt.Errorf("declaring named types: %w", err))
			}
			if err = accessors.Implement(b); err != nil {
				t.Fatal(fmt.Errorf("implement: %w", err))
			}

			testloc, err := testutils.PrepareTestCase(tc, []string{"go.mod", "go.sum", "config_test.go", "config.yml"})
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
