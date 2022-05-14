package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("test.txt")
	if err != nil {
		panic(err)
	}
	reader := strings.NewReader("this is q4\n")
	io.Copy(writer, reader)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
