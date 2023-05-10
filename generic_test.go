package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	int64s := []int64{1, 2, 3, 4, 5}
	int32s := []int32{1, 2, 3, 4, 5}
	strings := []string{"a", "be", "see", "deep"}
	float64s := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	float32s := []float32{1.2, 2.3, 3.4, 4.5, 5.6}

	assert.True(t, Contains[int64](int64s, int64(1)))
	assert.False(t, Contains[int64](int64s, int64(0)))

	assert.True(t, Contains[int32](int32s, int32(1)))
	assert.False(t, Contains[int32](int32s, int32(0)))

	assert.True(t, Contains[string](strings, "be"))
	assert.False(t, Contains[string](strings, "zebra"))

	assert.True(t, Contains[float64](float64s, 1.1))
	assert.False(t, Contains[float64](float64s, 0.1))

	assert.True(t, Contains[float32](float32s, float32(1.2)))
	assert.False(t, Contains[float32](float32s, float32(0.9)))
}

func Test_Difference(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"c", "d", "e"}

	assert.Equal(t, []string{"a", "b", "d", "e"}, Difference[string](s1, s2))
	assert.Equal(t, []string{"d", "e", "a", "b"}, Difference[string](s2, s1))

	i1 := []int64{1, 2, 3}
	i2 := []int64{3, 4, 5}

	assert.Equal(t, []int64{1, 2, 4, 5}, Difference[int64](i1, i2))
	assert.Equal(t, []int64{4, 5, 1, 2}, Difference[int64](i2, i1))
}

func Test_Unique(t *testing.T) {
	s := []string{"a", "a", "b", "a", "d", "b"}
	assert.Equal(t, []string{"a", "b", "d"}, Unique[string](s))

	i := []int64{1, 1, 2, 4, 2}
	assert.Equal(t, []int64{1, 2, 4}, Unique[int64](i))
}

func Test_InterfaceBytesToType(t *testing.T) {
	someInteger := 8
	bt, _ := json.Marshal(someInteger)
	resultInt64 := InterfaceBytesToType[int64](interface{}(bt))
	assert.EqualValues(t, someInteger, resultInt64)

	someString := "this-is-string"
	bt, _ = json.Marshal(someString)
	resultString := InterfaceBytesToType[string](interface{}(bt))
	assert.EqualValues(t, someString, resultString)

	type myStruct struct {
		Name string
		Age  int
	}
	someStruct := myStruct{Name: someString, Age: someInteger}
	bt, _ = json.Marshal(someStruct)
	resultStruct := InterfaceBytesToType[myStruct](interface{}(bt))
	assert.EqualValues(t, someStruct, resultStruct)
}

func Test_ValueOrDefault(t *testing.T) {
	assert.EqualValues(t, 10, ValueOrDefault[int64](0, 10))
	assert.EqualValues(t, -5, ValueOrDefault[int64](-5, 10))
	assert.EqualValues(t, 10, ValueOrDefault[int64](10, 100))

	assert.EqualValues(t, "hello", ValueOrDefault[string]("", "hello"))
	assert.EqualValues(t, "hello", ValueOrDefault[string]("hello", "hi"))

	assert.EqualValues(t, 10.1, ValueOrDefault[float64](0.0, 10.1))
	assert.EqualValues(t, 11.2, ValueOrDefault[float64](11.2, 10.1))
}

func Test_DeleteByValue(t *testing.T) {
	int64s := []int64{1, 2, 3, 4, 5}
	int64ExpectTrueDelete := []int64{2, 3, 4, 5}

	int64sUnique := []int64{1, 2, 3, 4, 5, 1}
	int64sUniqueExpectTrueDelete := []int64{2, 3, 4, 5}

	int32s := []int32{1, 2, 3, 4, 5}
	int32ExpectTrueDelete := []int32{1, 2, 4, 5}

	strings := []string{"a", "be", "see", "deep"}
	stringExpectTrueDelete := []string{"a", "be", "deep"}

	float64s := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	float64ExpectTrueDelete := []float64{1.1, 3.3, 4.4, 5.5}

	float32s := []float32{1.2, 2.3, 3.4, 4.5, 5.6}
	float32ExpectTrueDelete := []float32{1.2, 2.3, 3.4, 4.5}

	assert.EqualValues(t, int64ExpectTrueDelete, DeleteByValue[int64](int64s, 1))
	assert.EqualValues(t, int64s, DeleteByValue[int64](int64s, 10))

	assert.EqualValues(t, int64sUniqueExpectTrueDelete, DeleteByValue[int64](int64sUnique, 1))
	assert.EqualValues(t, int64sUnique, DeleteByValue[int64](int64sUnique, 10))

	assert.EqualValues(t, int32ExpectTrueDelete, DeleteByValue[int32](int32s, 3))
	assert.EqualValues(t, int32s, DeleteByValue[int32](int32s, 10))

	assert.EqualValues(t, stringExpectTrueDelete, DeleteByValue[string](strings, "see"))
	assert.EqualValues(t, strings, DeleteByValue[string](strings, "lol"))

	assert.EqualValues(t, float64ExpectTrueDelete, DeleteByValue[float64](float64s, 2.2))
	assert.EqualValues(t, float64s, DeleteByValue[float64](float64s, 5.7))

	assert.EqualValues(t, float32ExpectTrueDelete, DeleteByValue[float32](float32s, 5.6))
	assert.EqualValues(t, float32s, DeleteByValue[float32](float32s, 5.7))

}

func Test_ParseCacheResultToPointerObject(t *testing.T) {
	type TestObj struct {
		ID int64
	}

	t.Run("ok, primitive types, string", func(t *testing.T) {
		data := "123"
		res, err := ParseCacheResultToPointerObject[string](ToByte(data))
		assert.NoError(t, err)
		assert.Equal(t, *res, data)
	})

	t.Run("ok, primitive types, integer", func(t *testing.T) {
		data := int64(123)
		res, err := ParseCacheResultToPointerObject[int64](ToByte(data))
		assert.NoError(t, err)
		assert.Equal(t, *res, data)
	})

	t.Run("ok, object", func(t *testing.T) {
		data := TestObj{ID: 1234}
		res, err := ParseCacheResultToPointerObject[TestObj](ToByte(data))
		assert.NoError(t, err)
		assert.Equal(t, *res, data)
	})

	t.Run("ok, null cache", func(t *testing.T) {
		res, err := ParseCacheResultToPointerObject[TestObj]([]byte(`null`))
		assert.NoError(t, err)
		assert.Nil(t, res)
		// check null type are equal;
		assert.Equal(t, fmt.Sprintf("%T", &TestObj{}), fmt.Sprintf("%T", res))
	})

	t.Run("error, cache result are not byte ", func(t *testing.T) {
		res, err := ParseCacheResultToPointerObject[TestObj](int64(123456))
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "failed to cast int64 to byte", err.Error())
	})

	t.Run("error, failed on unmarshal ", func(t *testing.T) {
		res, err := ParseCacheResultToPointerObject[TestObj](ToByte([]int64{1, 2, 3, 4, 5, 6}))
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "failed to unmarshal [1,2,3,4,5,6] to *utils.TestObj", err.Error())
	})
}
