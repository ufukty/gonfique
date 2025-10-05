package input

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"go.ufukty.com/gonfique/internal/files/input/encoders"
	"gopkg.in/yaml.v3"
)

func readYaml(f io.Reader) (any, error) {
	var y any
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	return y, nil
}

func readJson(f io.Reader) (any, error) {
	var y any
	if err := json.NewDecoder(f).Decode(&y); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	return y, nil
}

func Read(f io.Reader, enc encoders.Encoding) (any, error) {
	switch enc {
	case encoders.Json:
		return readJson(f)
	case encoders.Yaml:
		return readYaml(f)
	}
	return nil, fmt.Errorf("unknown encoding: %s", enc)
}

func ReadFile(path string, enc encoders.Encoding) (any, error) {
	h, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer h.Close()
	c, err := Read(h, enc)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	return c, nil
}
