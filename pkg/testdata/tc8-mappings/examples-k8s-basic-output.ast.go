package tc8

import (
	"go/ast"
	"go/token"
)

var KubernetesExampleOutput = &ast.File{
    Name: &ast.Ident{
        Name: "config",
    },
    Decls: []ast.Decl{
        &ast.GenDecl{
            Tok: token.IMPORT,
            Specs: []ast.Spec{
                &ast.ImportSpec{
                    Name: nil,
                    Path: &ast.BasicLit{
                        Kind: token.STRING,
                        Value: "\"fmt\"",
                    },
                    Comment: nil,
                },
                &ast.ImportSpec{
                    Name: nil,
                    Path: &ast.BasicLit{
                        Kind: token.STRING,
                        Value: "\"os\"",
                    },
                    Comment: nil,
                },
                &ast.ImportSpec{
                    Name: nil,
                    Path: &ast.BasicLit{
                        Kind: token.STRING,
                        Value: "\"gopkg.in/yaml.v3\"",
                    },
                    Comment: nil,
                },
            },
        },
        &ast.GenDecl{
            Tok: token.TYPE,
            Specs: []ast.Spec{
                &ast.TypeSpec{
                    Name: &ast.Ident{
                        Name: "Config",
                    },
                    TypeParams: nil,
                    Assign: 0,
                    Type: &ast.StructType{
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
                                        Kind: token.STRING,
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
                                                        Kind: token.STRING,
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
                                                        Kind: token.STRING,
                                                        Value: "`yaml:\"password\"`",
                                                    },
                                                    Comment: nil,
                                                },
                                            },
                                        },
                                        Incomplete: false,
                                    },
                                    Tag: &ast.BasicLit{
                                        Kind: token.STRING,
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
                                        Kind: token.STRING,
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
                                                        Kind: token.STRING,
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
                                                        Kind: token.STRING,
                                                        Value: "`yaml:\"namespace\"`",
                                                    },
                                                    Comment: nil,
                                                },
                                            },
                                        },
                                        Incomplete: false,
                                    },
                                    Tag: &ast.BasicLit{
                                        Kind: token.STRING,
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
                                                        Len: nil,
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
                                                                            Kind: token.STRING,
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
                                                                            Kind: token.STRING,
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
                                                                            Kind: token.STRING,
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
                                                        Kind: token.STRING,
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
                                                        Kind: token.STRING,
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
                                                        Len: nil,
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
                                                                            Kind: token.STRING,
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
                                                                                            Len: nil,
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
                                                                                                                                                Kind: token.STRING,
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
                                                                                                                                                                Kind: token.STRING,
                                                                                                                                                                Value: "`yaml:\"number\"`",
                                                                                                                                                            },
                                                                                                                                                            Comment: nil,
                                                                                                                                                        },
                                                                                                                                                    },
                                                                                                                                                },
                                                                                                                                                Incomplete: false,
                                                                                                                                            },
                                                                                                                                            Tag: &ast.BasicLit{
                                                                                                                                                Kind: token.STRING,
                                                                                                                                                Value: "`yaml:\"port\"`",
                                                                                                                                            },
                                                                                                                                            Comment: nil,
                                                                                                                                        },
                                                                                                                                    },
                                                                                                                                },
                                                                                                                                Incomplete: false,
                                                                                                                            },
                                                                                                                            Tag: &ast.BasicLit{
                                                                                                                                Kind: token.STRING,
                                                                                                                                Value: "`yaml:\"service\"`",
                                                                                                                            },
                                                                                                                            Comment: nil,
                                                                                                                        },
                                                                                                                    },
                                                                                                                },
                                                                                                                Incomplete: false,
                                                                                                            },
                                                                                                            Tag: &ast.BasicLit{
                                                                                                                Kind: token.STRING,
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
                                                                                                                Kind: token.STRING,
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
                                                                                                                Kind: token.STRING,
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
                                                                                            Kind: token.STRING,
                                                                                            Value: "`yaml:\"paths\"`",
                                                                                        },
                                                                                        Comment: nil,
                                                                                    },
                                                                                },
                                                                            },
                                                                            Incomplete: false,
                                                                        },
                                                                        Tag: &ast.BasicLit{
                                                                            Kind: token.STRING,
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
                                                        Kind: token.STRING,
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
                                                                                        Kind: token.STRING,
                                                                                        Value: "`yaml:\"app\"`",
                                                                                    },
                                                                                    Comment: nil,
                                                                                },
                                                                            },
                                                                        },
                                                                        Incomplete: false,
                                                                    },
                                                                    Tag: &ast.BasicLit{
                                                                        Kind: token.STRING,
                                                                        Value: "`yaml:\"matchLabels\"`",
                                                                    },
                                                                    Comment: nil,
                                                                },
                                                            },
                                                        },
                                                        Incomplete: false,
                                                    },
                                                    Tag: &ast.BasicLit{
                                                        Kind: token.STRING,
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
                                                                                                        Kind: token.STRING,
                                                                                                        Value: "`yaml:\"app\"`",
                                                                                                    },
                                                                                                    Comment: nil,
                                                                                                },
                                                                                            },
                                                                                        },
                                                                                        Incomplete: false,
                                                                                    },
                                                                                    Tag: &ast.BasicLit{
                                                                                        Kind: token.STRING,
                                                                                        Value: "`yaml:\"labels\"`",
                                                                                    },
                                                                                    Comment: nil,
                                                                                },
                                                                            },
                                                                        },
                                                                        Incomplete: false,
                                                                    },
                                                                    Tag: &ast.BasicLit{
                                                                        Kind: token.STRING,
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
                                                                                        Len: nil,
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
                                                                                                            Len: nil,
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
                                                                                                                                                Kind: token.STRING,
                                                                                                                                                Value: "`yaml:\"name\"`",
                                                                                                                                            },
                                                                                                                                            Comment: nil,
                                                                                                                                        },
                                                                                                                                    },
                                                                                                                                },
                                                                                                                                Incomplete: false,
                                                                                                                            },
                                                                                                                            Tag: &ast.BasicLit{
                                                                                                                                Kind: token.STRING,
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
                                                                                                                                                Kind: token.STRING,
                                                                                                                                                Value: "`yaml:\"name\"`",
                                                                                                                                            },
                                                                                                                                            Comment: nil,
                                                                                                                                        },
                                                                                                                                    },
                                                                                                                                },
                                                                                                                                Incomplete: false,
                                                                                                                            },
                                                                                                                            Tag: &ast.BasicLit{
                                                                                                                                Kind: token.STRING,
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
                                                                                                            Kind: token.STRING,
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
                                                                                                            Kind: token.STRING,
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
                                                                                                            Kind: token.STRING,
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
                                                                                                            Len: nil,
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
                                                                                                                                Kind: token.STRING,
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
                                                                                                            Kind: token.STRING,
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
                                                                                        Kind: token.STRING,
                                                                                        Value: "`yaml:\"containers\"`",
                                                                                    },
                                                                                    Comment: nil,
                                                                                },
                                                                            },
                                                                        },
                                                                        Incomplete: false,
                                                                    },
                                                                    Tag: &ast.BasicLit{
                                                                        Kind: token.STRING,
                                                                        Value: "`yaml:\"spec\"`",
                                                                    },
                                                                    Comment: nil,
                                                                },
                                                            },
                                                        },
                                                        Incomplete: false,
                                                    },
                                                    Tag: &ast.BasicLit{
                                                        Kind: token.STRING,
                                                        Value: "`yaml:\"template\"`",
                                                    },
                                                    Comment: nil,
                                                },
                                            },
                                        },
                                        Incomplete: false,
                                    },
                                    Tag: &ast.BasicLit{
                                        Kind: token.STRING,
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
                                        Kind: token.STRING,
                                        Value: "`yaml:\"type\"`",
                                    },
                                    Comment: nil,
                                },
                            },
                        },
                        Incomplete: false,
                    },
                    Comment: nil,
                },
            },
        },
        &ast.FuncDecl{
            Recv: nil,
            Name: &ast.Ident{
                Name: "ReadConfig",
            },
            Type: &ast.FuncType{
                TypeParams: nil,
                Params: &ast.FieldList{
                    List: []*ast.Field{
                        &ast.Field{
                            Names: []*ast.Ident{
                                &ast.Ident{
                                    Name: "path",
                                },
                            },
                            Type: &ast.Ident{
                                Name: "string",
                            },
                            Tag: nil,
                            Comment: nil,
                        },
                    },
                },
                Results: &ast.FieldList{
                    List: []*ast.Field{
                        &ast.Field{
                            Names: nil,
                            Type: &ast.Ident{
                                Name: "Config",
                            },
                            Tag: nil,
                            Comment: nil,
                        },
                        &ast.Field{
                            Names: nil,
                            Type: &ast.Ident{
                                Name: "error",
                            },
                            Tag: nil,
                            Comment: nil,
                        },
                    },
                },
            },
            Body: &ast.BlockStmt{
                List: []ast.Stmt{
                    &ast.AssignStmt{
                        Lhs: []ast.Expr{
                            &ast.Ident{
                                Name: "f",
                            },
                            &ast.Ident{
                                Name: "err",
                            },
                        },
                        Tok: token.DEFINE,
                        Rhs: []ast.Expr{
                            &ast.CallExpr{
                                Fun: &ast.SelectorExpr{
                                    X: &ast.Ident{
                                        Name: "os",
                                    },
                                    Sel: &ast.Ident{
                                        Name: "Open",
                                    },
                                },
                                Args: []ast.Expr{
                                    &ast.Ident{
                                        Name: "path",
                                    },
                                },
                            },
                        },
                    },
                    &ast.IfStmt{
                        If: 1866,
                        Init: nil,
                        Cond: &ast.BinaryExpr{
                            X: &ast.Ident{
                                Name: "err",
                            },
                            OpPos: 1873,
                            Op: token.NEQ,
                            Y: &ast.Ident{
                                Name: "nil",
                            },
                        },
                        Body: &ast.BlockStmt{
                            List: []ast.Stmt{
                                &ast.ReturnStmt{
                                    Results: []ast.Expr{
                                        &ast.CompositeLit{
                                            Type: &ast.Ident{
                                                Name: "Config",
                                            },
                                            Elts: nil,
                                            Incomplete: false,
                                        },
                                        &ast.CallExpr{
                                            Fun: &ast.SelectorExpr{
                                                X: &ast.Ident{
                                                    Name: "fmt",
                                                },
                                                Sel: &ast.Ident{
                                                    Name: "Errorf",
                                                },
                                            },
                                            Args: []ast.Expr{
                                                &ast.BasicLit{
                                                    Kind: token.STRING,
                                                    Value: "\"opening config file: %w\"",
                                                },
                                                &ast.Ident{
                                                    Name: "err",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        Else: nil,
                    },
                    &ast.AssignStmt{
                        Lhs: []ast.Expr{
                            &ast.Ident{
                                Name: "cfg",
                            },
                        },
                        Tok: token.DEFINE,
                        Rhs: []ast.Expr{
                            &ast.CompositeLit{
                                Type: &ast.Ident{
                                    Name: "Config",
                                },
                                Elts: nil,
                                Incomplete: false,
                            },
                        },
                    },
                    &ast.AssignStmt{
                        Lhs: []ast.Expr{
                            &ast.Ident{
                                Name: "err",
                            },
                        },
                        Tok: token.ASSIGN,
                        Rhs: []ast.Expr{
                            &ast.CallExpr{
                                Fun: &ast.SelectorExpr{
                                    X: &ast.CallExpr{
                                        Fun: &ast.SelectorExpr{
                                            X: &ast.Ident{
                                                Name: "yaml",
                                            },
                                            Sel: &ast.Ident{
                                                Name: "NewDecoder",
                                            },
                                        },
                                        Args: []ast.Expr{
                                            &ast.Ident{
                                                Name: "f",
                                            },
                                        },
                                    },
                                    Sel: &ast.Ident{
                                        Name: "Decode",
                                    },
                                },
                                Args: []ast.Expr{
                                    &ast.UnaryExpr{
                                        OpPos: 1997,
                                        Op: token.AND,
                                        X: &ast.Ident{
                                            Name: "cfg",
                                        },
                                    },
                                },
                            },
                        },
                    },
                    &ast.IfStmt{
                        If: 2004,
                        Init: nil,
                        Cond: &ast.BinaryExpr{
                            X: &ast.Ident{
                                Name: "err",
                            },
                            OpPos: 2011,
                            Op: token.NEQ,
                            Y: &ast.Ident{
                                Name: "nil",
                            },
                        },
                        Body: &ast.BlockStmt{
                            List: []ast.Stmt{
                                &ast.ReturnStmt{
                                    Results: []ast.Expr{
                                        &ast.CompositeLit{
                                            Type: &ast.Ident{
                                                Name: "Config",
                                            },
                                            Elts: nil,
                                            Incomplete: false,
                                        },
                                        &ast.CallExpr{
                                            Fun: &ast.SelectorExpr{
                                                X: &ast.Ident{
                                                    Name: "fmt",
                                                },
                                                Sel: &ast.Ident{
                                                    Name: "Errorf",
                                                },
                                            },
                                            Args: []ast.Expr{
                                                &ast.BasicLit{
                                                    Kind: token.STRING,
                                                    Value: "\"decoding config file: %w\"",
                                                },
                                                &ast.Ident{
                                                    Name: "err",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        Else: nil,
                    },
                    &ast.ReturnStmt{
                        Results: []ast.Expr{
                            &ast.Ident{
                                Name: "cfg",
                            },
                            &ast.Ident{
                                Name: "nil",
                            },
                        },
                    },
                },
            },
        },
    },
    Comments: nil,
}