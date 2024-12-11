// based on matcher/testdata/appendix
package appendix

import (
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/paths/models"
)

var ConfigType = &ast.StructType{
	Struct: 74,
	Fields: &ast.FieldList{
		List: []*ast.Field{
			&ast.Field{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: "ApiVersion",
					},
				},
				Type: &ast.Ident{
					Name: "string",
				},
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "`yaml:\"apiVersion\"`",
				},
				Comment: nil,
			},
			&ast.Field{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: "Data",
					},
				},
				Type: &ast.StructType{
					Struct: 134,
					Fields: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "MyKey",
									},
								},
								Type: &ast.Ident{
									Name: "string",
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"my-key\"`",
								},
								Comment: nil,
							},
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "Password",
									},
								},
								Type: &ast.Ident{
									Name: "string",
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"password\"`",
								},
								Comment: nil,
							},
						},
					},
					Incomplete: false,
				},
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "`yaml:\"data\"`",
				},
				Comment: nil,
			},
			&ast.Field{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: "Kind",
					},
				},
				Type: &ast.Ident{
					Name: "string",
				},
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "`yaml:\"kind\"`",
				},
				Comment: nil,
			},
			&ast.Field{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: "Metadata",
					},
				},
				Type: &ast.StructType{
					Struct: 271,
					Fields: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "Name",
									},
								},
								Type: &ast.Ident{
									Name: "string",
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"name\"`",
								},
								Comment: nil,
							},
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "Namespace",
									},
								},
								Type: &ast.Ident{
									Name: "string",
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"namespace\"`",
								},
								Comment: nil,
							},
						},
					},
					Incomplete: false,
				},
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "`yaml:\"metadata\"`",
				},
				Comment: nil,
			},
			&ast.Field{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: "Spec",
					},
				},
				Type: &ast.StructType{
					Struct: 378,
					Fields: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "Ports",
									},
								},
								Type: &ast.ArrayType{
									Lbrack: 395,
									Len:    nil,
									Elt: &ast.StructType{
										Struct: 397,
										Fields: &ast.FieldList{
											List: []*ast.Field{
												&ast.Field{
													Names: []*ast.Ident{
														&ast.Ident{
															Name: "Port",
														},
													},
													Type: &ast.Ident{
														Name: "int",
													},
													Tag: &ast.BasicLit{
														Kind:  token.STRING,
														Value: "`yaml:\"port\"`",
													},
													Comment: nil,
												},
												&ast.Field{
													Names: []*ast.Ident{
														&ast.Ident{
															Name: "Protocol",
														},
													},
													Type: &ast.Ident{
														Name: "string",
													},
													Tag: &ast.BasicLit{
														Kind:  token.STRING,
														Value: "`yaml:\"protocol\"`",
													},
													Comment: nil,
												},
												&ast.Field{
													Names: []*ast.Ident{
														&ast.Ident{
															Name: "TargetPort",
														},
													},
													Type: &ast.Ident{
														Name: "int",
													},
													Tag: &ast.BasicLit{
														Kind:  token.STRING,
														Value: "`yaml:\"targetPort\"`",
													},
													Comment: nil,
												},
											},
										},
										Incomplete: false,
									},
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"ports\"`",
								},
								Comment: nil,
							},
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "Replicas",
									},
								},
								Type: &ast.Ident{
									Name: "int",
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"replicas\"`",
								},
								Comment: nil,
							},
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "Rules",
									},
								},
								Type: &ast.ArrayType{
									Lbrack: 584,
									Len:    nil,
									Elt: &ast.StructType{
										Struct: 586,
										Fields: &ast.FieldList{
											List: []*ast.Field{
												&ast.Field{
													Names: []*ast.Ident{
														&ast.Ident{
															Name: "Host",
														},
													},
													Type: &ast.Ident{
														Name: "string",
													},
													Tag: &ast.BasicLit{
														Kind:  token.STRING,
														Value: "`yaml:\"host\"`",
													},
													Comment: nil,
												},
												&ast.Field{
													Names: []*ast.Ident{
														&ast.Ident{
															Name: "Http",
														},
													},
													Type: &ast.StructType{
														Struct: 632,
														Fields: &ast.FieldList{
															List: []*ast.Field{
																&ast.Field{
																	Names: []*ast.Ident{
																		&ast.Ident{
																			Name: "Paths",
																		},
																	},
																	Type: &ast.ArrayType{
																		Lbrack: 651,
																		Len:    nil,
																		Elt: &ast.StructType{
																			Struct: 653,
																			Fields: &ast.FieldList{
																				List: []*ast.Field{
																					&ast.Field{
																						Names: []*ast.Ident{
																							&ast.Ident{
																								Name: "Backend",
																							},
																						},
																						Type: &ast.StructType{
																							Struct: 675,
																							Fields: &ast.FieldList{
																								List: []*ast.Field{
																									&ast.Field{
																										Names: []*ast.Ident{
																											&ast.Ident{
																												Name: "Service",
																											},
																										},
																										Type: &ast.StructType{
																											Struct: 698,
																											Fields: &ast.FieldList{
																												List: []*ast.Field{
																													&ast.Field{
																														Names: []*ast.Ident{
																															&ast.Ident{
																																Name: "Name",
																															},
																														},
																														Type: &ast.Ident{
																															Name: "string",
																														},
																														Tag: &ast.BasicLit{
																															Kind:  token.STRING,
																															Value: "`yaml:\"name\"`",
																														},
																														Comment: nil,
																													},
																													&ast.Field{
																														Names: []*ast.Ident{
																															&ast.Ident{
																																Name: "Port",
																															},
																														},
																														Type: &ast.StructType{
																															Struct: 752,
																															Fields: &ast.FieldList{
																																List: []*ast.Field{
																																	&ast.Field{
																																		Names: []*ast.Ident{
																																			&ast.Ident{
																																				Name: "Number",
																																			},
																																		},
																																		Type: &ast.Ident{
																																			Name: "int",
																																		},
																																		Tag: &ast.BasicLit{
																																			Kind:  token.STRING,
																																			Value: "`yaml:\"number\"`",
																																		},
																																		Comment: nil,
																																	},
																																},
																															},
																															Incomplete: false,
																														},
																														Tag: &ast.BasicLit{
																															Kind:  token.STRING,
																															Value: "`yaml:\"port\"`",
																														},
																														Comment: nil,
																													},
																												},
																											},
																											Incomplete: false,
																										},
																										Tag: &ast.BasicLit{
																											Kind:  token.STRING,
																											Value: "`yaml:\"service\"`",
																										},
																										Comment: nil,
																									},
																								},
																							},
																							Incomplete: false,
																						},
																						Tag: &ast.BasicLit{
																							Kind:  token.STRING,
																							Value: "`yaml:\"backend\"`",
																						},
																						Comment: nil,
																					},
																					&ast.Field{
																						Names: []*ast.Ident{
																							&ast.Ident{
																								Name: "Path",
																							},
																						},
																						Type: &ast.Ident{
																							Name: "string",
																						},
																						Tag: &ast.BasicLit{
																							Kind:  token.STRING,
																							Value: "`yaml:\"path\"`",
																						},
																						Comment: nil,
																					},
																					&ast.Field{
																						Names: []*ast.Ident{
																							&ast.Ident{
																								Name: "PathType",
																							},
																						},
																						Type: &ast.Ident{
																							Name: "string",
																						},
																						Tag: &ast.BasicLit{
																							Kind:  token.STRING,
																							Value: "`yaml:\"pathType\"`",
																						},
																						Comment: nil,
																					},
																				},
																			},
																			Incomplete: false,
																		},
																	},
																	Tag: &ast.BasicLit{
																		Kind:  token.STRING,
																		Value: "`yaml:\"paths\"`",
																	},
																	Comment: nil,
																},
															},
														},
														Incomplete: false,
													},
													Tag: &ast.BasicLit{
														Kind:  token.STRING,
														Value: "`yaml:\"http\"`",
													},
													Comment: nil,
												},
											},
										},
										Incomplete: false,
									},
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"rules\"`",
								},
								Comment: nil,
							},
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "Selector",
									},
								},
								Type: &ast.StructType{
									Struct: 1012,
									Fields: &ast.FieldList{
										List: []*ast.Field{
											&ast.Field{
												Names: []*ast.Ident{
													&ast.Ident{
														Name: "MatchLabels",
													},
												},
												Type: &ast.StructType{
													Struct: 1036,
													Fields: &ast.FieldList{
														List: []*ast.Field{
															&ast.Field{
																Names: []*ast.Ident{
																	&ast.Ident{
																		Name: "App",
																	},
																},
																Type: &ast.Ident{
																	Name: "string",
																},
																Tag: &ast.BasicLit{
																	Kind:  token.STRING,
																	Value: "`yaml:\"app\"`",
																},
																Comment: nil,
															},
														},
													},
													Incomplete: false,
												},
												Tag: &ast.BasicLit{
													Kind:  token.STRING,
													Value: "`yaml:\"matchLabels\"`",
												},
												Comment: nil,
											},
										},
									},
									Incomplete: false,
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"selector\"`",
								},
								Comment: nil,
							},
							&ast.Field{
								Names: []*ast.Ident{
									&ast.Ident{
										Name: "Template",
									},
								},
								Type: &ast.StructType{
									Struct: 1132,
									Fields: &ast.FieldList{
										List: []*ast.Field{
											&ast.Field{
												Names: []*ast.Ident{
													&ast.Ident{
														Name: "Metadata",
													},
												},
												Type: &ast.StructType{
													Struct: 1153,
													Fields: &ast.FieldList{
														List: []*ast.Field{
															&ast.Field{
																Names: []*ast.Ident{
																	&ast.Ident{
																		Name: "Labels",
																	},
																},
																Type: &ast.StructType{
																	Struct: 1173,
																	Fields: &ast.FieldList{
																		List: []*ast.Field{
																			&ast.Field{
																				Names: []*ast.Ident{
																					&ast.Ident{
																						Name: "App",
																					},
																				},
																				Type: &ast.Ident{
																					Name: "string",
																				},
																				Tag: &ast.BasicLit{
																					Kind:  token.STRING,
																					Value: "`yaml:\"app\"`",
																				},
																				Comment: nil,
																			},
																		},
																	},
																	Incomplete: false,
																},
																Tag: &ast.BasicLit{
																	Kind:  token.STRING,
																	Value: "`yaml:\"labels\"`",
																},
																Comment: nil,
															},
														},
													},
													Incomplete: false,
												},
												Tag: &ast.BasicLit{
													Kind:  token.STRING,
													Value: "`yaml:\"metadata\"`",
												},
												Comment: nil,
											},
											&ast.Field{
												Names: []*ast.Ident{
													&ast.Ident{
														Name: "Spec",
													},
												},
												Type: &ast.StructType{
													Struct: 1264,
													Fields: &ast.FieldList{
														List: []*ast.Field{
															&ast.Field{
																Names: []*ast.Ident{
																	&ast.Ident{
																		Name: "Containers",
																	},
																},
																Type: &ast.ArrayType{
																	Lbrack: 1288,
																	Len:    nil,
																	Elt: &ast.StructType{
																		Struct: 1290,
																		Fields: &ast.FieldList{
																			List: []*ast.Field{
																				&ast.Field{
																					Names: []*ast.Ident{
																						&ast.Ident{
																							Name: "EnvFrom",
																						},
																					},
																					Type: &ast.ArrayType{
																						Lbrack: 1312,
																						Len:    nil,
																						Elt: &ast.StructType{
																							Struct: 1314,
																							Fields: &ast.FieldList{
																								List: []*ast.Field{
																									&ast.Field{
																										Names: []*ast.Ident{
																											&ast.Ident{
																												Name: "ConfigMapRef",
																											},
																										},
																										Type: &ast.StructType{
																											Struct: 1342,
																											Fields: &ast.FieldList{
																												List: []*ast.Field{
																													&ast.Field{
																														Names: []*ast.Ident{
																															&ast.Ident{
																																Name: "Name",
																															},
																														},
																														Type: &ast.Ident{
																															Name: "string",
																														},
																														Tag: &ast.BasicLit{
																															Kind:  token.STRING,
																															Value: "`yaml:\"name\"`",
																														},
																														Comment: nil,
																													},
																												},
																											},
																											Incomplete: false,
																										},
																										Tag: &ast.BasicLit{
																											Kind:  token.STRING,
																											Value: "`yaml:\"configMapRef\"`",
																										},
																										Comment: nil,
																									},
																									&ast.Field{
																										Names: []*ast.Ident{
																											&ast.Ident{
																												Name: "SecretRef",
																											},
																										},
																										Type: &ast.StructType{
																											Struct: 1430,
																											Fields: &ast.FieldList{
																												List: []*ast.Field{
																													&ast.Field{
																														Names: []*ast.Ident{
																															&ast.Ident{
																																Name: "Name",
																															},
																														},
																														Type: &ast.Ident{
																															Name: "string",
																														},
																														Tag: &ast.BasicLit{
																															Kind:  token.STRING,
																															Value: "`yaml:\"name\"`",
																														},
																														Comment: nil,
																													},
																												},
																											},
																											Incomplete: false,
																										},
																										Tag: &ast.BasicLit{
																											Kind:  token.STRING,
																											Value: "`yaml:\"secretRef\"`",
																										},
																										Comment: nil,
																									},
																								},
																							},
																							Incomplete: false,
																						},
																					},
																					Tag: &ast.BasicLit{
																						Kind:  token.STRING,
																						Value: "`yaml:\"envFrom\"`",
																					},
																					Comment: nil,
																				},
																				&ast.Field{
																					Names: []*ast.Ident{
																						&ast.Ident{
																							Name: "Image",
																						},
																					},
																					Type: &ast.Ident{
																						Name: "string",
																					},
																					Tag: &ast.BasicLit{
																						Kind:  token.STRING,
																						Value: "`yaml:\"image\"`",
																					},
																					Comment: nil,
																				},
																				&ast.Field{
																					Names: []*ast.Ident{
																						&ast.Ident{
																							Name: "Name",
																						},
																					},
																					Type: &ast.Ident{
																						Name: "string",
																					},
																					Tag: &ast.BasicLit{
																						Kind:  token.STRING,
																						Value: "`yaml:\"name\"`",
																					},
																					Comment: nil,
																				},
																				&ast.Field{
																					Names: []*ast.Ident{
																						&ast.Ident{
																							Name: "Ports",
																						},
																					},
																					Type: &ast.ArrayType{
																						Lbrack: 1599,
																						Len:    nil,
																						Elt: &ast.StructType{
																							Struct: 1601,
																							Fields: &ast.FieldList{
																								List: []*ast.Field{
																									&ast.Field{
																										Names: []*ast.Ident{
																											&ast.Ident{
																												Name: "ContainerPort",
																											},
																										},
																										Type: &ast.Ident{
																											Name: "int",
																										},
																										Tag: &ast.BasicLit{
																											Kind:  token.STRING,
																											Value: "`yaml:\"containerPort\"`",
																										},
																										Comment: nil,
																									},
																								},
																							},
																							Incomplete: false,
																						},
																					},
																					Tag: &ast.BasicLit{
																						Kind:  token.STRING,
																						Value: "`yaml:\"ports\"`",
																					},
																					Comment: nil,
																				},
																			},
																		},
																		Incomplete: false,
																	},
																},
																Tag: &ast.BasicLit{
																	Kind:  token.STRING,
																	Value: "`yaml:\"containers\"`",
																},
																Comment: nil,
															},
														},
													},
													Incomplete: false,
												},
												Tag: &ast.BasicLit{
													Kind:  token.STRING,
													Value: "`yaml:\"spec\"`",
												},
												Comment: nil,
											},
										},
									},
									Incomplete: false,
								},
								Tag: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "`yaml:\"template\"`",
								},
								Comment: nil,
							},
						},
					},
					Incomplete: false,
				},
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "`yaml:\"spec\"`",
				},
				Comment: nil,
			},
			&ast.Field{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: "Type",
					},
				},
				Type: &ast.Ident{
					Name: "string",
				},
				Tag: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "`yaml:\"type\"`",
				},
				Comment: nil,
			},
		},
	},
	Incomplete: false,
}

var (
	K8sCfgFs = ConfigType.Fields.List

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

var Keypaths = map[ast.Node]models.FlattenKeypath{
	ApiVersion:                               "apiVersion",
	Data:                                     "data",
	DataMyKey:                                "data.my-key",
	DataPassword:                             "data.password",
	Kind:                                     "kind",
	Metadata:                                 "metadata",
	MetadataName:                             "metadata.name",
	MetadataNamespace:                        "metadata.namespace",
	Spec:                                     "spec",
	SpecPorts:                                "spec.ports",
	SpecPortsItem:                            "spec.ports.[]",
	SpecPortsItemPort:                        "spec.ports.[].port",
	SpecPortsItemProtocol:                    "spec.ports.[].protocol",
	SpecPortsItemTargetPort:                  "spec.ports.[].targetPort",
	SpecReplicas:                             "spec.replicas",
	SpecRules:                                "spec.rules",
	SpecRulesItem:                            "spec.rules.[]",
	SpecRulesItemHost:                        "spec.rules.[].host",
	SpecRulesItemHttp:                        "spec.rules.[].http",
	SpecRulesItemHttpPaths:                   "spec.rules.[].http.paths",
	SpecRulesItemHttpPathsItem:               "spec.rules.[].http.paths.[]",
	SpecRulesItemHttpPathsItemBackend:        "spec.rules.[].http.paths.[].backend",
	SpecRulesItemHttpPathsItemBackendService: "spec.rules.[].http.paths.[].backend.service",
	SpecRulesItemHttpPathsItemBackendServiceName:              "spec.rules.[].http.paths.[].backend.service.name",
	SpecRulesItemHttpPathsItemBackendServicePort:              "spec.rules.[].http.paths.[].backend.service.port",
	SpecRulesItemHttpPathsItemBackendServicePortNumber:        "spec.rules.[].http.paths.[].backend.service.port.number",
	SpecRulesItemHttpPathsItemPath:                            "spec.rules.[].http.paths.[].path",
	SpecRulesItemHttpPathsItemPathType:                        "spec.rules.[].http.paths.[].pathType",
	SpecSelector:                                              "spec.selector",
	SpecSelectorMatchLabels:                                   "spec.selector.matchLabels",
	SpecSelectorMatchLabelsApp:                                "spec.selector.matchLabels.app",
	SpecTemplate:                                              "spec.template",
	SpecTemplateMetadata:                                      "spec.template.metadata",
	SpecTemplateMetadataLabels:                                "spec.template.metadata.labels",
	SpecTemplateMetadataLabelsApp:                             "spec.template.metadata.labels.app",
	SpecTemplateSpec:                                          "spec.template.spec",
	SpecTemplateSpecContainers:                                "spec.template.spec.containers",
	SpecTemplateSpecContainersItem:                            "spec.template.spec.containers.[]",
	SpecTemplateSpecContainersItemEnvFrom:                     "spec.template.spec.containers.[].envFrom",
	SpecTemplateSpecContainersItemEnvFromItem:                 "spec.template.spec.containers.[].envFrom.[]",
	SpecTemplateSpecContainersItemEnvFromItemConfigMapRef:     "spec.template.spec.containers.[].envFrom.[].configMapRef",
	SpecTemplateSpecContainersItemEnvFromItemConfigMapRefName: "spec.template.spec.containers.[].envFrom.[].configMapRef.name",
	SpecTemplateSpecContainersItemEnvFromItemSecretRef:        "spec.template.spec.containers.[].envFrom.[].secretRef",
	SpecTemplateSpecContainersItemEnvFromItemSecretRefName:    "spec.template.spec.containers.[].envFrom.[].secretRef.name",
	SpecTemplateSpecContainersItemImage:                       "spec.template.spec.containers.[].image",
	SpecTemplateSpecContainersItemName:                        "spec.template.spec.containers.[].name",
	SpecTemplateSpecContainersItemPorts:                       "spec.template.spec.containers.[].ports",
	SpecTemplateSpecContainersItemPortsItem:                   "spec.template.spec.containers.[].ports.[]",
	SpecTemplateSpecContainersItemPortsItemContainerPort:      "spec.template.spec.containers.[].ports.[].containerPort",
	Type: "type",
}
