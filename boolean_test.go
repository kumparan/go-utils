package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsEmailValid(t *testing.T) {
	assert.Equal(t, false, IsEmailValid("bebek123"))
	assert.Equal(t, false, IsEmailValid("bebek123/gmal"))
	assert.Equal(t, true, IsEmailValid("bebek123@gmal.com"))
}

func Test_IsNumeric(t *testing.T) {
	assert.Equal(t, false, IsNumeric("bebek123"))
	assert.Equal(t, true, IsNumeric("123"))
}

func Test_BoolToString(t *testing.T) {
	assert.Equal(t, strconv.FormatBool(false), BoolToString(false))
	assert.Equal(t, strconv.FormatBool(true), BoolToString(true))
}

func Test_BoolPointerToBool(t *testing.T) {
	var b *bool
	assert.Equal(t, false, BoolPointerToBool(b))
	bb := false
	b = &bb
	assert.Equal(t, bb, BoolPointerToBool(b))
	*b = true
	assert.Equal(t, true, BoolPointerToBool(b))
}
