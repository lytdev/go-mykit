package gfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadWithLine 按行读取文件的文本
func ReadWithLine(fp string) ([]string, error) {
	var lines = make([]string, 0)
	// 创建句柄
	fi, err := os.Open(fp)
	if err != nil {
		return lines, err
	}

	// 创建 Reader
	r := bufio.NewReader(fi)

	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			return make([]string, 0), nil
		}
		if err == io.EOF {
			if len(line) > 0 {
				lines = append(lines, line)
			}
			break
		}
		lines = append(lines, line)
	}
	return lines, nil
}

// CopyFile 复制文件
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	//返回拷贝的字节数和发生错误的信息
	n, err := io.Copy(destination, source)
	return n, err
}
