package directives

import (
	"fmt"
	"testing"
)

func TestDirectiveFile(t *testing.T) {
	df, err := ReadDirectiveFile("testdata/directives.yml")
	if err != nil {
		t.Fatal(fmt.Errorf("prep: %w", err))
	}
	if len(*df) != 2 {
		t.Fatal(fmt.Errorf("assert, length. want %d got %d", 2, len(*df)))
	}
}
