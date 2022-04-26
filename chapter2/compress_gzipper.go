package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	// gzファイルを作成
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	// gzファイル内のファイルを作成
	writer.Header.Name = "test.txt"
	// test.txtに書き込み
	io.WriteString((writer), "gzip.Writer example\n")
	writer.Close()
}
