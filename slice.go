package utils

import (
	"strconv"
)

// ContainsInt64 tells whether a slice contains x.
func ContainsInt64(a []int64, x int64) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// ContainsString tells whether a slice contains x.
func ContainsString(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// SliceAtoi -> convert array of string to array of integer
func SliceAtoi(s []string) ([]int, error) {
	var arr []int

	for _, val := range s {
		i, err := strconv.Atoi(val)
		if err != nil {
			return arr, err
		}

		arr = append(arr, i)
	}
	return arr, nil
}

// DifferenceString :nodoc:
func DifferenceString(slice1 []string, slice2 []string) []string {
	var diff []string

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

// DifferenceInt64 :nodoc:
func DifferenceInt64(slice1 []int64, slice2 []int64) []int64 {
	var diff []int64

	// Loop two times, first to find slice1 int64 not in slice2,
	// second loop to find slice2 int64 not in slice1
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

// UniqueString :nodoc:
func UniqueString(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for idx := range elements {
		if _, ok := encountered[elements[idx]]; ok {
			continue
		}
		encountered[elements[idx]] = true
		result = append(result, elements[idx])
	}

	return result
}

// UniqueInt64 :nodoc:
func UniqueInt64(elements []int64) []int64 {
	encountered := map[int64]bool{}
	result := []int64{}

	for idx := range elements {
		if encountered[elements[idx]] {
			continue
		}
		encountered[elements[idx]] = true
		result = append(result, elements[idx])
	}

	return result
}

// SlicePointerInt32PointerToSliceInt64 :nodoc:
func SlicePointerInt32PointerToSliceInt64(i *[]*int32) (result []int64) {
	if i != nil {
		dump := *i
		for _, element := range dump {
			result = append(result, int64(*element))
		}
		return result
	}
	return result
}

// PaginateSlice :nodoc:
func PaginateSlice[T comparable](data []T, page, size int64) []T {
	if page < 1 || size < 1 {
		return nil
	}

	offset := Offset(page, size)
	count := len(data)
	switch {
	case count < int(offset):
		return []T{}
	case count < int(offset+size):
		return data[offset:count]
	default:
		return data[offset : offset+size]
	}
}
