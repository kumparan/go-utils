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
