package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDifferenceTime_GetDifferenceDaysForHumans(t *testing.T) {

	now := time.Now()
	testCases := []struct {
		input    time.Time
		expected string // 期望值
	}{
		{now.Add(72 * time.Hour), "3 hari lagi"},
		{now.Add(48 * time.Hour), "2 hari lagi"},
		{now.Add(24 * time.Hour), "besok"},
		{now.Add(-72 * time.Hour), "3 hari yang lalu"},
		{now.Add(-48 * time.Hour), "2 hari yang lalu"},
		{now.Add(-24 * time.Hour), "kemarin"},
		{now.Add(8784 * time.Hour), "366 hari lagi"},
		{now.Add(-8784 * time.Hour), "366 hari yang lalu"},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, GetDifferenceDaysForHumans(time.Now(), testCase.input))
	}
}
