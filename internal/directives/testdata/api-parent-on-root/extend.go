package config

import (
	"path/filepath"
)

type Addressable interface {
	GetPath() string
}

func Join(adds ...Addressable) string {
	addr := []string{}
	for _, adr := range adds {
		addr = append(addr, adr.GetPath())
	}
	return filepath.Join(addr...)
}
