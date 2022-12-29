package utils

import (
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

// TimeFromObjectIDHex extracts time data from bson's objectID
func TimeFromObjectIDHex(s string) (t time.Time, err error) {
	if !bson.IsObjectIdHex(s) {
		err = errors.New("invalid hex objectID")
		return
	}

	return bson.ObjectIdHex(s).Time(), nil
}
