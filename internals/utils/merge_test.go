package utils_test

import (
	"testing"
)

// TestMergeStringSliceToMap tests the MergeStringSliceToMap function.
func TestMergeStringSliceToMap(t *testing.T) {
	t.Run("MergeStringSliceToMap", func(t *testing.T) {
		t.Parallel()
		m := make(map[string][]interface{})
		k := "key"
		v := []interface{}{"value"}
		MergeStringSliceToMap(m, k, v)
		if len(m) != 1 {
			t.Errorf("Expected 1, got %d", len(m))
		}
		if len(m[k]) != 1 {
			t.Errorf("Expected 1, got %d", len(m[k]))
		}
		if m[k][0] != "value" {
			t.Errorf("Expected value, got %s", m[k][0])
		}
	})
}

func MergeStringSliceToMap(m map[string][]interface{}, k string, v []interface{}) {
	if m[k] == nil {
		m[k] = make([]interface{}, len(v))
		copy(m[k], v)
	} else {
		m[k] = append(m[k], v...)
	}
}
