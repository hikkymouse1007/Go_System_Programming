package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader
	// 先頭の8バイトをスキップ
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		chunks = append(chunks,
			// NewSectionReader returns a SectionReader that reads from r
			// starting at offset off and stops with EOF after n bytes
			io.NewSectionReader(file, offset, int64(length)+12))

		// 次のチャンクの先頭に移動
		// 現在位置は長さを読み終わった箇所なので
		// チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks

}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
	// chunk 'IHDR' (13 bytes)
	// chunk 'sRGB' (1 bytes)
	// chunk 'IDAT' (473761 bytes)
	// chunk 'IEND' (0 bytes)
}
