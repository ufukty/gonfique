package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/ufukty/gonfique/internal/transform"
	"gopkg.in/yaml.v3"
)

func readYamlConfig(src string, typename string) (*File, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer f.Close()
	var y any
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decoding input file: %w", err)
	}

	cfg, imports, keys := transform.Transform(y, transform.Yaml)
	imports = slices.Concat([]string{"fmt", "os", "gopkg.in/yaml.v3"}, imports)
	slices.Sort(imports)

	file := &File{
		Encoding:        transform.Yaml,
		Keys:            keys,
		TypeName:        typename,
		TypeNameInitial: strings.ToLower(string(([]rune(typename))[0])),
		Cfg:             cfg,
		Named:           nil,
		Isolated:        nil,
		Iterators:       nil,
		Imports:         imports,
	}

	return file, nil
}

func readJsonConfig(src string, typename string) (*File, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer f.Close()
	var y any
	if err := json.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decoding input file: %w", err)
	}

	cfg, imports, keys := transform.Transform(y, transform.Json)
	imports = slices.Concat([]string{"fmt", "os", "encoding/json"}, imports)
	slices.Sort(imports)

	file := &File{
		Encoding:        transform.Json,
		Keys:            keys,
		TypeName:        typename,
		TypeNameInitial: strings.ToLower(string(([]rune(typename))[0])),
		Cfg:             cfg,
		Named:           nil,
		Isolated:        nil,
		Iterators:       nil,
		Imports:         imports,
	}

	return file, nil
}

func ReadConfigFile(src string, typename string) (*File, error) {
	switch ext := filepath.Ext(src); ext {
	case ".json":
		return readJsonConfig(src, typename)
	case ".yaml", ".yml":
		return readYamlConfig(src, typename)
	default:
		return nil, fmt.Errorf("unsupported file extension %q", ext)
	}
}
