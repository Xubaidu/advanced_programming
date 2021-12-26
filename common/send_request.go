package common

import (
	. "advanced_programming/schema"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func HttpMethod(httpMethodReq *HttpMethodRequest) (httpResp *http.Response, err error) {
	// 声明一个 writer
	buf := new(bytes.Buffer)

	// 把 httpMethodReq 解码到 buf
	err = json.NewEncoder(buf).Encode(httpMethodReq.Params)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	// 声明一个 request
	req, err := http.NewRequest(httpMethodReq.Method, httpMethodReq.URL, buf)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	// 设置 request 的请求头
	for k, v := range httpMethodReq.Header {
		req.Header.Set(k, v)
	}

	// 根据 req 的设置发 https 请求
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

func SendHttpRequest(httpReq *HttpRequest, respParams interface{}) (err error) {
	// 构造请求
	httpMethodReq := &HttpMethodRequest{
		Method: httpReq.Method,
		URL:    httpReq.URL,
		Header: make(map[string]string),
		Params: httpReq.ReqParams,
	}

	// 填充 header
	if httpReq.ContentType != "" {
		httpMethodReq.Header["Content-Type"] = httpReq.ContentType
	}
	if httpReq.Token != "" {
		httpMethodReq.Header["Authorization"] = "Bearer " + httpReq.Token
	}

	// 获取响应
	resp, err := HttpMethod(httpMethodReq)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	// 解码到 respParams
	err = json.NewDecoder(resp.Body).Decode(respParams)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	return nil

}
