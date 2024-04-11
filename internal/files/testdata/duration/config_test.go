package config

import (
	"fmt"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	cfg, err := ReadConfig("config.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("reading config: %w", err))
	}

	var d time.Duration = cfg.Github.Gateways.Public.GracePeriod
	if d != time.Microsecond*200 {
		t.Fatal("not 200µs")
	}
}
