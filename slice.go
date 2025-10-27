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

// FindDifferencesFromSlices find item that not exists in all slices but exists in one or more of them
func FindDifferencesFromSlices[T comparable](slices ...[]T) []T {
	if len(slices) < 2 {
		return nil
	}

	var allItems []T
	itemCountMap := map[T]int{}
	for _, slice := range slices {
		handledItem := map[T]bool{}
		for _, item := range slice {
			// handle duplicate item in one slice
			if !handledItem[item] {
				itemCountMap[item]++
				allItems = append(allItems, item)
				handledItem[item] = true
			}
		}
	}
	allItems = Unique(allItems)
	var result []T
	for _, item := range allItems {
		count := itemCountMap[item]
		if count < len(slices) {
			result = append(result, item)
		}
	}

	return result
}

// IsUniqueSliceItem :nodoc:
func IsUniqueSliceItem[T comparable](data []T) bool {
	mapData := make(map[T]bool, len(data))
	for _, d := range data {
		if _, ok := mapData[d]; ok {
			return false
		}
		mapData[d] = true
	}

	return true
}

// ConvertSlice can change slice data type or even manipulate the data using converter func
func ConvertSlice[T1 any, T2 any](in []T1, converter func(T1) T2) []T2 {
	var res []T2
	for _, v := range in {
		res = append(res, converter(v))
	}
	return res
}

// ReverseSlice :nodoc:
func ReverseSlice[T any](in []T) []T {
	reversed := make([]T, 0)
	for i := len(in) - 1; i >= 0; i-- {
		reversed = append(reversed, in[i])
	}
	return reversed
}
