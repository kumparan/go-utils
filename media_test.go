package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GeneratePushNotificationMediaURL(t *testing.T) {
	externalMediaURL := "https://this.is/test/image/upload/123/234.jpg"
	res := GeneratePushNotificationMediaURL("http://blue.kumparan.com", externalMediaURL, _LargeIcon)
	assert.Equal(t, "http://blue.kumparan.com/image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_256,ar_1:1/123/234.jpg", res)
}
