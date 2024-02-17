package pkg

import (
	"go/ast"
	"slices"
	"testing"

	"github.com/ufukty/gonfique/pkg/testdata/appendix"
)

func TestMatch(t *testing.T) {
	var (
		k8sFile  = appendix.KubernetesExampleOutput
		k8sCfg   = k8sFile.Decls[1].(*ast.GenDecl)
		k8sCfgTs = k8sCfg.Specs[0].(*ast.TypeSpec)
		k8sCfgFs = k8sCfgTs.Type.(*ast.StructType).Fields.List

		apiVersion   = k8sCfgFs[0]
		metadata     = k8sCfgFs[3]
		metadataName = metadata.Type.(*ast.StructType).Fields.List[0]
		spec         = k8sCfgFs[4]

		specTemplate                  = spec.Type.(*ast.StructType).Fields.List[4]
		specTemplateMetadata          = specTemplate.Type.(*ast.StructType).Fields.List[0]
		specTemplateMetadataLabels    = specTemplateMetadata.Type.(*ast.StructType).Fields.List[0]
		specTemplateMetadataLabelsApp = specTemplateMetadataLabels.Type.(*ast.StructType).Fields.List[0]

		specTemplateSpec                   = specTemplate.Type.(*ast.StructType).Fields.List[1]
		specTemplateSpecContainers         = specTemplateSpec.Type.(*ast.StructType).Fields.List[0]
		specTemplateSpecContainersItem     = specTemplateSpecContainers.Type.(*ast.ArrayType).Elt
		specTemplateSpecContainersItemName = specTemplateSpecContainersItem.(*ast.StructType).Fields.List[2]

		specSelector               = spec.Type.(*ast.StructType).Fields.List[3]
		specSelectorMatchLabels    = specSelector.Type.(*ast.StructType).Fields.List[0]
		specSelectorMatchLabelsApp = specSelectorMatchLabels.Type.(*ast.StructType).Fields.List[0]
	)

	type testcase struct {
		input string
		want  []ast.Node
	}

	tcs := []testcase{
		{"apiVersion", []ast.Node{apiVersion}},
		{"metadata.name", []ast.Node{metadataName}},

		{"spec.template.metadata.labels.app", []ast.Node{specTemplateMetadataLabelsApp}},
		{"spec.template.*.labels.app", []ast.Node{specTemplateMetadataLabelsApp}},
		{"spec.*.*.labels.app", []ast.Node{specTemplateMetadataLabelsApp}},

		{"spec.**.app", []ast.Node{specSelectorMatchLabelsApp, specTemplateMetadataLabelsApp}},

		{"spec.template.spec.containers", []ast.Node{specTemplateSpecContainers}},
		{"spec.template.spec.containers.[]", []ast.Node{specTemplateSpecContainersItem}},
		{"spec.template.spec.containers.[].name", []ast.Node{specTemplateSpecContainersItemName}},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			got := Match(k8sCfgTs, tc.input)
			if d := slices.CompareFunc(tc.want, got, func(i, j ast.Node) int {
				if i != j {
					return 1
				}
				return 0
			}); d != 0 {
				t.Fatalf("want %#v got %#v", nodeSliceString(tc.want), nodeSliceString(got))
			}
		})
	}
}
