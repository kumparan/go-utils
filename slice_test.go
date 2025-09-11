package utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContainsInt64(t *testing.T) {
	s := []int64{1, 2, 3, 4, 5}

	assert.Equal(t, true, ContainsInt64(s, int64(1)))
	assert.Equal(t, true, ContainsInt64(s, int64(2)))
	assert.Equal(t, true, ContainsInt64(s, int64(3)))
	assert.Equal(t, true, ContainsInt64(s, int64(4)))
	assert.Equal(t, false, ContainsInt64(s, int64(6)))
	assert.Equal(t, false, ContainsInt64(s, int64(10)))
}

func Test_ContainsString(t *testing.T) {
	s := []string{"kuda", "horse", "flower"}

	assert.Equal(t, true, ContainsString(s, "kuda"))
	assert.Equal(t, true, ContainsString(s, "horse"))
	assert.Equal(t, true, ContainsString(s, "flower"))
	assert.Equal(t, false, ContainsString(s, "house"))
	assert.Equal(t, false, ContainsString(s, "rainbos"))
}

func Test_SliceAtoi(t *testing.T) {
	s := []string{"1", "2", "3"}

	i, err := SliceAtoi(s)
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, i)

	s2 := []string{"kuda", "horse", "flower"}
	_, err = SliceAtoi(s2)
	assert.Error(t, err)
}

func Test_DifferenceString(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"c", "d", "e"}

	assert.Equal(t, []string{"a", "b", "d", "e"}, DifferenceString(s1, s2))
	assert.Equal(t, []string{"d", "e", "a", "b"}, DifferenceString(s2, s1))
}

func Test_DifferenceInt64(t *testing.T) {
	s1 := []int64{1, 2, 3}
	s2 := []int64{3, 4, 5}

	assert.Equal(t, []int64{1, 2, 4, 5}, DifferenceInt64(s1, s2))
	assert.Equal(t, []int64{4, 5, 1, 2}, DifferenceInt64(s2, s1))
}

func Test_UniqueString(t *testing.T) {
	s := []string{"a", "a", "b", "a", "d", "b"}

	assert.Equal(t, []string{"a", "b", "d"}, UniqueString(s))
}

func Test_UniqueInt64(t *testing.T) {
	s := []int64{1, 1, 2, 4, 2}

	assert.Equal(t, []int64{1, 2, 4}, UniqueInt64(s))
}

func Test_SlicePointerInt32PointerToSliceInt64(t *testing.T) {
	var i *[]*int32
	var j []int64
	assert.Equal(t, j, SlicePointerInt32PointerToSliceInt64(i))
}

func Test_PaginateSlice(t *testing.T) {
	slice1 := []string{"a", "a", "b", "a", "d", "b"}
	assert.Equal(t, []string{"a", "a", "b"}, PaginateSlice(slice1, 1, 3))

	slice2 := []int64{1, 1, 2, 4, 2}
	assert.Equal(t, []int64{1, 1, 2}, PaginateSlice(slice2, 1, 3))

	type dummy struct{ a int8 }
	slice3 := []dummy{{a: 1}, {a: 2}, {a: 3}, {a: 4}}
	assert.Equal(t, []dummy{{a: 1}, {a: 2}, {a: 3}}, PaginateSlice(slice3, 1, 3))

	// offset are too high, return empty
	slice4 := []int64{1, 1, 2, 4, 2}
	assert.Equal(t, []int64{}, PaginateSlice(slice4, 10, 3))

	// secondPage
	slice5 := []int64{1, 1, 2, 4, 2}
	assert.Equal(t, []int64{4, 2}, PaginateSlice(slice5, 2, 3))

	// invalid page size input, return nil
	assert.Equal(t, []int64(nil), PaginateSlice(slice5, -123123, -4242))
}

func Test_FindDifferencesFromSlices(t *testing.T) {
	type tc struct {
		slices [][]string
		result []string
	}

	testCases := []tc{
		{
			slices: [][]string{
				{"a", "b", "c"},
				{"a", "b"},
			},
			result: []string{"c"},
		},
		{
			slices: [][]string{
				{"a", "b", "c"},
				{"c", "d", "e"},
				{"e", "f", "g"},
			},
			result: []string{"a", "b", "c", "d", "e", "f", "g"},
		},
		{
			slices: [][]string{
				{"a", "b", "c"},
				{"a", "b", "c"},
				{"a", "b", "c"},
			},
			result: nil,
		},
		{
			slices: [][]string{
				{"a", "b", "c"},
			},
			result: nil,
		},
	}

	for _, tc := range testCases {
		res := FindDifferencesFromSlices(tc.slices...)
		for _, it := range tc.result {
			assert.Contains(t, res, it)
		}
	}
}

func Test_IsUniqueSliceItem(t *testing.T) {
	type tc[T comparable] struct {
		slice  []T
		result bool
	}

	testCasesString := []tc[string]{
		{
			slice: []string{
				"haha", "hehe",
			},
			result: true,
		},
		{
			slice: []string{
				"HAHA", "haha",
			},
			result: true,
		},
		{
			slice: []string{
				"haha", "haha",
			},
			result: false,
		},
	}

	for _, tc := range testCasesString {
		assert.Equal(t, tc.result, IsUniqueSliceItem(tc.slice))
	}

	testCasesInt := []tc[int]{
		{
			slice: []int{
				1, 2,
			},
			result: true,
		},
		{
			slice: []int{
				1, 1,
			},
			result: false,
		},
	}

	for _, tc := range testCasesInt {
		assert.Equal(t, tc.result, IsUniqueSliceItem(tc.slice))
	}
}

func Test_ConvertSliceType(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		type (
			a struct {
				B string
			}
			b struct {
				A string
			}
		)

		srcSlices := []a{{B: "1. this is B inside a"}, {B: "2. this is B inside a"}}
		converter := func(in a) b {
			return b{A: in.B}
		}
		var dstSlices []b
		for _, v := range srcSlices {
			dstSlices = append(dstSlices, converter(v))
		}

		res := ConvertSlice[a, b](srcSlices, converter)
		assert.Equal(t, dstSlices, res)
	})

	t.Run("int - string", func(t *testing.T) {
		srcSlice := []int{1, 2, 3}
		dstSlice := []string{"1", "2", "3"}
		converter := func(in int) string {
			return fmt.Sprint(in)
		}

		assert.Equal(t, dstSlice, ConvertSlice(srcSlice, converter))
	})

	t.Run("string - string", func(t *testing.T) {
		srcSlice := []string{"a  ", "  b  ", "  c  "}
		dstSlice := []string{"a", "b", "c"}
		assert.Equal(t, dstSlice, ConvertSlice(srcSlice, strings.TrimSpace))
	})
}

func Test_ReverseSlice(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		source := []int{1, 2, 3, 4, 5}
		dest := []int{5, 4, 3, 2, 1}
		assert.Equal(t, dest, ReverseSlice(source))
	})

	t.Run("empty slice", func(t *testing.T) {
		source := []int{}
		dest := []int{}
		assert.Equal(t, dest, ReverseSlice(source))
	})
}
