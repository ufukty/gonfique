package config

import (
	"fmt"
	"go/ast"
	"os"
	"strings"

	"github.com/ufukty/gonfique/internal/files/config/meta"
	"github.com/ufukty/gonfique/internal/transform"
	"gopkg.in/yaml.v3"
)

type Dict string

const (
	Static      Dict = "static"
	DynamicKeys Dict = "dynamic-keys"
	Dynamic     Dict = "dynamic"
)

type PathConfig struct {
	Export  bool
	Declare string
	Dict    Dict
	Replace string
}

type TypeConfig struct {
	Parent    string   `yaml:"parent"`
	Embed     string   `yaml:"embed"`
	Accessors []string `yaml:"accessors"`
	Iterator  bool     `yaml:"iterator"`
}

type Path string

func (p Path) Segments() []string {
	return strings.Split(string(p), ".")
}

func (p Path) WithField(f transform.FieldName) Path {
	return Path(fmt.Sprintf("%s.%s", p, f))
}

type Typename string

func (t Typename) Ident() *ast.Ident {
	return ast.NewIdent(string(t))
}

type File struct {
	Meta  meta.Meta               `yaml:"meta"`
	Paths map[Path]PathConfig     `yaml:"paths"`
	Types map[Typename]TypeConfig `yaml:"types"`
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