package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic((err))
	}
	// Go 言語にはなんでも表示できる %v というフォーマット指定子があり、
	// プリミティ ブ型でもそうでない型でも String() メソッドがあればそれを表示に使って出力してくれます。
	// これも fmt.Stringer インタフェースとして定義されています。
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
	// 第一引数にio.Writerを渡せればなんでもOK
	fmt.Fprintf(file, "Write with file at %v", time.Now())
}
