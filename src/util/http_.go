package util

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// http get 返回json
func HttpGetJson[respClass any](url string) (*respClass, error) {
	logMaker := LogMaker{}
	defer logMaker.LogString()
	logMaker.AddLine("http GET ")
	logMaker.AddKV(" url", url)

	resp, err := http.Get(url)
	if err != nil {
		logMaker.AddKV(" http.Get error", err)
		return nil, err
	}

	bytex, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		logMaker.AddKV(" io.ReadAll error", err2)
		return nil, err2
	}

	logMaker.AddKV(" result", string(bytex))

	// json
	var obj respClass
	err3 := json.Unmarshal(bytex, &obj)
	if err3 != nil {
		logMaker.AddKV(" json.Unmarshal error", err3)
		return nil, err3
	}

	return &obj, nil
}

// http get 返回json
func HttpGetJson2[respClass any](url1 string, params map[string]string) (*respClass, error) {
	// 解析。
	url2, err1 := url.Parse(url1)
	CheckError(err1)
	// 填参数。
	query := url2.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	// 编码
	url2.RawQuery = query.Encode()
	url3 := url2.String()
	return HttpGetJson[respClass](url3)
}

// http get 入参json，返回json
func HttpPostJson[reqClass any, respClass any](url1 string, req *reqClass) (*respClass, error) {
	logMaker := LogMaker{}
	defer logMaker.LogString()
	logMaker.AddLine("http POST ")
	logMaker.AddKV(" url", url1)

	reqJson := ToJson(req)
	logMaker.AddKV(" reqJson", reqJson)

	request, err1 := http.NewRequest("POST", url1, strings.NewReader(reqJson))
	if err1 != nil {
		logMaker.AddKV(" http.NewRequest error", err1)
		return nil, err1
	}

	resp, err2 := http.DefaultClient.Do(request)
	if err2 != nil {
		return nil, err2
	}

	bytex, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		logMaker.AddKV(" io.ReadAll error", err2)
		return nil, err2
	}

	logMaker.AddKV(" result", string(bytex))

	// json
	var obj respClass
	err3 := json.Unmarshal(bytex, &obj)
	if err3 != nil {
		logMaker.AddKV(" json.Unmarshal error", err3)
		return nil, err3
	}

	return &obj, nil
}
