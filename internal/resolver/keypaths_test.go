package resolver

import (
	"testing"

	"github.com/ufukty/gonfique/internal/bundle"
	"github.com/ufukty/gonfique/internal/resolver/testdata/appendix"
)

func TestAllKeypathsForHolders(t *testing.T) {
	b := bundle.Bundle{
		OriginalKeys: appendix.Keys,
		CfgType:      appendix.ConfigType,
	}
	got := AllKeypathsForHolders(b)
	for holder, wantkp := range appendix.Keypaths {
		if gotkp, ok := got[holder]; !ok {
			t.Errorf("assert 1, existance")
		} else if gotkp != wantkp {
			t.Errorf("assert 2, mismatch. \nwant %q\ngot  %q", wantkp, gotkp)
		}
	}
}
