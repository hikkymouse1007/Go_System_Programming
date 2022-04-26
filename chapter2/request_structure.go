package main

import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write((os.Stdout))
	// GET / HTTP/1.1
	// Host: ascii.jp
	// User-Agent: Go-http-client/1.1
	// X-Test: ヘッダーも追加できます
}
