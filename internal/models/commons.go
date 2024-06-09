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

type WildcardKeypath string

func (kp WildcardKeypath) Segments() []string {
	return strings.Split(string(kp), ".")
}

func (kp WildcardKeypath) WithField(f FieldName) WildcardKeypath {
	return WildcardKeypath(fmt.Sprintf("%s.%s", kp, f))
}

func (kp WildcardKeypath) Parent() WildcardKeypath {
	ss := kp.Segments()
	l := max(len(ss)-1, 0)
	return WildcardKeypath(strings.Join(ss[:l], "."))
}

type FlattenKeypath string

func (kp FlattenKeypath) Segments() []string {
	return strings.Split(string(kp), ".")
}

func (kp FlattenKeypath) WithField(f FieldName) FlattenKeypath {
	return FlattenKeypath(fmt.Sprintf("%s.%s", kp, f))
}

func (kp FlattenKeypath) Parent() FlattenKeypath {
	ss := kp.Segments()
	l := max(len(ss)-1, 0)
	return FlattenKeypath(strings.Join(ss[:l], "."))
}

type TypeName string

func (tn TypeName) Ident() *ast.Ident {
	return ast.NewIdent(string(tn))
}

type FieldName string

func (fn FieldName) Ident() *ast.Ident {
	return ast.NewIdent(string(fn))
}
