package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func StrToMSS(s string) map[string]string {
	if s == "" {
		return nil
	}
	fields := strings.Split(s, "&")
	mss := make(map[string]string)
	for _, field := range fields {
		if strings.Contains(field, "=") {
			keyValue := strings.Split(field, "=")
			key, _ := url.PathUnescape(keyValue[0])
			value := ""
			if len(keyValue) == 2 {
				value, _ = url.PathUnescape(keyValue[1])
				value = strings.ReplaceAll(value, "+", " ")
			}
			mss[key] = value
		}
	}
	return mss
}

func JsonToMSS(s string) map[string]string {
	if s == "" {
		return nil
	}
	msi := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &msi)
	if err != nil {
		return nil
	}
	mss := make(map[string]string)
	for k, v := range msi {
		mss[k] = fmt.Sprint(v)
	}
	return mss
}
