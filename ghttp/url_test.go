package ghttp

import "testing"

func TestUrl(t *testing.T) {
	rawUrl := "https://www.aspxfans.com:8080/news/index.asp?boardID=520&page=1&page=2&show=true&type=4"
	t.Log("-------------s1-------------")
	s1, _ := RawURLGetParam(rawUrl, "page")
	t.Log(s1)
	t.Log("-------------s2-------------")
	s2, _ := RawURLGetParams(rawUrl, "page")
	t.Log(s2)
	t.Log("-------------s3-------------")
	s3, _ := RawURLGetAllParams(rawUrl)
	t.Log(s3)

	t.Log("------------s4--------------")
	s4 := RawURLAddParam(rawUrl, "isDel", "0")
	t.Log(s4)
	t.Log("-------------s5-------------")
	s5 := RawURLAddParams(rawUrl, map[string]string{
		"a": "1",
		"b": "2",
	})
	t.Log(s5)
	t.Log("-------------s6-------------")
	s6 := RawURLDelParam(rawUrl, "page")
	t.Log(s6)
	t.Log("-------------s7-------------")
	s7 := RawURLDelParams(rawUrl, []string{"type", "page"})
	t.Log(s7)
	t.Log("-------------s8-------------")
	s8 := RawURLSetParam(rawUrl, "page", "100")
	t.Log(s8)
	t.Log("-------------s9-------------")
	s9 := RawURLSetParams(rawUrl, map[string]string{
		"page": "100",
		"type": "101",
	})
	t.Log(s9)
	t.Log("-------------s10-------------")
	s10, _ := RawQueryGetParam(rawUrl, "page")
	t.Log(s10)
	t.Log("-------------s11-------------")
	s11, _ := RawQueryGetParams(rawUrl, "page")
	t.Log(s11)
}
