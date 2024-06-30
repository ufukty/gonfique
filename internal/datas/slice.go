package datas

import "golang.org/x/exp/maps"

func Uniq[K comparable](ss []K) []K {
	m := make(map[K]bool, len(ss))
	for _, s := range ss {
		m[s] = true
	}
	return maps.Keys(m)
}
