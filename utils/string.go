package utils

import (
	"fmt"
	"strconv"
)

func ToString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if v == nil {
			return ""
		}
		switch v.(type) {
		case int:
			return strconv.FormatInt(int64(v.(int)), 10)
		case float64:
			return strconv.FormatFloat(v.(float64), 'f', -1, 64)
		case string:
			return v.(string)
		case nil:
			return ""
		default:
			return fmt.Sprint(v)
		}
	}
	return ""
}
