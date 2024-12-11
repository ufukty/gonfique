package input

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Encoding string

var (
	Json = Encoding("json")
	Yaml = Encoding("yaml")
)

func readYaml(src string) (any, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer f.Close()
	var y any
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decoding input file: %w", err)
	}
	return y, nil
}

func readJson(src string) (any, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer f.Close()
	var y any
	if err := json.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decoding input file: %w", err)
	}
	return y, nil
}

func Read(src string) (any, Encoding, error) {
	switch ext := filepath.Ext(src); ext {
	case ".json":
		c, err := readJson(src)
		if err != nil {
			return nil, "", fmt.Errorf("json: %w", err)
		}
		return c, Json, nil
	case ".yaml", ".yml":
		c, err := readYaml(src)
		if err != nil {
			return nil, "", fmt.Errorf("yaml: %w", err)
		}
		return c, Yaml, nil
	default:
		return nil, "", fmt.Errorf("unsupported file extension %q", ext)
	}
}
