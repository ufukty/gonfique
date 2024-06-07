package resolver

import (
	"testing"

	"github.com/ufukty/gonfique/internal/resolver/testdata/appendix"
)

func TestAllKeypathsForHolders(t *testing.T) {
	got := AllKeypathsForHolders(appendix.ConfigType, appendix.Keys)
	if len(got) != len(appendix.Keypaths) {
		t.Errorf("assert 1, length. want %d got %d", len(appendix.Keys), len(got))
	}
	for holder, wantkp := range appendix.Keypaths {
		if gotkp, ok := got[holder]; !ok {
			t.Errorf("assert 2, existence. want %q", wantkp)
		} else if gotkp != wantkp {
			t.Errorf("assert 3, mismatch. \nwant %q\ngot  %q", wantkp, gotkp)
		}
	}
}
