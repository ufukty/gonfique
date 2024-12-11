package files

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	f, err := ReadConfigFile("testdata/config.yml")
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
	t.Run("e.Dict static", func(t *testing.T) {
		if f.Paths["e"].Dict != "static" {
			t.Errorf("expected 'static', got %s", f.Paths["e"].Dict)
		}
	})
	t.Run("f.Dict dynamic-keys", func(t *testing.T) {
		if f.Paths["f"].Dict != "dynamic-keys" {
			t.Errorf("expected 'dynamic-keys', got %s", f.Paths["f"].Dict)
		}
	})
	t.Run("g.Dict dynamic", func(t *testing.T) {
		if f.Paths["g"].Dict != "dynamic" {
			t.Errorf("expected 'dynamic', got %s", f.Paths["g"].Dict)
		}
	})
	t.Run("d.Replace", func(t *testing.T) {
		if f.Paths["d"].Replace != "Date time" {
			t.Errorf("expected %s, got %s", "Date time", f.Paths["d"].Replace)
		}
	})
}
