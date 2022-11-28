package gid

import (
	"strconv"
	"testing"
)

// https://github.com/satori/go.uuid/blob/master/generator_test.go
func TestUuid(t *testing.T) {
	t.Log("uuid1:" + Uuid1())
	t.Log("uuid2:" + Uuid2())
	t.Log("uuid3:" + Uuid3())
	t.Log("uuid4:" + Uuid4())
	t.Log("fastUuid:" + FastUuid())
}

func TestSnowflake(t *testing.T) {
	snowflake, _ := NewSnowflake(0)
	for i := 0; i < 1000; i++ {
		t.Log("雪花算法ID：" + strconv.Itoa(int(snowflake.NextId())))
	}
}
func TestNanoId(t *testing.T) {
	for i := 0; i < 1000; i++ {
		idStr1, _ := NewNanoId()
		idStr2, _ := GenerateNanoId("abcde", 8)
		t.Log("NanoId1：" + idStr1)
		t.Log("NanoId2：" + idStr2)
	}
}
