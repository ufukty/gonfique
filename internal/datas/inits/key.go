package inits

func Key[K comparable, V any](m map[K]V, k K) {
	if _, ok := m[k]; !ok {
		m[k] = *new(V)
	}
}

func Key2[K1, K2 comparable, V any](m map[K1]map[K2]V, k1 K1, k2 K2) {
	if _, ok := m[k1]; !ok {
		m[k1] = make(map[K2]V)
	}
	Key(m[k1], k2)
}

func Key3[K1, K2, K3 comparable, V any](m map[K1]map[K2]map[K3]V, k1 K1, k2 K2, k3 K3) {
	if _, ok := m[k1]; !ok {
		m[k1] = make(map[K2]map[K3]V)
	}
	Key2(m[k1], k2, k3)
}
