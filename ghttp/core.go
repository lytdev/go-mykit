package ghttp

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"sync"
)

// https://github.com/go-resty/resty

var mutex sync.Mutex

// GetRestyClient 获取resty执行请求的客户端
func GetRestyClient() *resty.Request {
	mutex.Lock()         //加锁
	defer mutex.Unlock() //解锁
	return resty.New().R()
}

// GetRespJsonData 获取响应的数据实体
func GetRespJsonData[T any](bodyStr string) (error, T) {
	var respData T
	if err := json.Unmarshal([]byte(bodyStr), &respData); err == nil {
		return nil, respData
	}
	return nil, respData
}

// GetJson GET请求获取json数据
func GetJson[T any](url string, queryParams map[string]string) (error, T) {
	var tmp T
	client := GetRestyClient()
	resp, reqErr := client.
		SetQueryParams(queryParams).
		SetHeader("Accept", "application/json").
		Get(url)
	if reqErr != nil {
		return reqErr, tmp
	}
	bodyStr := resp.String()
	return GetRespJsonData[T](bodyStr)
}

// Get GET请求获取数据
func Get(url string, queryParams map[string]string, headers map[string]string) (error, string) {
	client := GetRestyClient()
	resp, reqErr := client.
		SetQueryParams(queryParams).
		SetHeaders(headers).
		Get(url)
	if reqErr != nil {
		return reqErr, ""
	}
	return nil, resp.String()
}
