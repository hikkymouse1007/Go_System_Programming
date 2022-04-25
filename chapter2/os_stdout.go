package main

import (
	"os"
)

func main() {
	// fmt.Println() では、最終的 に os.Stdout の Write() メソッドを呼び出していました。
	// それと等価なコードは次 のとおりです。
	// 標準エラー出力に出力するための os.Stderr もあります。
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
