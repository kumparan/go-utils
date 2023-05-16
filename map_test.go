package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LowerMapStringKey(t *testing.T) {
	value := map[string]interface{}{
		"storyID": "testing",
	}
	res := LowerMapStringKey(value)
	assert.Equal(t, res["storyid"].(string), "testing")
}

func Test_MapValuesToOrderedSlice(t *testing.T) {
	value := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
		"d": "D",
		"e": "E",
	}

	type tc struct {
		order  []string
		output []string
	}
	testCases := []tc{
		{
			order:  []string{"a", "b", "c"},
			output: []string{"A", "B", "C"},
		},
		{
			order:  []string{"d", "e"},
			output: []string{"D", "E"},
		},
		{
			order:  []string{"f", "g", "h"},
			output: []string{"", "", ""},
		},
		{
			order:  []string{"c", "a", "d"},
			output: []string{"C", "A", "D"},
		},
		{
			order:  []string{"d", "f", "a"},
			output: []string{"D", "", "A"},
		},
	}

	for _, tc := range testCases {
		res := MapValuesToOrderedSlice(value, tc.order)
		assert.EqualValues(t, tc.output, res)
	}
}

func Test_MapValuesToOrderedSliceExcludeNil(t *testing.T) {
	value := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
		"d": "D",
		"e": "E",
	}

	type tc struct {
		order  []string
		output []string
	}
	testCases := []tc{
		{
			order:  []string{"a", "b", "c"},
			output: []string{"A", "B", "C"},
		},
		{
			order:  []string{"d", "e"},
			output: []string{"D", "E"},
		},
		{
			order:  []string{"f", "g", "h"},
			output: nil,
		},
		{
			order:  []string{"c", "a", "d"},
			output: []string{"C", "A", "D"},
		},
		{
			order:  []string{"d", "f", "a"},
			output: []string{"D", "A"},
		},
	}

	for _, tc := range testCases {
		res := MapValuesToOrderedSliceExcludeNil(value, tc.order)
		assert.EqualValues(t, tc.output, res)
	}
}
