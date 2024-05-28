package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := ReadConfig("config.json")
	if err != nil {
		t.Fatal(fmt.Errorf("reading config: %w", err))
	}

	if len(cfg.Digitalocean.Fra1.Services.Account) != 1 {
		t.Fatal("expected 1 item")
	}

	if len(cfg.Digitalocean.Fra1.Services.Account[0].SshKeys) != 0 {
		t.Fatal("expected empty array")
	}

}
