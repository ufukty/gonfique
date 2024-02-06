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

	if cfg.Github.Domain != "github.com" {
		t.Fatal(fmt.Errorf("expected %q got %q", "github.com", cfg.Github.Domain))
	}

	for key, value := range cfg.Github.Gateways.Public.Services.Tags.Endpoints.Range() {
		fmt.Println(key, value.Method) // enough to not have syntax or runtime error
	}
}
