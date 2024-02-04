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

	if cfg.Spec.Selector.MatchLabels.App != "my-app" {
		t.Fatal(fmt.Errorf("expected %q got %q", "my-app", cfg.Spec.Selector.MatchLabels.App))
	}
}
