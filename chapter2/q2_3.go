package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Content-Encoding
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json") // json 化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	writer := gzip.NewWriter(w)
	writer.Header.Name = "testttt"
	multiWriter := io.MultiWriter(writer, os.Stdout)
	// multiWriterに書き込みを行うNewEncoderを作成
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)
	writer.Flush()
	writer.Close()
}

// ここにコードを書く }

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// Zipファイルを取得
// curl http://localhost:8080 --output myzip.json.gz
