package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/ufukty/gonfique/internal/transform"
	"gopkg.in/yaml.v3"
)

func readYamlConfig(src string) (*File, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer f.Close()
	var y any
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decoding input file: %w", err)
	}

	cfg, imports := transform.Transform(y)
	imports = slices.Concat([]string{"fmt", "os", "gopkg.in/yaml.v3"}, imports)
	slices.Sort(imports)

	file := &File{
		Lang:          Yaml,
		ConfigContent: y,
		Cfg:           cfg,
		Named:         nil,
		Isolated:      nil,
		Iterators:     nil,
		Imports:       imports,
	}

	return file, nil
}

func readJsonConfig(src string) (*File, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer f.Close()
	var y any
	if err := json.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decoding input file: %w", err)
	}

	cfg, imports := transform.Transform(y)
	imports = slices.Concat([]string{"fmt", "os", "encoding/json"}, imports)
	slices.Sort(imports)

	file := &File{
		Lang:          Json,
		ConfigContent: y,
		Cfg:           cfg,
		Named:         nil,
		Isolated:      nil,
		Iterators:     nil,
		Imports:       imports,
	}

	return file, nil
}

func ReadConfigFile(src string) (*File, error) {
	switch ext := filepath.Ext(src); ext {
	case ".json":
		return readJsonConfig(src)
	case ".yaml", ".yml":
		return readYamlConfig(src)
	default:
		return nil, fmt.Errorf("unsupported file extension %q", ext)
	}
}
