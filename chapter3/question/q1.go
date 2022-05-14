package main

import (
	"io"
	"os"
)

const BUFSIZE = 1024

func main() {
	old, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer old.Close()

	new, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	new.Close()
	io.Copy(new, old)
}
