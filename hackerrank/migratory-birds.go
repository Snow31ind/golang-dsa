package hackerrank

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func Entries[M ~map[K]V, K comparable, V any](m M) []struct {
	Key   K
	Value V
} {
	r := make([]struct {
		Key   K
		Value V
	}, 0, len(m))
	for k, v := range m {
		r = append(r, struct {
			Key   K
			Value V
		}{Key: k, Value: v})
	}
	return r
}

func MigratoryBirds(arr []int32) int32 {
	m := make(map[int32]int32)
	for _, v := range arr {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}

	entries := Entries(m)
	if len(entries) == 1 {
		return entries[0].Key
	}

	key := entries[0].Key
	value := entries[0].Value
	for _, entry := range entries[1:] {
		if entry.Value > value || (entry.Value == value && entry.Key < key) {
			value = entry.Value
			key = entry.Key
		}
	}

	return key
}
