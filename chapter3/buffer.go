package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var buffer1 bytes.Buffer
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	buffer3 := bytes.NewBufferString("初期文字列")
	fmt.Println(buffer1)
	fmt.Println(buffer2)
	fmt.Println(buffer3)
	bReader1 := bytes.NewReader([]byte{0x10, 0x20, 0x30})
	bReader2 := bytes.NewReader([]byte("文字列をバイト配列にキャストして設定"))
	sReader := strings.NewReader("Readerの出力内容は文字列で渡す")
	fmt.Println(bReader1)
	fmt.Println(bReader2)
	fmt.Println(sReader)
	lReader := io.LimitReader(sReader, 16)
	fmt.Println(lReader)
}
