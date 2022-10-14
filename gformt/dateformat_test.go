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
	"sort"
	"testing"
)

func Test_Format(t *testing.T) {
	dataMap := map[string]string{
		"pubDate01 ": "2022.02",
		"pubDate02 ": "2022.2",
		"pubDate03 ": "2022.02.01",
		"pubDate04 ": "2022.02.1",
		"pubDate05 ": "2022.2.01",
		"pubDate06 ": "2022.2.1",
		"pubDate07 ": "2022/02",
		"pubDate08 ": "2022/02/01",
		"pubDate09 ": "2022/02/1",
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
		"pubDate30 ": "2022/12/21 09:59",
		"pubDate31 ": "2022/12/30 9:59",
		"pubDate32 ": "2022/12/31 9:1:10",
		"pubDate33 ": "2022.12.31 9:1:10",
		"pubDate34 ": "2022-12-31 9:1:10",
		"pubDate35 ": "2022-12-31 9:1:10 AM",
		"pubDate36 ": "2022年12月31 9:1:10 PM",
		"pubDate37 ": "2022年12月31日 9时1分10秒",
		"pubDate38 ": "2022年12月31日9时1分10秒",
		"pubDate39 ": "2022-12-31 9:1:61",
		"pubDate40 ": "2022-12-31 9:1 PM",
		"pubDate41 ": "2022年12月31日 9:62:10",
		"pubDate42 ": "2022-12-31 9:1:10 PM",
		"pubDate43 ": "2022-12-31 25:1:10 AM",
	}

	keys := make([]string, 0, len(dataMap))
	for k := range dataMap {
		keys = append(keys, k)
	}

	//对切片进行排序
	sort.Strings(keys)

	for _, key := range keys {
		val := dataMap[key]
		p1, err := GetFormatDateStr(val)
		if err != nil {
			t.Log(err.Error())
		} else {
			t.Logf("key:%s --> value:%s --> format:%s", key, val, p1)
		}
	}

}
