package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := ReadConfig("config.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("reading config: %w", err))
	}

	for _, container := range cfg.Spec.Template.Spec.Containers {
		fmt.Println(container.Name)
	}
}
