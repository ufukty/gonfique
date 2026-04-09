package version

import (
	"testing"
)

func TestOfBuild(t *testing.T) {
	v, err := OfBuild()
	if err != nil {
		t.Errorf("act, unexpected error: %v", err)
	}
	if v != "(devel)" {
		t.Errorf("assert, unexpected empty value")
	}
}
