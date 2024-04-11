package files

import (
	"fmt"
	"os"
	"slices"

	"github.com/ufukty/gonfique/internal/transform"
	"gopkg.in/yaml.v3"
)

func ReadConfigYaml(src string) (*File, error) {
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
		ConfigContent: y,
		Cfg:           cfg,
		Named:         nil,
		Isolated:      nil,
		Iterators:     nil,
		Imports:       imports,
	}

	return file, nil
}
