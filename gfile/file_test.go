package gfile

import (
	"os"
	"path"
	"testing"
)

func TestRecurveListDirFile(t *testing.T) {
	dirPath := "D:\\云展活动\\数据报告\\2022年第2季"
	files, err := RecurveListDirFile(dirPath)
	if err != nil {
		t.Error(err)
	}
	for _, item := range files {
		t.Log(item)
	}
}

func TestFileDemo(t *testing.T) {
	filePath := "D:\\云展活动\\数据报告\\2022年第2季\\2022年第2季.zip"
	s1 := MainName(filePath)
	t.Log(s1)

}
func TestPath(t *testing.T) {
	filePath2 := path.Join("v1", "v2", "v3/v4", "v6.exe")
	t.Log(filePath2)
	t.Log("D:\\_SYNC_PRESS_BOOK\\_TMP\\中职\\" + string(os.PathSeparator))

	fp := PathJoin("D:\\_SYNC_PRESS_BOOK\\_TMP\\中职\\\\", "中国", "hello.jpg")
	t.Log(fp)
}
