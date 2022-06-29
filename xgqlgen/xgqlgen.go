package xgqlgen

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kumparan/go-utils"
	"gopkg.in/guregu/null.v4"
)

// MarshalInt64ID marshal int64 to string ID
func MarshalInt64ID(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(fmt.Sprintf(`"%d"`, i)))
	})
}

// UnmarshalInt64ID unmarshal ID into int64
func UnmarshalInt64ID(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case json.Number:
		return strconv.ParseInt(string(v), 10, 64)
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("%T is not a number", v)
	}
}

// MarshalNullInt64ID marshal int64 to string ID
func MarshalNullInt64ID(i null.Int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(fmt.Sprintf(`"%d"`, i.Int64)))
	})
}

// UnmarshalNullInt64ID unmarshal ID into int64
func UnmarshalNullInt64ID(v interface{}) (null.Int, error) {
	switch v := v.(type) {
	case string:
		valInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return null.Int{}, err
		}
		return null.IntFrom(valInt), nil
	case json.Number:
		valInt, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return null.Int{}, err
		}
		return null.IntFrom(valInt), nil
	case int:
		return null.IntFrom(int64(v)), nil
	case int64:
		return null.IntFrom(v), nil
	case int32:
		return null.IntFrom(int64(v)), nil
	case nil:
		return null.Int{}, nil
	default:
		return null.Int{}, fmt.Errorf("%T is not a number", v)
	}
}

// MarshalNullInt64 marshal int64 to string ID
func MarshalNullInt64(i null.Int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(fmt.Sprintf(`%d`, i.Int64)))
	})
}

// UnmarshalNullInt64 unmarshal ID into int64
func UnmarshalNullInt64(v interface{}) (null.Int, error) {
	switch v := v.(type) {
	case string:
		valInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return null.Int{}, err
		}
		return null.IntFrom(valInt), nil
	case json.Number:
		valInt, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return null.Int{}, err
		}
		return null.IntFrom(valInt), nil
	case int:
		return null.IntFrom(int64(v)), nil
	case int64:
		return null.IntFrom(v), nil
	case int32:
		return null.IntFrom(int64(v)), nil
	case nil:
		return null.Int{}, nil
	default:
		return null.Int{}, fmt.Errorf("%T is not a number", v)
	}
}

// MarshalTimeRFC3339Nano marshal time.Time to string RFC3339Nano
func MarshalTimeRFC3339Nano(tt time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		ts := utils.FormatTimeRFC3339(&tt)
		_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, ts)))
	})
}

// UnmarshalTimeRFC3339Nano unmarshal v into time.Time
func UnmarshalTimeRFC3339Nano(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		tt, err := time.Parse(time.RFC3339Nano, v)
		return tt, err
	default:
		return time.Time{}, fmt.Errorf("%T is not a valid RFC3339Nano time", v)
	}
}

// MarshalGormDeletedAt marshal gorm.DeletedAt to string
func MarshalGormDeletedAt(gd gorm.DeletedAt) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if !gd.Valid {
			_, _ = w.Write([]byte(`null`))
			return
		}
		ts := utils.FormatTimeRFC3339(&gd.Time)
		_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, ts)))
	})
}

// UnmarshalGormDeletedAt unmarshal v into gorm.DeletedAt
func UnmarshalGormDeletedAt(v interface{}) (gorm.DeletedAt, error) {
	switch v := v.(type) {
	case string:
		tt, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return gorm.DeletedAt{}, err
		}
		gd := gorm.DeletedAt{}
		err = gd.Scan(tt)
		return gd, err
	default:
		return gorm.DeletedAt{}, errors.New("v is not a valid string time")
	}
}

// MarshalNullTimeRFC3339Nano :nodoc:
func MarshalNullTimeRFC3339Nano(nt null.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if !nt.Valid {
			_, _ = w.Write([]byte(`null`))
			return
		}
		ts := utils.FormatTimeRFC3339(&nt.Time)
		_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, ts)))
	})
}

// UnmarshalNullTimeRFC3339Nano unmarshal v into time.Time
func UnmarshalNullTimeRFC3339Nano(v interface{}) (null.Time, error) {
	switch v := v.(type) {
	case string:
		tt, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return null.Time{}, err
		}
		return null.NewTime(tt, true), nil
	case nil:
		return null.Time{}, nil
	default:
		return null.Time{}, fmt.Errorf("%T is not a valid RFC3339Nano time", v)
	}
}

// UnmarshalNullString unmarshal string into null.String
func UnmarshalNullString(v interface{}) (null.String, error) {
	switch v := v.(type) {
	case string:
		return null.StringFrom(v), nil
	case nil:
		return null.String{}, nil
	default:
		return null.String{}, fmt.Errorf("%T is not a string", v)
	}
}

// MarshalNullString marshal int64 to string ID
func MarshalNullString(i null.String) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, i.String)))
	})
}

// ConstraintSize directive to constrain field between min and max values. if field is above max, then directive returns max. if field is below min, then directive returns min. else return field.
func ConstraintSize(ctx context.Context, obj interface{}, next graphql.Resolver, min int64, max int64, field *string) (interface{}, error) {
	val, ok := obj.(map[string]interface{}) // safe check is valid map
	if !ok {
		return next(ctx) // skip if invalid
	}

	fieldName := "size"
	if field != nil {
		fieldName = *field
	}

	valInt, ok := val[fieldName].(int64) // safe check is valid int64
	if !ok {
		return next(ctx) // skip if invalid
	}

	return utils.Int64WithMinAndMaxLimit(valInt, min, max), nil
}
