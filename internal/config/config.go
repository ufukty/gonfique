package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Dict string

const (
	Static      Dict = "static"
	DynamicKeys Dict = "dynamic-keys"
	Dynamic     Dict = "dynamic"
)

type Meta struct {
	Package string `yaml:"package"`
	Type    string `yaml:"config-type"`
}

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

type File struct {
	Meta  Meta                  `yaml:"meta"`
	Paths map[string]PathConfig `yaml:"paths"`
	Types map[string]TypeConfig `yaml:"types"`
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
