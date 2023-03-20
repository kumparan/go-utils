package utils

import "strings"

// LowerMapStringKey :nodoc:
func LowerMapStringKey(v map[string]interface{}) map[string]interface{} {
	lv := make(map[string]interface{}, len(v))
	for mk, mv := range v {
		lv[strings.ToLower(mk)] = mv
	}
	return lv
}

// MapValuesToOrderedSlice convert map values to ordered slice
// if the map didn't contain one or more key in the order slice, then the key will be skipped
func MapValuesToOrderedSlice[K comparable, V any](order []K, src map[K]V) []V {
	res := make([]V, 0)
	for _, o := range order {
		if v, ok := src[o]; ok {
			res = append(res, v)
			continue
		}
		var zero V
		res = append(res, zero)
	}
	return res
}
