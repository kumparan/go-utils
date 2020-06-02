package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// DumpOutGoingContext :nodoc:
func DumpOutGoingContext(c context.Context) string {
	md, _ := metadata.FromOutgoingContext(c)
	return Dump(md)
}

// DumpIncomingContext :nodoc:
func DumpIncomingContext(c context.Context) string {
	md, _ := metadata.FromIncomingContext(c)
	return Dump(md)
}

// GetContextValueByKey get value from context by key. metadata will return string
func GetContextValueByKey(c context.Context, k string) string {
	md, _ := metadata.FromIncomingContext(c)
	s := md.Get(k)
	if len(s) > 0 {
		return s[0]
	}

	return ""
}
