package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Float64PointerToFloat64(t *testing.T) {
	var f *float64
	assert.Equal(t, 0.0, Float64PointerToFloat64(f))
	ff := 12.1
	f = &ff
	assert.Equal(t, ff, Float64PointerToFloat64(f))
	*f = 0
	assert.Equal(t, float64(0), Float64PointerToFloat64(f))
}
