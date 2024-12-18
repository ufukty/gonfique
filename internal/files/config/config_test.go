package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	f, err := Read("testdata/config.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("read: %w", err))
	}
	t.Run("b.Export", func(t *testing.T) {
		if f.Paths["b"].Export != true {
			t.Errorf("expected %t, got %t", true, f.Paths["b"].Export)
		}
	})
	t.Run("c.Declare", func(t *testing.T) {
		if f.Paths["c"].Declare != "Employee" {
			t.Errorf("expected %s, got %s", "Employee", f.Paths["c"].Declare)
		}
	})
	t.Run("e.Dict struct (default value = zero value)", func(t *testing.T) {
		if f.Paths["e"].Dict != "" {
			t.Errorf("expected '', got %s", f.Paths["e"].Dict)
		}
	})
	t.Run("g.Dict map", func(t *testing.T) {
		if f.Paths["g"].Dict != "map" {
			t.Errorf("expected 'map', got %s", f.Paths["g"].Dict)
		}
	})
	t.Run("d.Replace", func(t *testing.T) {
		if f.Paths["d"].Replace != "Date time" {
			t.Errorf("expected %s, got %s", "Date time", f.Paths["d"].Replace)
		}
	})
}
