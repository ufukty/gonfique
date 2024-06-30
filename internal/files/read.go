package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/models"
	"gopkg.in/yaml.v3"
)

func readYamlConfig(src string) (any, error) {
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

func readJsonConfig(src string) (any, error) {
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

func ReadConfigFile(b *bundle.Bundle, src string) error {
	var err error
	switch ext := filepath.Ext(src); ext {
	case ".json":
		b.Encoding = models.Json
		b.Imports = append(b.Imports, "encoding/json")
		b.Cfgcontent, err = readJsonConfig(src)
		if err != nil {
			return fmt.Errorf("reading json config: %w", err)
		}
		return nil
	case ".yaml", ".yml":
		b.Encoding = models.Yaml
		b.Imports = append(b.Imports, "gopkg.in/yaml.v3")
		b.Cfgcontent, err = readYamlConfig(src)
		if err != nil {
			return fmt.Errorf("reading yaml config: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported file extension %q", ext)
	}
}
