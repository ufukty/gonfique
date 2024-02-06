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

	if cfg.Spec.Ports[0].Port != 80 {
		t.Fatal(fmt.Errorf("array access, expected %q got %q", "80", cfg.Spec.Ports[0].Port))
	}

	if l := len(cfg.Spec.Template.Spec.Containers[0].EnvFrom); l != 2 {
		t.Fatal(fmt.Errorf("array lentgh, expected %q got %q", "2", l))
	}
}
