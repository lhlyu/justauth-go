package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Get(apiUrl string, params map[string]string, headers map[string]string) (string, error) {
	u := &url.Values{}
	for k, v := range params {
		u.Add(k, v)
	}
	if len(u.Encode()) > 0 {
		apiUrl += "?" + u.Encode()
	}
	req, _ := http.NewRequest("GET", apiUrl, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	byts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byts), nil
}

func Post(apiUrl string, params map[string]string, headers map[string]string, data interface{}) (string, error) {
	byts, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	u := &url.Values{}
	for k, v := range params {
		u.Add(k, v)
	}
	if len(u.Encode()) > 0 {
		apiUrl += "?" + u.Encode()
	}
	req, _ := http.NewRequest("POST", apiUrl, strings.NewReader(string(byts)))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	byts, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byts), nil
}
