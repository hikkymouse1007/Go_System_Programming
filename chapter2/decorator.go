package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	// 複数のwriterを一つのio.Writerオブジェクトとして受け取る
	writer := io.MultiWriter(file, os.Stdout)
	// writeに渡したfileと標準出力両方に対して、io.WriteStringでの第二引数の文字列を書き込む
	// 作成されたtxtファイル、コンソールの出力両方に同じ文字列が書き込まれる。
	io.WriteString(writer, "io.MultiWriter example\n")
}
