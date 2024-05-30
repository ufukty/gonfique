package mappings

import (
	"fmt"
	"go/ast"
	"strings"
	"testing"

	"github.com/ufukty/gonfique/internal/mappings/testdata/appendix"
	"github.com/ufukty/gonfique/internal/testutils"
)

func compareMatchs(got []matchitem, want []ast.Node) bool {
	if len(got) != len(want) {
		return false
	}
	for i := 0; i < len(got); i++ {
		if got[i].holder != want[i] {
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
		{"metadata.name", []ast.Node{appendix.MetadataName}},

		{"spec.template.metadata.labels.app", []ast.Node{appendix.SpecTemplateMetadataLabelsApp}},
		{"spec.template.*.labels.app", []ast.Node{appendix.SpecTemplateMetadataLabelsApp}},
		{"spec.*.*.labels.app", []ast.Node{appendix.SpecTemplateMetadataLabelsApp}},

		{"spec.**.app", []ast.Node{appendix.SpecSelectorMatchLabelsApp, appendix.SpecTemplateMetadataLabelsApp}},

		{"spec.template.spec.containers", []ast.Node{appendix.SpecTemplateSpecContainers}},
		{"spec.template.spec.containers.[]", []ast.Node{appendix.SpecTemplateSpecContainersItem}},
		{"spec.template.spec.containers.[].name", []ast.Node{appendix.SpecTemplateSpecContainersItemName}},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			got, err := matchTypeDefHolder(appendix.K8sCfgTs, tc.input, appendix.Keys)
			if err != nil {
				t.Fatal(fmt.Errorf("act: %w", err))
			}
			if !compareMatchs(got, tc.want) {
				t.Fatalf("want %#v got %#v", testutils.NodeSliceString(tc.want), matchitemsToString(got))
			}
		})
	}
}
