package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
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

func textChunk(text string) io.Reader {
	var buffer bytes.Buffer
	byteData := []byte(text)
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)

	crc := crc32.NewIEEE()
	crc.Write(byteData)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer

}

func main() {
	err := download()
	if err != nil {
		panic(err)
	}

	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	io.Copy(newFile, chunks[0])
	io.Copy(newFile, textChunk("ASCII PROGRAMING++"))
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	chunks = readChunks(newFile)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
	// chunk 'IHDR' (13 bytes)
	// chunk 'sRGB' (1 bytes)
	// chunk 'IDAT' (473761 bytes)
	// chunk 'IEND' (0 bytes)
}

func download() error {
	resp, err := http.Get("https://upload.wikimedia.org/wikipedia/en/7/7d/Lenna_%28test_image%29.png")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create("Lenna.png")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
