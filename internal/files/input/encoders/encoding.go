package encoders

import (
	"fmt"
	"path/filepath"
)

type Encoding string

var (
	Json = Encoding("json")
	Yaml = Encoding("yaml")
)

func FromExtension(path string) (Encoding, error) {
	switch ext := filepath.Ext(path); ext {
	case ".json":
		return Json, nil
	case ".yml", ".yaml":
		return Yaml, nil
	default:
		return "", fmt.Errorf("unknown file extension: %s", ext)
	}
}

func FromString(s string) (Encoding, error) {
	switch s {
	case "json":
		return Json, nil
	case "yml", "yaml":
		return Yaml, nil
	default:
		return "", fmt.Errorf("unknown identifier: %s", s)
	}
}
