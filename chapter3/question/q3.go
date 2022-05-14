package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()
	// archive.zip内にnewfile.txtを作る
	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	reader := strings.NewReader("this is Q3 \n")
	io.Copy(writer, reader)
}
