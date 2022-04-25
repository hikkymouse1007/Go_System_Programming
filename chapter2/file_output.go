package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic((err))
	}
	// Write()が受け取るのは文字列ではなくてバイト列なので、
	// 変換を行ってから Write() メソッドに渡しています。
	file.Write([]byte("os.File example\n"))
	file.Close()
}
