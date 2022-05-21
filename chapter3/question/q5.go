package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func CopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	srcLimited := io.LimitReader(src, n)
	written, err = io.Copy(dst, srcLimited)
	return
}

func main() {
	r := strings.NewReader("test text")
	// 書き込み先のバッファを用意
	buf := bytes.NewBufferString("")
	fmt.Println(buf)
	written, err := CopyN(buf, r, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(written)
	fmt.Println(buf)
}
