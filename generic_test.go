package utils

import (
	"encoding/json"
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
