package utils

// MergeStringSliceToMap merges a string slice to a map
// with a string key and a slice of interfaces. If the
// key does not exist, it creates a new key-value pair.
// If the key exists, it appends the slice to the existing
// slice.
func MergeStringSliceToMap(m map[string][]interface{}, k string, v []interface{}) {
	if m[k] == nil {
		m[k] = make([]interface{}, len(v))
		copy(m[k], v)
	} else {
		m[k] = append(m[k], v...)
	}
}
