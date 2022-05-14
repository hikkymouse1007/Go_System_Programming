package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReadr")
	teeReader := io.TeeReader(reader, &buffer)
	// データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)
	// けどバッファに残っている
	fmt.Println(buffer.String())
}
