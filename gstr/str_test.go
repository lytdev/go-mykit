package gstr

import (
	"testing"
)

func TestChunkString(t *testing.T) {
	result1 := ChunkString("1234567", 3)
	result2 := ChunkString("12345", 2)
	result3 := ChunkString("", 2)
	result4 := ChunkString("1", 2)
	t.Logf("%v\n", result1)
	t.Logf("%v\n", result2)
	t.Logf("%v\n", result3)
	t.Logf("%v\n", result4)
}
