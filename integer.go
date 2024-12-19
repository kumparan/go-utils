package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"
)

// Int64ToString :nodoc:
func Int64ToString(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}

// Offset to get offset from page and limit, minimum value for page = 1
func Offset(page, limit int64) int64 {
	offset := (page - 1) * limit
	if offset < 0 {
		return 0
	}
	return offset
}

// GenerateID based on current time
func GenerateID() int64 {
	now := time.Now().UnixNano()
	randomInt, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		return now
	}

	return now + randomInt.Int64()
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

// Int64WithMinAndMaxLimit check input value. if bigger than maximum, then return maximum. if smaller than minimum, then return minimum. else return input.
func Int64WithMinAndMaxLimit(input, minimum, maximum int64) int64 {
	if input < minimum {
		return minimum
	}

	if input > maximum {
		return maximum
	}

	return input
}
