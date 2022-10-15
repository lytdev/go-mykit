package gfrmt

import (
	"regexp"
	"sort"
	"testing"
)

func TestRegexp(t *testing.T) {
	numStr := "2022.02_6"
	if matched, _ := regexp.MatchString("^(\\d{4}[.]\\d{1,2}[.]\\d{1,2})$", numStr); matched {
		t.Log("匹配成功")
	} else {
		t.Error("匹配失败")
	}
}

func TestFormat(t *testing.T) {
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
		"pubDate26 ": "2022.2.120", //err
		"pubDate27 ": "2022/0262",  //err
		"pubDate28 ": "2022/13/12", //err
		"pubDate29 ": "2022/12/32", //err
		"pubDate30 ": "2022/12/21 09:59",
		"pubDate31 ": "2022/12/30 9:59",
		"pubDate32 ": "2022/12/31 9:1:10",
		"pubDate33 ": "2022.12.31 9:1:10",
		"pubDate34 ": "2022-12-31 9:1:10",
		"pubDate35 ": "2022-12-31 9:1:10 AM",
		"pubDate36 ": "2022年12月31 9:1:10 PM",
		"pubDate37 ": "2022年12月31日 9时1分10秒",
		"pubDate38 ": "2022年12月31日9时1分10秒",
		"pubDate39 ": "2022-12-31 9:1:61", //err
		"pubDate40 ": "2022-12-31 9:1 PM",
		"pubDate41 ": "2022年12月31日 9:62:10", //err
		"pubDate42 ": "2022-12-31 9:1:10 PM",
		"pubDate43 ": "2022-12-31 25:1:10 AM", //err
		"pubDate44 ": "2022_12-31 25:1:10 AM", //err
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
			t.Logf("日期的名称:%s --> 错误信息:%s", key, err.Error())
		} else {
			t.Logf("日期的名称:%s --> 日期的值:%s --> format:%s", key, val, p1)
		}
	}

}
