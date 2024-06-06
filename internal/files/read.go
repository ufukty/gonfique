package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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

func ReadConfigFile(src string) (any, models.Encoding, error) {
	switch ext := filepath.Ext(src); ext {
	case ".json":
		cfgcontent, err := readJsonConfig(src)
		if err != nil {
			return nil, models.Json, fmt.Errorf("reading json config: %w", err)
		}
		return cfgcontent, models.Json, nil
	case ".yaml", ".yml":
		cfgcontent, err := readYamlConfig(src)
		if err != nil {
			return nil, models.Yaml, fmt.Errorf("reading yaml config: %w", err)
		}
		return cfgcontent, models.Yaml, nil
	default:
		return nil, models.Yaml, fmt.Errorf("unsupported file extension %q", ext)
	}
}
