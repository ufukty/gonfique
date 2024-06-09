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

	t.Fatal("add getter/setter calls")

	fmt.Println(cfg.Gateways.Public.Services.Tags.Endpoints.Assign)
}
