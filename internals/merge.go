package utils

import "reflect"

// MergeStringSliceToMap ...
func MergeStringSliceToMap(m map[string][]interface{}, k string, v []interface{}) {
	if m[k] == nil {
		m[k] = make([]interface{}, len(v))
		for i, s := range v {
			m[k][i] = interface{}(s)
		}
	} else {
		m[k] = append(m[k], v...)
	}
}

func MergeStructs(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.FieldByName(srcVal.Type().Field(i).Name)

		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(srcField)
		}
	}
}
