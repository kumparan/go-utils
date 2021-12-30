package xgqlgen

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kumparan/go-utils"
)

// MarshalInt64ID marshal int64 to string ID
func MarshalInt64ID(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(fmt.Sprintf(`"%d"`, i)))
	})
}

// UnmarshalInt64ID unmarshal ID into int64
func UnmarshalInt64ID(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return utils.StringToInt64(v), nil
	case json.Number:
		return utils.StringToInt64(string(v)), nil
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("%T is not a int64", v)
	}
}
