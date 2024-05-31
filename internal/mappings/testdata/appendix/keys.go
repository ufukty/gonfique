package appendix

import "go/ast"

// sorted:
// | apiVersion: apps/v1
// | data:
// |   my-key: my-value
// |   password: cGFzc3dvcmQ=
// | kind: Deployment
// | metadata:
// |   name: my-deployment
// |   namespace: my-namespace
// | spec:
// |   ports:
// |     - port: 80
// |       protocol: TCP
// |       targetPort: 80
// |   replicas: 3
// |   rules:
// |     - host: myapp.example.com
// |       http:
// |         paths:
// |           - backend:
// |               service:
// |                 name: my-service
// |                 port:
// |                   number: 80
// |             path: /
// |             pathType: Prefix
// |   selector:
// |     matchLabels:
// |       app: my-app
// |   template:
// |     metadata:
// |       labels:
// |         app: my-app
// |     spec:
// |       containers:
// |         - envFrom:
// |             - configMapRef:
// |                 name: my-config
// |             - secretRef:
// |                 name: my-secret
// |           image: my-image
// |           name: my-container
// |           ports:
// |             - containerPort: 80
// | type: Opaque

var (
	K8sCfg   = KubernetesExampleOutput.Decls[1].(*ast.GenDecl)
	K8sCfgTs = K8sCfg.Specs[0].(*ast.TypeSpec)
	K8sCfgFs = K8sCfgTs.Type.(*ast.StructType).Fields.List

	ApiVersion                                                = K8sCfgFs[0]
	Data                                                      = K8sCfgFs[1]
	DataMyKey                                                 = Data.Type.(*ast.StructType).Fields.List[0]
	DataPassword                                              = Data.Type.(*ast.StructType).Fields.List[1]
	Kind                                                      = K8sCfgFs[2]
	Metadata                                                  = K8sCfgFs[3]
	MetadataName                                              = Metadata.Type.(*ast.StructType).Fields.List[0]
	MetadataNamespace                                         = Metadata.Type.(*ast.StructType).Fields.List[1]
	Spec                                                      = K8sCfgFs[4]
	SpecPorts                                                 = Spec.Type.(*ast.StructType).Fields.List[0]
	SpecPortsItem                                             = SpecPorts.Type.(*ast.ArrayType)
	SpecPortsItemPort                                         = SpecPortsItem.Elt.(*ast.StructType).Fields.List[0]
	SpecPortsItemProtocol                                     = SpecPortsItem.Elt.(*ast.StructType).Fields.List[1]
	SpecPortsItemTargetPort                                   = SpecPortsItem.Elt.(*ast.StructType).Fields.List[2]
	SpecReplicas                                              = Spec.Type.(*ast.StructType).Fields.List[1]
	SpecRules                                                 = Spec.Type.(*ast.StructType).Fields.List[2]
	SpecRulesItem                                             = SpecRules.Type.(*ast.ArrayType)
	SpecRulesItemHost                                         = SpecRulesItem.Elt.(*ast.StructType).Fields.List[0]
	SpecRulesItemHttp                                         = SpecRulesItem.Elt.(*ast.StructType).Fields.List[1]
	SpecRulesItemHttpPaths                                    = SpecRulesItemHttp.Type.(*ast.StructType).Fields.List[0]
	SpecRulesItemHttpPathsItem                                = SpecRulesItemHttpPaths.Type.(*ast.ArrayType)
	SpecRulesItemHttpPathsItemBackend                         = SpecRulesItemHttpPathsItem.Elt.(*ast.StructType).Fields.List[0]
	SpecRulesItemHttpPathsItemBackendService                  = SpecRulesItemHttpPathsItemBackend.Type.(*ast.StructType).Fields.List[0]
	SpecRulesItemHttpPathsItemBackendServiceName              = SpecRulesItemHttpPathsItemBackendService.Type.(*ast.StructType).Fields.List[0]
	SpecRulesItemHttpPathsItemBackendServicePort              = SpecRulesItemHttpPathsItemBackendService.Type.(*ast.StructType).Fields.List[1]
	SpecRulesItemHttpPathsItemBackendServicePortNumber        = SpecRulesItemHttpPathsItemBackendServicePort.Type.(*ast.StructType).Fields.List[0]
	SpecRulesItemHttpPathsItemPath                            = SpecRulesItemHttpPathsItem.Elt.(*ast.StructType).Fields.List[1]
	SpecRulesItemHttpPathsItemPathType                        = SpecRulesItemHttpPathsItem.Elt.(*ast.StructType).Fields.List[2]
	SpecSelector                                              = Spec.Type.(*ast.StructType).Fields.List[3]
	SpecSelectorMatchLabels                                   = SpecSelector.Type.(*ast.StructType).Fields.List[0]
	SpecSelectorMatchLabelsApp                                = SpecSelectorMatchLabels.Type.(*ast.StructType).Fields.List[0]
	SpecTemplate                                              = Spec.Type.(*ast.StructType).Fields.List[4]
	SpecTemplateMetadata                                      = SpecTemplate.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateMetadataLabels                                = SpecTemplateMetadata.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateMetadataLabelsApp                             = SpecTemplateMetadataLabels.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateSpec                                          = SpecTemplate.Type.(*ast.StructType).Fields.List[1]
	SpecTemplateSpecContainers                                = SpecTemplateSpec.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateSpecContainersItem                            = SpecTemplateSpecContainers.Type.(*ast.ArrayType)
	SpecTemplateSpecContainersItemEnvFrom                     = SpecTemplateSpecContainersItem.Elt.(*ast.StructType).Fields.List[0]
	SpecTemplateSpecContainersItemEnvFromItem                 = SpecTemplateSpecContainersItemEnvFrom.Type.(*ast.ArrayType)
	SpecTemplateSpecContainersItemEnvFromItemConfigMapRef     = SpecTemplateSpecContainersItemEnvFromItem.Elt.(*ast.StructType).Fields.List[0]
	SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName = SpecTemplateSpecContainersItemEnvFromItemConfigMapRef.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateSpecContainersItemEnvFromItemSecretRef        = SpecTemplateSpecContainersItemEnvFromItem.Elt.(*ast.StructType).Fields.List[1]
	SpecTemplateSpecContainersItemEnvFromItemSecretRefName    = SpecTemplateSpecContainersItemEnvFromItemSecretRef.Type.(*ast.StructType).Fields.List[0]
	SpecTemplateSpecContainersItemImage                       = SpecTemplateSpecContainersItem.Elt.(*ast.StructType).Fields.List[1]
	SpecTemplateSpecContainersItemName                        = SpecTemplateSpecContainersItem.Elt.(*ast.StructType).Fields.List[2]
	SpecTemplateSpecContainersItemPorts                       = SpecTemplateSpecContainersItem.Elt.(*ast.StructType).Fields.List[3]
	SpecTemplateSpecContainersItemPortsItem                   = SpecTemplateSpecContainersItemPorts.Type.(*ast.ArrayType)
	SpecTemplateSpecContainersItemPortsItemContainerPort      = SpecTemplateSpecContainersItemPortsItem.Elt.(*ast.StructType).Fields.List[0]
	Type                                                      = K8sCfgFs[5]
)

var Keys = map[ast.Node]string{
	ApiVersion:                               "apiVersion",
	Data:                                     "data",
	DataMyKey:                                "my-key",
	DataPassword:                             "password",
	Kind:                                     "kind",
	Metadata:                                 "metadata",
	MetadataName:                             "name",
	MetadataNamespace:                        "namespace",
	Spec:                                     "spec",
	SpecPorts:                                "ports",
	SpecPortsItemPort:                        "port",
	SpecPortsItemProtocol:                    "protocol",
	SpecPortsItemTargetPort:                  "targetPort",
	SpecReplicas:                             "replicas",
	SpecRules:                                "rules",
	SpecRulesItemHost:                        "host",
	SpecRulesItemHttp:                        "http",
	SpecRulesItemHttpPaths:                   "paths",
	SpecRulesItemHttpPathsItemBackend:        "backend",
	SpecRulesItemHttpPathsItemBackendService: "service",
	SpecRulesItemHttpPathsItemBackendServiceName:              "name",
	SpecRulesItemHttpPathsItemBackendServicePort:              "port",
	SpecRulesItemHttpPathsItemBackendServicePortNumber:        "number",
	SpecRulesItemHttpPathsItemPath:                            "path",
	SpecRulesItemHttpPathsItemPathType:                        "pathType",
	SpecSelector:                                              "selector",
	SpecSelectorMatchLabels:                                   "matchLabels",
	SpecSelectorMatchLabelsApp:                                "app",
	SpecTemplate:                                              "template",
	SpecTemplateMetadata:                                      "metadata",
	SpecTemplateMetadataLabels:                                "labels",
	SpecTemplateMetadataLabelsApp:                             "app",
	SpecTemplateSpec:                                          "spec",
	SpecTemplateSpecContainers:                                "containers",
	SpecTemplateSpecContainersItemEnvFrom:                     "envFrom",
	SpecTemplateSpecContainersItemEnvFromItemConfigMapRef:     "configMapRef",
	SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName: "name",
	SpecTemplateSpecContainersItemEnvFromItemSecretRef:        "secretRef",
	SpecTemplateSpecContainersItemEnvFromItemSecretRefName:    "name",
	SpecTemplateSpecContainersItemImage:                       "image",
	SpecTemplateSpecContainersItemName:                        "name",
	SpecTemplateSpecContainersItemPorts:                       "ports",
	SpecTemplateSpecContainersItemPortsItemContainerPort:      "containerPort",
	Type: "type",
}
