package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// Int64ToString :nodoc:
func Int64ToString(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

// Offset to get offset from page and limit, min value for page = 1
func Offset(page, limit int64) int64 {
	offset := (page - 1) * limit
	if offset < 0 {
		return 0
	}
	return offset
}

// GenerateID based on current time
func GenerateID() int64 {
	return time.Now().UnixNano() + int64(rand.Intn(10000))
}

// Int32PointerToInt64 :nodoc:
func Int32PointerToInt64(i *int32) int64 {
	if i != nil {
		return int64(*i)
	}
	return int64(0)
}

// Int32PointerToInt32 :nodoc:
func Int32PointerToInt32(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

// Int64PointerToInt64 :nodoc:
func Int64PointerToInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// IsSameSliceIgnoreOrder to compare slice without order
func IsSameSliceIgnoreOrder(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	diff := make(map[int64]bool, len(a))
	for _, v := range a {
		diff[v] = true
	}
	for _, v := range b {
		if _, ok := diff[v]; !ok {
			return false
		}
		delete(diff, v)
	}

	return len(diff) == 0
}

// Int64WithLimit -> Check req value bigger or not from limit.
func Int64WithLimit(input int64, limit int64) int64 {
	if input < 0 || input > limit {
		return limit
	}

	return input
}

// Int64WithMinAndMaxLimit check input value. if bigger than max, then return max. if smaller than min, then return min. else return input.
func Int64WithMinAndMaxLimit(input, min, max int64) int64 {
	if input < min {
		return min
	}

	if input > max {
		return max
	}

	return input
}
