package models

import (
	"fmt"
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

func (kp Keypath) WithField(f FieldName) Keypath {
	return Keypath(fmt.Sprintf("%s.%s", kp, f))
}

func (kp Keypath) Parent() Keypath {
	ss := kp.Segments()
	l := max(len(ss)-1, 0)
	return Keypath(strings.Join(ss[:l], "."))
}

type TypeName string

func (tn TypeName) Ident() *ast.Ident {
	return ast.NewIdent(string(tn))
}

type FieldName string

func (fn FieldName) Ident() *ast.Ident {
	return ast.NewIdent(string(fn))
}
