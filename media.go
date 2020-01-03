package utils

import (
	"strings"
)

const (
	// BigPicture :nodoc:
	BigPicture = "big_picture"
	// LargeIcon :nodoc:
	LargeIcon = "large_icon"
	// IOSAttachment :nodoc:
	IOSAttachment = "ios_attahcment"
)

// GeneratePushNotificationMediaURL :nodoc:
func GeneratePushNotificationMediaURL(baseImageURL, srcImageURL, imageType string) string {
	if srcImageURL == "" || baseImageURL == "" {
		return ""
	}

	splitMediasrcImageURL := strings.Split(srcImageURL, "/")
	coverMediaFile := splitMediasrcImageURL[len(splitMediasrcImageURL)-2] + "/" + splitMediasrcImageURL[len(splitMediasrcImageURL)-1]

	var param string
	switch imageType {

	case LargeIcon:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_256,ar_1:1"

	case BigPicture:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_1440,h_720"

	case IOSAttachment:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_1024,h_1024"
	}

	mediaURL := baseImageURL + "/" + param + "/" + coverMediaFile

	return mediaURL
}
