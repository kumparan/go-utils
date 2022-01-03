package xgqlgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
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

func TestGormDeletedAt(t *testing.T) {
	now := time.Now()
	deletedAt := gorm.DeletedAt{}
	err := deletedAt.Scan(now)
	require.NoError(t, err)

	t.Run("marshal", func(t *testing.T) {
		ts := fmt.Sprintf(`"%s"`, now.Format(time.RFC3339Nano))
		assert.Equal(t, ts, marshalerToString(MarshalGormDeletedAt(deletedAt)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		ts := now.Format(time.RFC3339Nano)
		gd, err := UnmarshalGormDeletedAt(ts)
		assert.NoError(t, err)
		assert.EqualValues(t, now.Format(time.RFC3339Nano), gd.Time.Format(time.RFC3339Nano))
	})
}

func TestTime(t *testing.T) {
	now := time.Now()

	t.Run("marshal", func(t *testing.T) {
		ts := fmt.Sprintf(`"%s"`, now.Format(time.RFC3339Nano))
		assert.Equal(t, ts, marshalerToString(MarshalTimeRFC3339Nano(now)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		ts := now.Format(time.RFC3339Nano)
		tt, err := UnmarshalTimeRFC3339Nano(ts)
		assert.NoError(t, err)
		assert.EqualValues(t, now.Format(time.RFC3339Nano), tt.Format(time.RFC3339Nano))
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
