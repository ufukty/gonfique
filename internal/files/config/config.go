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
	Struct Dict = "struct"
	Map    Dict = "map"
)

type PathConfig struct {
	Export  bool     `yaml:"export"`
	Declare Typename `yaml:"declare"`
	Dict    Dict     `yaml:"dict"`
	Replace string   `yaml:"replace"`
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

func defaults(f *File) {
	for cp, pc := range f.Paths {
		if pc.Dict == "" {
			d := f.Paths[cp]
			d.Dict = Struct
			f.Paths[cp] = d
		}
	}
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
	defaults(n)
	return n, nil
}
