package utils

import "strings"

// LowerMapStringKey :nodoc:
func LowerMapStringKey(v map[string]interface{}) map[string]interface{} {
	lv := make(map[string]interface{}, len(v))
	for mk, mv := range v {
		lv[strings.ToLower(mk)] = mv
	}
	return lv
}
