package utils

import "strings"

const (
	_BigPicture    = "big_picture"
	_LargeIcon     = "large_icon"
	_IOSAttachment = "ios_attahcment"
)

// GenerateMediaURL :nodoc:
func GenerateMediaURL(cdnURL, externalURL, imageType string) string {
	if externalURL == "" {
		return ""
	}
	splittedMediaExternalURL := strings.Split(externalURL, "/")
	coverMediaFile := splittedMediaExternalURL[len(splittedMediaExternalURL)-2] + "/" + splittedMediaExternalURL[len(splittedMediaExternalURL)-1]
	var param string
	switch imageType {

	case _LargeIcon:
		param = "fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_256,ar_1:1"

	case _BigPicture:
		param = "fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_1440,h_720"

	case _IOSAttachment:
		param = "fl_progressive,fl_lossy,c_fill,g_face,q_auto:best,w_1024,h_1024"
	}

	mediaURL := cdnURL + "/image/upload/" + param + "/" + coverMediaFile

	return mediaURL
}
