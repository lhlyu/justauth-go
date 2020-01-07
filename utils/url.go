package utils

import (
	"fmt"
	"net/url"
)

type UrlBuilder struct {
	baseUrl string
	params  map[string]interface{}
}

func NewUrlBuilder(baseUrl string) *UrlBuilder {
	return &UrlBuilder{
		baseUrl: baseUrl,
		params:  make(map[string]interface{}),
	}
}

func (this *UrlBuilder) QueryParam(key string, value interface{}) *UrlBuilder {
	if key == "" {
		return this
	}
	this.params[key] = value
	return this
}

func (this *UrlBuilder) Build() string {
	if len(this.params) == 0 {
		return this.baseUrl
	}
	uv := url.Values{}
	for k, v := range this.params {
		uv.Add(k, fmt.Sprint(v))
	}
	return this.baseUrl + "?" + uv.Encode()
}
