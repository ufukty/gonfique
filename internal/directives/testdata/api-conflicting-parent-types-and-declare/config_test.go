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

	want := "/api/v1.0.0/tags/assign"
	got := Join(
		cfg.Gateways.Public,
		cfg.Gateways.Public.Services.Tags,
		cfg.Gateways.Public.Services.Tags.Endpoints.Assign,
	)
	if want != got {
		t.Fatal(fmt.Errorf("assert:\nwant: %s\ngot : %s", want, got))
	}
}
