package xgqlgen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kumparan/go-utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
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

func TestNullInt64ID(t *testing.T) {
	id := int64(1640848405303336961)
	idStr := fmt.Sprint(id)
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, `"1640848405303336961"`, marshalerToString(MarshalNullInt64ID(null.NewInt(id, true))))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, id, mustUnmarshalNullInt64ID(t, id).Int64)
		assert.Equal(t, id, mustUnmarshalNullInt64ID(t, json.Number(idStr)).Int64)
		assert.Equal(t, id, mustUnmarshalNullInt64ID(t, idStr).Int64)
	})
}

func TestNullInt64(t *testing.T) {
	val := int64(1640848405303336961)
	idStr := fmt.Sprint(val)
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, val, marshalerToInt64(MarshalNullInt64(null.NewInt(val, true))))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, val, mustUnmarshalInt64(t, val).Int64)
		assert.Equal(t, val, mustUnmarshalInt64(t, json.Number(idStr)).Int64)
		assert.Equal(t, val, mustUnmarshalInt64(t, idStr).Int64)
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

func TestNullTime(t *testing.T) {
	now := time.Now()
	nullNow := null.NewTime(now, true)
	t.Run("marshal", func(t *testing.T) {
		ts := fmt.Sprintf(`"%s"`, now.Format(time.RFC3339Nano))
		assert.Equal(t, ts, marshalerToString(MarshalNullTimeRFC3339Nano(nullNow)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		ts := now.Format(time.RFC3339Nano)
		tt, err := UnmarshalNullTimeRFC3339Nano(ts)
		assert.NoError(t, err)
		assert.EqualValues(t, now.Format(time.RFC3339Nano), tt.Time.Format(time.RFC3339Nano))
	})
}

func TestConstraintSize(t *testing.T) {
	var (
		min        int64           = 1
		max        int64           = 25
		a          int64           = 5
		b          int64           = 0
		c          int64           = 26
		sizeField  string          = "size"
		otherField string          = "other"
		ctx        context.Context = context.TODO()
	)

	t.Run("within constraint", func(t *testing.T) {
		obj := map[string]interface{}{sizeField: a}
		res, err := ConstraintSize(ctx, obj, nil, min, max, nil)
		assert.NoError(t, err)
		assert.EqualValues(t, a, res)
	})

	t.Run("below min constraint", func(t *testing.T) {
		obj := map[string]interface{}{sizeField: b}
		res, err := ConstraintSize(ctx, obj, nil, min, max, nil)
		assert.NoError(t, err)
		assert.EqualValues(t, min, res)
	})

	t.Run("above max constraint", func(t *testing.T) {
		obj := map[string]interface{}{sizeField: c}
		res, err := ConstraintSize(ctx, obj, nil, min, max, nil)
		assert.NoError(t, err)
		assert.EqualValues(t, max, res)
	})

	t.Run("non-default field", func(t *testing.T) {
		obj := map[string]interface{}{otherField: c}
		res, err := ConstraintSize(ctx, obj, nil, min, max, &otherField)
		assert.NoError(t, err)
		assert.EqualValues(t, max, res)
	})
}

func mustUnmarshalInt64ID(t *testing.T, v interface{}) int64 {
	res, err := UnmarshalInt64ID(v)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func mustUnmarshalInt64(t *testing.T, v interface{}) null.Int {
	res, err := UnmarshalNullInt64(v)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func mustUnmarshalNullInt64ID(t *testing.T, v interface{}) null.Int {
	res, err := UnmarshalNullInt64ID(v)
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

func marshalerToInt64(m graphql.Marshaler) int64 {
	var b bytes.Buffer
	m.MarshalGQL(&b)
	return utils.StringToInt64(b.String())
}
