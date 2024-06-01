package mappings

import (
	"fmt"
	"go/ast"
	"slices"
	"strings"
	"testing"

	"github.com/ufukty/gonfique/internal/mappings/testdata/appendix"
	"github.com/ufukty/gonfique/internal/testutils"
)

func compareMatchs(got []matchitem, want []ast.Node) bool {
	if len(got) != len(want) {
		return false
	}
	for _, w := range want {
		if slices.IndexFunc(got, func(mi matchitem) bool { return mi.holder == w }) == -1 {
			return false
		}
	}
	return true
}

func matchitemsToString(mis []matchitem) string {
	s := []string{}
	for _, mi := range mis {
		s = append(s, testutils.NodeString(mi.holder))
	}
	return strings.Join(s, ", ")
}

func TestMatch(t *testing.T) {
	type testcase struct {
		input string
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
		{"**.configMapRef.Name", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName}},
		{"**.configMapRef.*", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName}},
		{"**.configMapRef.**", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName}},
		{"spec.**.configMapRef.**", []ast.Node{appendix.SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName}},
		{"**.envFrom.**", []ast.Node{
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
		t.Run(tc.input, func(t *testing.T) {
			got, err := matchTypeDefHolder(appendix.K8sCfgTs.Type, tc.input, appendix.Keys)
			if err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}
			if !compareMatchs(got, tc.want) {
				t.Fatalf("mismatch\nwant: %#v\ngot : %#v", testutils.NodeSliceString(tc.want), matchitemsToString(got))
			}
		})
	}
}
