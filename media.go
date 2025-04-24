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

	case LargeIcon:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_256,ar_1:1,f_jpeg"

	case BigPicture:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_1440,h_720,f_jpeg"

	case IOSAttachment:
		param = "image/upload/fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_1024,h_1024,f_jpeg"
	}

	mediaURL := cdnURL + "/" + param + "/" + coverMediaFile

	return mediaURL
}
