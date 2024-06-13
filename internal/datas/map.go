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
