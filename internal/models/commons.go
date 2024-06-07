package models

import (
	"go/ast"
	"strings"
)

type Encoding string

var (
	Json = Encoding("json")
	Yaml = Encoding("yaml")
)

type Keypath string

func (kp Keypath) Segments() []string {
	return strings.Split(string(kp), ".")
}

type TypeName string

func (tn TypeName) Ident() *ast.Ident {
	return ast.NewIdent(string(tn))
}

type FieldName string

func (fn FieldName) Ident() *ast.Ident {
	return ast.NewIdent(string(fn))
}
