package auto

import (
	"fmt"
	"testing"

	"github.com/ufukty/gonfique/internal/files/config"
	"github.com/ufukty/gonfique/internal/paths/resolve"
	"golang.org/x/exp/maps"
)

func TestAutogen(t *testing.T) {
	type testcase struct {
		testname string
		output   map[resolve.Path]config.Typename // input set is [maps.Keys]
	}

	tcs := []testcase{
		{
			testname: "empty",
			output:   map[resolve.Path]config.Typename{},
		},
		// TODO: consider producing args.TypeName for empty keypath
		// {
		// 	testname: "config",
		// 	output: map[paths.FlattenKeypath]config.Typename{
		// 		"": "Config",
		// 	},
		// },
		{
			testname: "1 letters",
			output: map[resolve.Path]config.Typename{
				"a": "A",
				"b": "B",
				"c": "C",
				"d": "D",
				"e": "E",
			},
		},
		{
			testname: "2 letters",
			output: map[resolve.Path]config.Typename{
				"a.a": "A", // alphabetical order rewarded with generic type name...
				"a.b": "B",
				"b.a": "BA", //
			},
		},
		{
			testname: "3 letters",
			output: map[resolve.Path]config.Typename{
				"a.a.a": "A",
				"a.b.a": "BA",
			},
		},
		{
			testname: "4 letters",
			output: map[resolve.Path]config.Typename{
				"a.a.a.a": "A",
				"a.b.a.a": "AA",
				"b.a.a.a": "AAA",
			},
		},
		{
			testname: "1 word",
			output: map[resolve.Path]config.Typename{
				"lorem.ipsum.dolor": "Dolor",
				"sit.amet":          "Amet",
				"consectetur":       "Consectetur",
			},
		},
		{
			testname: "2 words conflicting on different leveltestcases",
			output: map[resolve.Path]config.Typename{
				"lorem.ipsum.dolor": "IpsumDolor",
				"sit.amet":          "Amet",
				"consectetur":       "Consectetur",
				"dolor":             "Dolor",
			},
		},
		{
			testname: "3 words conflicting on different leveltestcases",
			output: map[resolve.Path]config.Typename{
				"lorem.dolor":       "Dolor",
				"lorem.ipsum.dolor": "IpsumDolor",
				"lorem.ipsum.sit":   "IpsumSit",
				"lorem.ipsum":       "Ipsum",
				"lorem.sit":         "Sit",
			},
		},
		{
			testname: "4 words conflicting on different leveltestcases",
			output: map[resolve.Path]config.Typename{
				"lorem.ipsum": "LoremIpsum",
				"ipsum":       "Ipsum",
			},
		},
		{
			testname: "specials",
			output: map[resolve.Path]config.Typename{
				"ipsum":         "Ipsum",
				"ipsum.[]":      "IpsumItem",
				"lorem":         "Lorem",
				"lorem.[key]":   "LoremKey",
				"lorem.[value]": "LoremValue",
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.testname, func(t *testing.T) {
			got := GenerateTypenames(maps.Keys(tc.output), []config.Typename{})

			fmt.Printf("%-20s %-10s %s\n", "Keypath", "Want", "Got")
			for kp, tn := range tc.output {
				fmt.Printf("%-20s %-10s %s\n", kp, tn, got[kp])
			}

			if len(got) != len(tc.output) {
				t.Fatalf("assert 1, length. want %d (%v) got %d (%v)", len(tc.output), tc.output, len(got), got)
			}

			for kp, tn := range tc.output {
				if _, ok := got[kp]; !ok {
					t.Fatalf("assert 2, existence. outputs doesn't contain the input and expected output %q: %q", kp, tn)
				}
				if got[kp] != tn {
					t.Errorf("assert 3, value. want %q got %q", tn, got[kp])
				}
			}

		})
	}
}
