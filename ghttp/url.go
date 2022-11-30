package ghttp

import (
	"errors"
	"net/url"
)

// RawURLGetParam 通过key获取请求参数的value,如果有多个的话,只获取第一个
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2  key="page" will get "1"
func RawURLGetParam(rawUrl, key string) (string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	m := stUrl.Query()
	if v, ok := m[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// RawURLGetParams 通过key获取请求参数所有value的字符串切片形式
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name and key=page will get [1 2]
func RawURLGetParams(rawUrl, key string) ([]string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	m := stUrl.Query()
	if v, ok := m[key]; ok {
		return v, nil
	}
	return nil, errors.New("no param")
}

// RawURLGetAllParams 获取请求参数的map形式
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2#name will get map[boardID:[520] page:[1 2]]
func RawURLGetAllParams(rawUrl string) (map[string][]string, error) {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	m := stUrl.Query()
	return m, nil
}

// RawURLAddParam 通过key和value重新构造新的url地址
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2;key=page,value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2&page=3
func RawURLAddParam(rawUrl, key, value string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Add(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// RawURLAddParams 通过map重新构造新的url地址
func RawURLAddParams(rawUrl string, params map[string]string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for k, v := range params {
		m.Add(k, v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// RawURLDelParam 根据key删除参数重新构造url地址
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2;key=page
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520
func RawURLDelParam(rawUrl, key string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Del(key)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// RawURLDelParams 根据keys切片删除所有参数重新构造url地址
func RawURLDelParams(rawUrl string, keys []string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for _, v := range keys {
		m.Del(v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// RawURLSetParam 根据key和value对url地址参数进行重新赋值
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2;key=page,value=3
// will get http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=3#name
func RawURLSetParam(rawUrl, key, value string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	m.Set(key, value)
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// RawURLSetParams 根据map对url地址参数进行重新赋值
func RawURLSetParams(rawUrl string, params map[string]string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return rawUrl
	}

	m := stUrl.Query()
	for k, v := range params {
		m.Set(k, v)
	}
	stUrl.RawQuery = m.Encode()
	return stUrl.String()
}

// RawUrlGetDomain 获取请求地址的域名
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2
// will get www.aspxfans.com
func RawUrlGetDomain(rawUrl string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return stUrl.Hostname()
}

// RawUrlGetPort 获取请求地址的端口
// if rawUrl=http://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2
// will get 8080
func RawUrlGetPort(rawUrl string) string {
	stUrl, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return stUrl.Port()
}

// RawQueryGetParam 从请求参数字符串(请求地址?后面的字符串)中通过key获取参数值
func RawQueryGetParam(rawQuery, key string) (string, error) {
	queries, err := url.ParseQuery(rawQuery)
	if err != nil {
		return "", err
	}

	if v, ok := queries[key]; ok && len(v) > 0 {
		return v[0], nil
	}
	return "", errors.New("no param")
}

// RawQueryGetParams 从请求参数字符串(请求地址?后面的字符串)中通过key获取参数值
func RawQueryGetParams(rawQuery, key string) ([]string, error) {
	queries, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}

	if v, ok := queries[key]; ok && len(v) > 0 {
		return v, nil
	}
	return nil, errors.New("no param")
}
