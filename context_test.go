package utils

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"testing"
)

func Test_GetContextValueByKey(t *testing.T) {
	key := "user_id"
	t.Run("found value", func(t *testing.T) {
		md := metadata.New(map[string]string{key: "12345"})
		ctx := metadata.NewIncomingContext(context.TODO(), md)

		v := GetContextValueByKey(ctx, key)
		assert.Equal(t, "12345", v)
	})

	t.Run("not found value", func(t *testing.T) {
		md := metadata.New(map[string]string{})
		ctx := metadata.NewIncomingContext(context.TODO(), md)

		v := GetContextValueByKey(ctx, key)
		assert.Equal(t, "", v)
	})
}