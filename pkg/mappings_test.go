package pkg

import (
	"fmt"
	"testing"
)

func TestMappings(t *testing.T) {

	cts, err := ReadConfigYaml("testdata/tc8-mappings/config.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("resolving the type spec needed: %w", err))
	}

	ms, err := ReadMappings("testdata/tc8-mappings/mappings.gonfique.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("reading mappings from user-provided file: %w", err))
	}

	// apply mappings before "organize" & "iterate"
	Mappings(cts, ms)
}
