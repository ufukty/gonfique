package config

import (
	"fmt"
	"testing"
)

func TestOrganize(t *testing.T) {
	cfg, err := ReadConfig("config.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("reading config: %w", err))
	}

	fmt.Println(cfg.Foo.Baz)
}
