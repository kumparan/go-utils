package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseDuration(t *testing.T) {
	type tableTest struct {
		Input         string
		Default       time.Duration
		ExpectDefault bool
		Expect        time.Duration
	}

	tts := []tableTest{
		{
			Input:  "500m50ms",
			Expect: (500 * time.Minute) + (50 * time.Millisecond),
		},
		{
			Input:         "",
			ExpectDefault: true,
			Default:       5 * time.Minute,
			Expect:        5 * time.Minute,
		},
		{
			Input:         "100",
			ExpectDefault: true,
			Default:       3 * time.Minute,
			Expect:        3 * time.Minute,
		},
		{
			Input:         "100x",
			ExpectDefault: true,
			Default:       1 * time.Minute,
			Expect:        1 * time.Minute,
		},
	}

	for _, test := range tts {
		dur := ParseDurationWithDefault(test.Input, test.Default)
		if test.ExpectDefault {
			require.EqualValues(t, test.Default, dur)
			continue
		}

		require.EqualValues(t, test.Expect, dur)
	}
}
