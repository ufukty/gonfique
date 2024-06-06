package namings

import (
	"fmt"
	"testing"

	"github.com/ufukty/gonfique/internal/models"
	"golang.org/x/exp/maps"
)

func TestAutogen(t *testing.T) {
	type testcase map[models.Keypath]models.TypeName

	tcs := map[string]testcase{
		"empty": {}, // empty
		"config": {
			"": "",
		},
		"letters, 1": {
			"a": "a",
			"b": "b",
			"c": "c",
			"d": "d",
			"e": "e",
		},
		"letters, 2": {
			"a.a": "a", // alphabetical order rewarded with generic type name...
			"a.b": "b",
			"b.a": "bA", //
		},
		"letters, 3": {
			"a.a.a": "a",
			"a.b.a": "bA",
		},
		"letters, 4": {
			"a.a.a": "a",
			"a.b.a": "bA",
			"b.a.a": "aA",
		},
		"words, 1": {
			"lorem.ipsum.dolor": "dolor",
			"sit.amet":          "amet",
			"consectetur":       "consectetur",
		},
		"words, 2": { // conflicting on different levels
			"lorem.ipsum.dolor": "ipsumDolor",
			"sit.amet":          "amet",
			"consectetur":       "consectetur",
			"dolor":             "dolor",
		},
		"words, 3": { // conflicting on different levels
			"lorem.dolor":       "dolor",
			"lorem.ipsum.dolor": "ipsumDolor",
			"lorem.ipsum.sit":   "ipsumSit",
			"lorem.ipsum":       "ipsum",
			"lorem.sit":         "sit",
		},
		"words, 4": { // conflicting on different levels
			"lorem.ipsum": "loremIpsum",
			"ipsum":       "ipsum",
		},
	}

	for tn, tc := range tcs {
		t.Run(tn, func(t *testing.T) {
			got := GenerateTypenames(maps.Keys(tc))

			fmt.Printf("%-20s %-10s %s\n", "Keypath", "Want", "Got")
			for kp, tn := range tc {
				fmt.Printf("%-20s %-10s %s\n", kp, tn, got[kp])
			}

			if len(got) != len(tc) {
				t.Fatalf("assert 1, length. want %d got %d", len(tc), len(got))
			}

			for in, want := range tc {
				if _, ok := got[in]; !ok {
					t.Fatalf("assert 2, existence. outputs doesn't contain the input and expected output %q: %q", in, want)
				}
				if got[in] != want {
					t.Errorf("assert 3, value. want %q got %q", want, got[in])
				}
			}

		})
	}
}
