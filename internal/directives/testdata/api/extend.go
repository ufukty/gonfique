package config

import (
	"path/filepath"
	"strings"
)

type Addressable interface {
	GetPath() string
}

func Join(adds ...Addressable) string {
	addr := []string{}
	for _, adr := range adds {
		strings.Join(addr, adr.GetPath())
	}
	return filepath.Join(addr...)
}
