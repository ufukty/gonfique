package appendix

import "go/ast"

// // apiVersion
// //     mykey
// //     password
// // data
// // kind
// //     name
// //     namespace
// // metadata
// //          port
// //          protocol
// //          targetPort
// //     ports
// //     replicas
// //          host
// //                           name
// //                               number
// //                           port
// //                       service
// //                   backend
// //                   path
// //                   pathType
// //              paths
// //          http
// //     rules
// //             app
// //         matchLabels
// //     selector
// //                 app
// //             labels
// //         metadata
// //                           name
// //                       configMapRef
// //                           name
// //                       secretRef
// //                  envFrom
// //                  image
// //                  name
// //                       containerPort
// //                  ports
// //             containers
// //         spec
// //     template
// // spec
// // type

var (
	K8sCfg   = KubernetesExampleOutput.Decls[1].(*ast.GenDecl)
	K8sCfgTs = K8sCfg.Specs[0].(*ast.TypeSpec)
	K8sCfgFs = K8sCfgTs.Type.(*ast.StructType).Fields.List

	ApiVersion = K8sCfgFs[0]
	Data       = K8sCfgFs[1]
	Kind       = K8sCfgFs[2]
	Metadata   = K8sCfgFs[3]
	Spec       = K8sCfgFs[4]
	Type       = K8sCfgFs[5]

	MetadataName      = Metadata.Type.(*ast.StructType).Fields.List[0]
	MetadataNamespace = Metadata.Type.(*ast.StructType).Fields.List[1]

	DataMyKey    = Data.Type.(*ast.StructType).Fields.List[0]
	DataPassword = Data.Type.(*ast.StructType).Fields.List[1]

	SpecPorts    = Spec.Type.(*ast.StructType).Fields.List[0]
	SpecReplicas = Spec.Type.(*ast.StructType).Fields.List[1]
	SpecRules    = Spec.Type.(*ast.StructType).Fields.List[2]
	SpecSelector = Spec.Type.(*ast.StructType).Fields.List[3]
	SpecTemplate = Spec.Type.(*ast.StructType).Fields.List[4]

	SpecTemplateMetadata = SpecTemplate.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateSpec     = SpecTemplate.Type.(*ast.StructType).Fields.List[1]

	SpecTemplateMetadataLabels    = SpecTemplateMetadata.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateMetadataLabelsApp = SpecTemplateMetadataLabels.Type.(*ast.StructType).Fields.List[0]

	SpecTemplateSpecContainers     = SpecTemplateSpec.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateSpecContainersItem = SpecTemplateSpecContainers.Type.(*ast.ArrayType)

	SpecTemplateSpecContainersItemEnvFrom = SpecTemplateSpecContainersItem.Elt.(*ast.StructType).Fields.List[0]
	SpecTemplateSpecContainersItemImage   = SpecTemplateSpecContainersItem.Elt.(*ast.StructType).Fields.List[1]
	SpecTemplateSpecContainersItemName    = SpecTemplateSpecContainersItem.Elt.(*ast.StructType).Fields.List[2]
	SpecTemplateSpecContainersItemPorts   = SpecTemplateSpecContainersItem.Elt.(*ast.StructType).Fields.List[3]

	SpecSelectorMatchLabels    = SpecSelector.Type.(*ast.StructType).Fields.List[0]
	SpecSelectorMatchLabelsApp = SpecSelectorMatchLabels.Type.(*ast.StructType).Fields.List[0]

	SpecPortsItem           = SpecPorts.Type.(*ast.ArrayType)
	SpecPortsItemPort       = SpecPortsItem.Elt.(*ast.StructType).Fields.List[0]
	SpecPortsItemProtocol   = SpecPortsItem.Elt.(*ast.StructType).Fields.List[1]
	SpecPortsItemTargetPort = SpecPortsItem.Elt.(*ast.StructType).Fields.List[2]
)

var Keys = map[ast.Node]string{
	ApiVersion: "apiVersion",
	Data:       "data",
	Kind:       "kind",
	Metadata:   "metadata",
	Spec:       "spec",
	Type:       "type",

	MetadataName:      "name",
	MetadataNamespace: "namespace",

	SpecPorts:    "ports",
	SpecReplicas: "replicas",
	SpecRules:    "rules",
	SpecSelector: "selector",
	SpecTemplate: "template",

	SpecTemplateMetadata: "metadata",
	SpecTemplateSpec:     "spec",

	SpecTemplateMetadataLabels:    "labels",
	SpecTemplateMetadataLabelsApp: "app",

	SpecTemplateSpecContainers: "containers",

	SpecTemplateSpecContainersItemEnvFrom: "envFrom",
	SpecTemplateSpecContainersItemImage:   "image",
	SpecTemplateSpecContainersItemName:    "name",
	SpecTemplateSpecContainersItemPorts:   "ports",

	SpecSelectorMatchLabels:    "matchLabels",
	SpecSelectorMatchLabelsApp: "app",

	SpecPortsItemPort:       "port",
	SpecPortsItemProtocol:   "protocol",
	SpecPortsItemTargetPort: "targetPort",
}
