package gfile

import (
	"os"
	"path"
	"testing"
)

func TestRecurveListDirFile(t *testing.T) {
	files, err := RecurveListDirFile("E:\\Share\\产品研发\\科研创新服务平台\\科研创新服务平台\\files")
	if err != nil {
		t.Error(err)
	}
	for _, item := range files {
		t.Log(item)
	}
}

func TestPath(t *testing.T) {
	filePath2 := path.Join("v1", "v2", "v3/v4", "v6.exe")
	t.Log(filePath2)
	t.Log("D:\\_SYNC_PRESS_BOOK\\_TMP\\中职\\" + string(os.PathSeparator))

	fp := PathJoin("D:\\_SYNC_PRESS_BOOK\\_TMP\\中职\\\\", "中国", "hello.jpg")
	t.Log(fp)
}
