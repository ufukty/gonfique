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

	if len(cfg.Droplet.SshKeys) != 0 {
		t.Fatal("expected empty array")
	}
}
