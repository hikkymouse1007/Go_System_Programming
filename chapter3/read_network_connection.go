package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	// ここの書き込みを行わないとレスポンスが返ってこない
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp.\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}
