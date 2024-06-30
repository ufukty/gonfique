package config

import "path/filepath"

type addressable interface {
	GetPath() string
}

func Join(adds ...addressable) string {
	addr := []string{}
	for _, adr := range adds {
		addr = append(addr, adr.GetPath())
	}
	return filepath.Join(addr...)
}
