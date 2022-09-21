package utils

import "encoding/json"

// Contains tells whether slice A contains x.
func Contains[T comparable](a []T, x T) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Difference get discrepancies between 2 slices
func Difference[T comparable](slice1 []T, slice2 []T) []T {
	var diff []T

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

// Unique returns unique value in as slice
func Unique[T comparable](elements []T) (result []T) {
	encountered := map[T]bool{}
	for idx := range elements {
		if _, ok := encountered[elements[idx]]; ok {
			continue
		}
		encountered[elements[idx]] = true
		result = append(result, elements[idx])
	}

	return result
}

// InterfaceBytesToType will transform cached value that get from the redis to any types
func InterfaceBytesToType[T any](i interface{}) (out T) {
	if i == nil {
		return
	}
	bt := i.([]byte)

	_ = json.Unmarshal(bt, &out)
	return
}

// ValueOrDefault use the given value or use default value if the value = empty value
func ValueOrDefault[T comparable](value, defaultValue T) T {
	var emptyValue T

	if value == emptyValue {
		return defaultValue
	}

	return value
}
