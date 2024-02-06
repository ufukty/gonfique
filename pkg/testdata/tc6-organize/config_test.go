package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOrganize(t *testing.T) {
	cfg, err := ReadConfig("config.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("reading config: %w", err))
	}

	eps := cfg.Github.Gateways.Public.Services.Tags.Endpoints
	if reflect.TypeOf(eps.Assign) != reflect.TypeOf(eps.Create) {
		t.Fatalf("two endpoints are assigned different types %q & %q", reflect.TypeOf(eps.Assign), reflect.TypeOf(eps.Create))
	}
}
