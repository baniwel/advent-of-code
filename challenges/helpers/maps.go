package helpers

func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, V any](m map[K]V) []any {
	values := make([]any, 0, len(m))
	for k := range m {
		values = append(values, m[k])
	}
	return values
}
