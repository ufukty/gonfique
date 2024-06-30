package datas

func Invmap[K, V comparable](src map[K]V) map[V]K {
	dst := make(map[V]K, len(src))
	for k, v := range src {
		dst[v] = k
	}
	return dst
}

func Revmap[K, V comparable](src map[K]V) map[V][]K {
	dst := make(map[V][]K, len(src))
	for k, v := range src {
		if _, ok := dst[v]; !ok {
			dst[v] = []K{}
		}
		dst[v] = append(dst[v], k)
	}
	return dst
}

func RevSliceMap[K, V comparable](src map[K][]V) map[V][]K {
	dst := make(map[V][]K, len(src))
	for k, vs := range src {
		for _, v := range vs {
			if _, ok := dst[v]; !ok {
				dst[v] = []K{}
			}
			dst[v] = append(dst[v], k)
		}
	}
	return dst
}

func MergeMaps[K comparable, V any](a, b map[K]V) map[K]V {
	c := make(map[K]V, len(a)+len(b))
	for k, v := range a {
		c[k] = v
	}
	for k, v := range b {
		c[k] = v
	}
	return c
}
