package sortby

import (
	"cmp"
	"iter"
	"slices"

	"golang.org/x/exp/maps"
)

func Key[K cmp.Ordered, V any](m map[K]V) iter.Seq2[K, V] {
	return KeyFunc(m, cmp.Compare)
}

func KeyFunc[K cmp.Ordered, V any](m map[K]V, cmp func(a, b K) int) iter.Seq2[K, V] {
	ks := maps.Keys(m)
	slices.SortFunc(ks, cmp)
	return func(yield func(K, V) bool) {
		for _, k := range ks {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}
