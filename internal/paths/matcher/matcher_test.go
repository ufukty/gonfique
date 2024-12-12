package matcher

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"
	"testing"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/matcher/testdata/appendix"
	"github.com/ufukty/gonfique/internal/testutils"
	"github.com/ufukty/gonfique/internal/transform"
)

func compareMatches(got []ast.Node, want []ast.Node) bool {
	if len(got) != len(want) {
		return false
	}
	for _, w := range want {
		if slices.IndexFunc(got, func(n ast.Node) bool { return n == w }) == -1 {
			return false
		}
	}
	return true
}

func matchitemsToString(nodes []ast.Node) string {
	s := []string{}
	for _, n := range nodes {
		s = append(s, testutils.NodeString(n))
	}
	return strings.Join(s, ", ")
}

func TestMatch(t *testing.T) {
	type testcase struct {
		input config.Path
		want  []ast.Node
	}

	tcs := []testcase{
		{"apiVersion", []ast.Node{appendix.ApiVersion}},
		{"metadata", []ast.Node{appendix.Metadata}},
		{"metadata.name", []ast.Node{appendix.MetadataName}},
		{"spec.template.metadata.labels.app", []ast.Node{appendix.SpecTemplateMetadataLabelsApp}},
		{"spec.template.*.labels.app", []ast.Node{appendix.SpecTemplateMetadataLabelsApp}},
		{"spec.*.*.labels.app", []ast.Node{appendix.SpecTemplateMetadataLabelsApp}},
		{"spec.template.spec.containers", []ast.Node{appendix.SpecTemplateSpecContainers}},
		{"spec.template.spec.containers.[]", []ast.Node{appendix.SpecTemplateSpecContainersItem}},
		{"spec.template.spec.containers.[].name", []ast.Node{appendix.SpecTemplateSpecContainersItemName}},
		{"**.configMapRef", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRef}},
		{"**.configMapRef.name", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName}},
		{"**.configMapRef.*", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName}},
		{"**.configMapRef.**", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName}},
		{"spec.**.configMapRef.**", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName}},
		{"**.envFrom.**", []ast.Node{
			appendix.SpecTemplateSpecContainersItemEnvFromItem,
			appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRef,
			appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName,
			appendix.SpecTemplateSpecContainersItemEnvFromItemSecretRef,
			appendix.SpecTemplateSpecContainersItemEnvFromItemSecretRefName,
		}},
		{"*.name", []ast.Node{appendix.MetadataName}},
		{"**.name", []ast.Node{
			appendix.MetadataName,
			appendix.SpecRulesItemHttpPathsItemBackendServiceName,
			appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName,
			appendix.SpecTemplateSpecContainersItemEnvFromItemSecretRefName,
			appendix.SpecTemplateSpecContainersItemName,
		}},
		{"spec.**.app", []ast.Node{
			appendix.SpecSelectorMatchLabelsApp,
			appendix.SpecTemplateMetadataLabelsApp,
		}},
	}

	for _, tc := range tcs {
		t.Run(string(tc.input), func(t *testing.T) {
			ti := &transform.Info{
				Type: appendix.K8sCfgTs.Type,
				Keys: appendix.Keys,
			}
			m := New(ti)
			got, err := m.FindHolders(tc.input)
			if err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}
			if !compareMatches(got, tc.want) {
				t.Fatalf("mismatch\nwant: %#v\ngot : %#v", testutils.NodeSliceString(tc.want), matchitemsToString(got))
			}
		})
	}
}
