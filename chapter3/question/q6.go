package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

// 文字列のバイト数
// http://itdoc.hitachi.co.jp/manuals/3000/30003D0820/GD080450.HTM#:~:text=%E3%82%B7%E3%83%95%E3%83%88JIS%E3%81%A7%E3%81%AF%EF%BC%8C1%E6%96%87%E5%AD%97,%E5%8F%AF%E5%A4%89%E9%95%B7%E3%81%AB%E3%81%AA%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82
func main() {
	var stream io.Reader
	a := io.NewSectionReader(programming, 5, 1)
	s := io.NewSectionReader(system, 0, 1)
	c := io.NewSectionReader(computer, 0, 1)
	i := io.NewSectionReader(programming, 8, 1)
	i2 := io.NewSectionReader(programming, 8, 1)
	n := strings.NewReader("\n")
	stream = io.MultiReader(a, s, c, i, i2, n) // iを使い回すことができないのでi2が必要
	io.Copy(os.Stdout, stream)
}
