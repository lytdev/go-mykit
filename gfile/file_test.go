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
	s2, _ := FileDir(filePath)
	t.Log(s1)
	t.Log(s2)

}
func TestPath(t *testing.T) {
	filePath2 := path.Join("v1", "v2", "v3/v4", "v6.exe")
	t.Log(filePath2)
	t.Log("D:\\_SYNC_PRESS_BOOK\\_TMP\\中职\\" + string(os.PathSeparator))

	fp := PathJoin("D:\\_SYNC_PRESS_BOOK\\_TMP\\中职\\\\", "中国", "hello.jpg")
	t.Log(fp)
}

func TestFileRead(t *testing.T) {
	fp := "D:\\_TMP\\demo.txt"
	lines, err := ReadWithLine(fp)
	if err != nil {
		t.Error(err)
	}
	for _, line := range lines {
		t.Log(line)
	}
}

func TestFileCopy(t *testing.T) {
	fp := "D:\\_TMP\\demo.txt"
	fp2 := "D:\\_TMP\\22\\demo1.txt"
	n, err := CopyFile(fp, fp2)
	if err != nil {
		t.Error(err)
	}
	t.Log("拷贝的字节数:", n)
}
