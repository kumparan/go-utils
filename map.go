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

// MapValuesToOrderedSlice convert map values to ordered slice.
// If the map didn't contain one or more key in the keysOrder slice, then it will be filled with zero value.
func MapValuesToOrderedSlice[K comparable, V any](src map[K]V, keysOrder []K) (orderedMapValues []V) {
	var zero V
	for _, o := range keysOrder {
		if v, ok := src[o]; ok {
			orderedMapValues = append(orderedMapValues, v)
			continue
		}
		orderedMapValues = append(orderedMapValues, zero)
	}
	return
}

// MapValuesToOrderedSliceExistOnly convert map values to ordered slice.
// If the map didn't contain one or more key in the keysOrder slice, then it will be ignored.
func MapValuesToOrderedSliceExistOnly[K comparable, V any](src map[K]V, keysOrder []K) (orderedMapValues []V) {
	for _, o := range keysOrder {
		if v, ok := src[o]; ok {
			orderedMapValues = append(orderedMapValues, v)
		}
	}
	return
}
