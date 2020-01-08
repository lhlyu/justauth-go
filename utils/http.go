package utils

import (
	"github.com/lhlyu/justauth-go/enums"
	"github.com/lhlyu/justauth-go/errcode"
	"io/ioutil"
	"net/http"
)

func Post(url string) (string, *errcode.ErrCode) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", errcode.NewErrCode(enums.FAILURE).WithMsg(err.Error())
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errcode.NewErrCode(enums.FAILURE).WithMsg(err.Error())
	}
	return string(body), nil
}

func Get(url string) (string, *errcode.ErrCode) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", errcode.NewErrCode(enums.FAILURE).WithMsg(err.Error())
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errcode.NewErrCode(enums.FAILURE).WithMsg(err.Error())
	}
	return string(body), nil
}
