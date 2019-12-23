package utils

import (
	"encoding/json"
	"strconv"
	"strings"
)

// StandardizeSpaces -> JoinURL long query to one line query
func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// UnescapeString UTF-8 string
// e.g. convert "\u0e27\u0e23\u0e0d\u0e32" to "วรญา"
func UnescapeString(str string) (ustr string) {
	_ = json.Unmarshal([]byte(`"`+str+`"`), &ustr)
	return
}

// StringToBool :nodoc:
func StringToBool(s string) bool {
	if s != "" {
		i, err := strconv.ParseBool(s)
		if err == nil {
			return i
		}
	}
	return false
}

// StringToInt64 :nodoc:
func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// StringToInt64WithDefault :nodoc:
func StringToInt64WithDefault(s string, d int64) int64 {
	i := StringToInt64(s)
	if i == 0 {
		return d
	}
	return i
}

// StringPointerToString :nodoc:
func StringPointerToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// StringPointerToFloat64 :nodoc:
func StringPointerToFloat64(s *string) float64 {
	if s != nil {
		f, err := strconv.ParseFloat(*s, 64)
		if err != nil {
			return float64(0)
		}
		return f
	}
	return float64(0)
}

// StringPointerToInt64 :nodoc:
func StringPointerToInt64(s *string) int64 {
	if s == nil {
		return int64(0)
	}

	return StringToInt64(*s)
}

// ArrayStringPointerToArrayInt64 :nodoc:
func ArrayStringPointerToArrayInt64(s *[]*string) []int64 {
	var i []int64
	if s != nil {
		for _, val := range *s {
			i = append(i, StringPointerToInt64(val))
		}
		return i
	}
	return nil
}
