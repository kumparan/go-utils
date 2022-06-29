package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Int32PointerToInt64(t *testing.T) {
	var i *int32
	assert.Equal(t, int64(0), Int32PointerToInt64(i))
	ii := int32(12)
	i = &ii
	assert.Equal(t, int64(ii), Int32PointerToInt64(i))
	*i = 0
	assert.Equal(t, int64(0), Int32PointerToInt64(i))
}

func Test_Int32PointerToInt32(t *testing.T) {
	var i *int32
	assert.Equal(t, int32(0), Int32PointerToInt32(i))
	ii := int32(12)
	i = &ii
	assert.Equal(t, ii, Int32PointerToInt32(i))
	*i = 0
	assert.Equal(t, int32(0), Int32PointerToInt32(i))
}

func Test_Int64PointerToInt64(t *testing.T) {
	var i *int64
	assert.Equal(t, int64(0), Int64PointerToInt64(i))
	ii := int64(12345678901)
	i = &ii
	assert.Equal(t, ii, Int64PointerToInt64(i))
	*i = 0
	assert.Equal(t, int64(0), Int64PointerToInt64(i))
}

func Test_IsSameSliceIgnoreOrder(t *testing.T) {
	a := []int64{2, 1, 3}
	b := []int64{2, 1, 3}
	assert.True(t, IsSameSliceIgnoreOrder(a, b))
	a = []int64{2, 1, 3}
	b = []int64{1, 2, 3}
	assert.True(t, IsSameSliceIgnoreOrder(a, b))
	a = []int64{2, 1, 3, 4}
	b = []int64{1, 2, 3}
	assert.False(t, IsSameSliceIgnoreOrder(a, b))
	a = []int64{}
	b = []int64{}
	assert.True(t, IsSameSliceIgnoreOrder(a, b))
}

func Test_Int64WithLimit(t *testing.T) {
	a := int64(5)
	b := int64(10)
	c := int64(15)
	d := int64(-1)

	assert.Equal(t, a, Int64WithLimit(b, a))
	assert.Equal(t, b, Int64WithLimit(b, c))
	assert.Equal(t, a, Int64WithLimit(d, a))
	assert.NotEqual(t, c, Int64WithLimit(c, a))

}

func Test_Int64WithMinAndMaxLimit(t *testing.T) {
	var (
		min int64 = 1
		max int64 = 25
		a   int64 = 5
		b   int64
		c   int64 = 26
	)

	assert.Equal(t, a, Int64WithMinAndMaxLimit(a, min, max))
	assert.Equal(t, min, Int64WithMinAndMaxLimit(b, min, max))
	assert.Equal(t, max, Int64WithMinAndMaxLimit(c, min, max))
}
