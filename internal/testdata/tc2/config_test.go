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

	if cfg.Logging.FileConfig.Path != "/var/log/app.log" {
		t.Fatal(fmt.Errorf("expected %q got %q", "/var/log/app.log", cfg.Logging.FileConfig.Path))

	}
}
