package utils

import (
	"strings"
)

const (
	// _BigPicture :nodoc:
	_BigPicture = "big_picture"
	// _LargeIcon :nodoc:
	_LargeIcon = "large_icon"
	// _IOSAttachment :nodoc:
	_IOSAttachment = "ios_attahcment"
)

// GeneratePushNotificationMediaURL Generates manipulated media URL for push notification purpose
// e.g. GeneratePushNotificationMediaURL("http://mycdn.com", "http://my.image.com/image/upload/v123/image.jpg", LargeIcon) => http://mycdn.com/image/upload/v123/image.jpg
func GeneratePushNotificationMediaURL(cdnURL, mediaSrcURL, imageType string) string {
	if mediaSrcURL == "" || cdnURL == "" {
		return ""
	}

	splitMediaSrcImageURL := strings.Split(mediaSrcURL, "/")
	coverMediaFile := splitMediaSrcImageURL[len(splitMediaSrcImageURL)-2] + "/" + splitMediaSrcImageURL[len(splitMediaSrcImageURL)-1]

	var param string
	switch imageType {

	case _LargeIcon:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_256,ar_1:1"

	case _BigPicture:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_1440,h_720"

	case _IOSAttachment:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_1024,h_1024"
	}

	mediaURL := cdnURL + "/" + param + "/" + coverMediaFile

	return mediaURL
}
