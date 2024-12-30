package config

import (
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	h, err := os.Open("testdata/config.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("prep, read file: %w", err))
	}
	defer h.Close()
	f, err := Read(h)
	if err != nil {
		t.Fatal(fmt.Errorf("read: %w", err))
	}
	t.Run("b.Export", func(t *testing.T) {
		if f.Rules["b"].Export != true {
			t.Errorf("expected %t, got %t", true, f.Rules["b"].Export)
		}
	})
	t.Run("c.Declare", func(t *testing.T) {
		if f.Rules["c"].Declare != "Employee" {
			t.Errorf("expected %s, got %s", "Employee", f.Rules["c"].Declare)
		}
	})
	t.Run("e.Dict struct", func(t *testing.T) {
		if f.Rules["e"].Dict != "struct" {
			t.Errorf("expected '', got %s", f.Rules["e"].Dict)
		}
	})
	t.Run("g.Dict map", func(t *testing.T) {
		if f.Rules["g"].Dict != "map" {
			t.Errorf("expected 'map', got %s", f.Rules["g"].Dict)
		}
	})
	t.Run("d.Replace", func(t *testing.T) {
		if f.Rules["d"].Replace != "Date time" {
			t.Errorf("expected %s, got %s", "Date time", f.Rules["d"].Replace)
		}
	})
}
