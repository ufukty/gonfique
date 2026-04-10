package sortby

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

func Key[K cmp.Ordered, V any](m map[K]V) iter.Seq2[K, V] {
	return KeyFunc(m, cmp.Compare)
}

func KeyFunc[K cmp.Ordered, V any](m map[K]V, cmp func(a, b K) int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range slices.SortedFunc(maps.Keys(m), cmp) {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}
