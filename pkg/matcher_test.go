package pkg

import (
	"go/ast"
	"slices"
	"strings"
	"testing"

	"github.com/ufukty/gonfique/pkg/testdata"
)

func TestMatcher(t *testing.T) {
	var (
		k8sFile  = testdata.KubernetesExampleOutput
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

		specTemplateSpec          = specTemplate.Type.(*ast.StructType).Fields.List[1]
		specTemplateSpecContainer = specTemplateSpec.Type.(*ast.StructType).Fields.List[0]
		specTemplateSpecContainerItem = specTemplateSpecContainer.Type.(*ast.ArrayType).Elt

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

		{"spec.**.app", []ast.Node{specTemplateMetadataLabelsApp, specSelectorMatchLabelsApp}},

		{"spec.template.spec.container.[*].", []ast.Node{specTemplateMetadataLabelsApp}}, // for array-type parents in keypath
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			got := match(k8sCfgTs.Type, strings.Split(tc.input, "."), []string{})
			if d := slices.Compare(tc.want, got); d != 0 {
				t.Fatalf("want %#v got %#v", got, tc.want)
			}
		})
	}
}
