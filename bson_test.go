package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTimeFromObjectIDHex(b *testing.B) {
	cases := []string{
		"63acfc824ffda9000ee65045",
		"5c4b288bac1b2972d0291377",
	}

	for _, v := range cases {
		_, _ = TimeFromObjectIDHex(v)

	}
}

func TestTimeFromObjectIDHex(t *testing.T) {
	t.Run("valid hex objectID", func(t *testing.T) {
		cases := map[string]time.Time{
			"63acfc824ffda9000ee65045": time.Date(2022, time.December, 29, 2, 33, 38, 0, time.Local),
			"5c4b288bac1b2972d0291377": time.Date(2019, time.January, 25, 15, 17, 31, 0, time.Local),
		}

		for k, v := range cases {
			r, err := TimeFromObjectIDHex(k)
			require.NoError(t, err)
			assert.Equal(t, v.UTC(), r.UTC())
		}
	})

	t.Run("invalid hex objectID", func(t *testing.T) {
		cases := []string{
			"anuanu",
			"876124892",
		}

		for _, c := range cases {
			_, err := TimeFromObjectIDHex(c)
			require.Error(t, err)
		}
	})
}
