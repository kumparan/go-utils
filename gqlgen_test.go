package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
)

func TestInt64(t *testing.T) {
	id := int64(1640848405303336961)
	idStr := fmt.Sprint(id)
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, `"1640848405303336961"`, marshalerToString(MarshalInt64ID(id)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, id, mustUnmarshalInt64ID(t, id))
		assert.Equal(t, id, mustUnmarshalInt64ID(t, json.Number(idStr)))
		assert.Equal(t, id, mustUnmarshalInt64ID(t, idStr))
	})
}

func mustUnmarshalInt64ID(t *testing.T, v interface{}) int64 {
	res, err := UnmarshalInt64ID(v)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func marshalerToString(m graphql.Marshaler) string {
	var b bytes.Buffer
	m.MarshalGQL(&b)
	return b.String()
}
