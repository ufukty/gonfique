package namings

import (
	"fmt"
	"testing"

	"github.com/ufukty/gonfique/internal/paths/models"
)

func TestAutogen(t *testing.T) {
	type testcase struct {
		input  map[models.FlattenKeypath]bool
		output map[models.FlattenKeypath]models.TypeName
	}

	tcs := map[string]testcase{
		"empty": {
			input:  map[models.FlattenKeypath]bool{},
			output: map[models.FlattenKeypath]models.TypeName{},
		},
		// TODO: consider producing args.TypeName for empty keypath
		// "config": {
		// 	input: map[models.FlattenKeypath]bool{
		// 		"": false,
		// 	},
		// 	output: map[models.FlattenKeypath]models.TypeName{
		// 		"": "Config",
		// 	},
		// },
		"letters, 1": {
			input: map[models.FlattenKeypath]bool{
				"a": false,
				"b": false,
				"c": false,
				"d": false,
				"e": false,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"a": "a",
				"b": "b",
				"c": "c",
				"d": "d",
				"e": "e",
			},
		},
		"letters, 1, exported": {
			input: map[models.FlattenKeypath]bool{
				"a": true,
				"b": true,
				"c": true,
				"d": true,
				"e": true,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"a": "A",
				"b": "B",
				"c": "C",
				"d": "D",
				"e": "E",
			},
		},
		"letters, 2": {
			input: map[models.FlattenKeypath]bool{
				"a.a": false,
				"a.b": false,
				"b.a": false,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"a.a": "a", // alphabetical order rewarded with generic type name...
				"a.b": "b",
				"b.a": "bA", //
			},
		},
		"letters, 3": {
			input: map[models.FlattenKeypath]bool{
				"a.a.a": false,
				"a.b.a": false,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"a.a.a": "a",
				"a.b.a": "bA",
			},
		},
		"letters, 3, exported": {
			input: map[models.FlattenKeypath]bool{
				"a.a.a": true,
				"a.b.a": true,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"a.a.a": "A",
				"a.b.a": "BA",
			},
		},
		"letters, 4": {
			input: map[models.FlattenKeypath]bool{
				"a.a.a": false,
				"a.b.a": false,
				"b.a.a": false,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"a.a.a": "a",
				"a.b.a": "bA",
				"b.a.a": "aA",
			},
		},
		"letters, 4, exported": {
			input: map[models.FlattenKeypath]bool{
				"a.a.a": true,
				"a.b.a": true,
				"b.a.a": true,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"a.a.a": "A",
				"a.b.a": "BA",
				"b.a.a": "AA",
			},
		},
		"words, 1": {
			input: map[models.FlattenKeypath]bool{
				"lorem.ipsum.dolor": false,
				"sit.amet":          false,
				"consectetur":       false,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"lorem.ipsum.dolor": "dolor",
				"sit.amet":          "amet",
				"consectetur":       "consectetur",
			},
		},
		"words, 2": { // conflicting on different leveltestcases
			input: map[models.FlattenKeypath]bool{
				"lorem.ipsum.dolor": false,
				"sit.amet":          false,
				"consectetur":       false,
				"dolor":             false,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"lorem.ipsum.dolor": "ipsumDolor",
				"sit.amet":          "amet",
				"consectetur":       "consectetur",
				"dolor":             "dolor",
			},
		},
		"words, 3": { // conflicting on different leveltestcases
			input: map[models.FlattenKeypath]bool{
				"lorem.dolor":       false,
				"lorem.ipsum.dolor": false,
				"lorem.ipsum.sit":   false,
				"lorem.ipsum":       false,
				"lorem.sit":         false,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"lorem.dolor":       "dolor",
				"lorem.ipsum.dolor": "ipsumDolor",
				"lorem.ipsum.sit":   "ipsumSit",
				"lorem.ipsum":       "ipsum",
				"lorem.sit":         "sit",
			},
		},
		"words, 4": { // conflicting on different leveltestcases
			input: map[models.FlattenKeypath]bool{
				"lorem.ipsum": false,
				"ipsum":       false,
			},
			output: map[models.FlattenKeypath]models.TypeName{
				"lorem.ipsum": "loremIpsum",
				"ipsum":       "ipsum",
			},
		},
	}

	for tn, tc := range tcs {
		t.Run(tn, func(t *testing.T) {
			got := GenerateTypenames(tc.input)

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
