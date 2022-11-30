package gfile

import "testing"

func TestRecurveListDirFile(t *testing.T) {
	files, err := RecurveListDirFile("E:\\Share\\产品研发\\科研创新服务平台\\科研创新服务平台\\files")
	if err != nil {
		t.Error(err)
	}
	for _, item := range files {
		t.Log(item)
	}
}
