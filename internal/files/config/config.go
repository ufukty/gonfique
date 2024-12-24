package config

import (
	"fmt"
	"go/ast"
	"os"
	"strings"

	"github.com/ufukty/gonfique/internal/files/config/meta"
	"gopkg.in/yaml.v3"
)

type Dict string

const (
	ZeroDict Dict = ""
	Struct   Dict = "struct"
	Map      Dict = "map"
)

type Directives struct {
	Export  bool     `yaml:"export"`
	Declare Typename `yaml:"declare"`
	Dict    Dict     `yaml:"dict"`
	Replace string   `yaml:"replace"`

	Parent    string   `yaml:"parent"`
	Embed     string   `yaml:"embed"`
	Accessors []string `yaml:"accessors"`
	Iterator  bool     `yaml:"iterator"`
}

func (d Directives) IsZero() bool {
	return !d.Export &&
		d.Declare == "" &&
		d.Dict == ZeroDict &&
		d.Replace == "" &&
		d.Parent == "" &&
		d.Embed == "" &&
		d.Accessors == nil &&
		!d.Iterator
}

type Path string

func (p Path) Segments() []string {
	return strings.Split(string(p), ".")
}

func (p Path) WithField(f FieldName) Path {
	return Path(fmt.Sprintf("%s.%s", p, f))
}

type Typename string

func (t Typename) Ident() *ast.Ident {
	return ast.NewIdent(string(t))
}

type FieldName string

func (fn FieldName) Ident() *ast.Ident {
	return ast.NewIdent(string(fn))
}

type File struct {
	Meta  meta.Meta           `yaml:"meta"`
	Rules map[Path]Directives `yaml:"rules"`
}

func Read(src string) (*File, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening: %w", err)
	}
	defer f.Close()
	n := &File{}
	err = yaml.NewDecoder(f).Decode(n)
	if err != nil {
		return nil, fmt.Errorf("decoding: %w", err)
	}
	return n, nil
}
