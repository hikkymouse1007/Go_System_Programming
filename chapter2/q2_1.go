package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("question2_1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "This is a %dst %s answered at %v\n", 1, "question", time.Now())
	fmt.Fprintf(file, "Your favorite decimal number is %f", 1.2235)
	file.Close()
}
