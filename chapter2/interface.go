package main

import (
	"fmt"
)

// インタフェースを定義
type Talker interface {
	Talk()
}

// 構造体を宣言
type Greeter struct {
	name string
}

// 構造体はTalker インタフェースで定義されているメソッドを持っている
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}
func main() {
	var talker Talker

	// インタフェースを満たす構造体のポインタは代入できる
	// 初期化パラメータを与えて Greeter型の構造体のインスタンスを作成し、そのポインタを代入
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
