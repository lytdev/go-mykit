/*
 * @Author       : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @Date         : 2022-10-12 19:23:50
 * @LastEditors  : 刘元涛 snoopy_718@mails.ccnu.edu.cn
 * @FilePath     : \go-myexcel\gformt\dateformat_test.go
 * @Description  :
 * Copyright (c) 2022 by 刘元涛 snoopy_718@mails.ccnu.edu.cn, All Rights Reserved.
 */
package gformt

import (
	"testing"
)

func Test_Format(t *testing.T) {
	data := map[string]string{
		"pubDate1 ":  "2022.02",
		"pubDate2 ":  "2022.2",
		"pubDate3 ":  "2022.02.01",
		"pubDate4 ":  "2022.02.1",
		"pubDate5 ":  "2022.2.01",
		"pubDate6 ":  "2022.2.1",
		"pubDate7 ":  "2022/02",
		"pubDate8 ":  "2022/02/01",
		"pubDate9 ":  "2022/02/1",
		"pubDate10 ": "2022/2",
		"pubDate11 ": "2022/2/1",
		"pubDate12 ": "2022/2/01",
		"pubDate13 ": "2022-02",
		"pubDate14 ": "2022-02-01",
		"pubDate15 ": "2022-2-01",
		"pubDate16 ": "2022-2-1",
		"pubDate17 ": "2022年",
		"pubDate18 ": "2022年2月",
		"pubDate19 ": "2022年02月",
		"pubDate20 ": "2022年02月01日",
		"pubDate21 ": "2022年2月1日",
		"pubDate22 ": "2022年02月1日",
		"pubDate23 ": "2022年2月01日",
		"pubDate24 ": "202202",
		"pubDate25 ": "20220201",
		"pubDate26 ": "2022.2.120",
		"pubDate27 ": "2022/0262",
		"pubDate28 ": "2022/13/12",
		"pubDate29 ": "2022/12/32",
	}

	for k, v := range data {
		p1, err := GetFormatDateStr(v)
		if err != nil {
			t.Error("请求错误:", err)
		} else {
			t.Log(k + ":" + p1)
		}

	}

}
