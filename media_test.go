package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMediaURL(t *testing.T) {
	externalMediaURL := "https://res.cloudinary.com/kumpar/image/upload/v1481882219/qaejp6n1amazr7ud2aaf.jpg"
	res := GenerateMediaURL("http://blue.kumparan.com", externalMediaURL, _LargeIcon)
	assert.Equal(t, "http://blue.kumparan.com/image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_256,ar_1:1/v1481882219/qaejp6n1amazr7ud2aaf.jpg", res)
}
