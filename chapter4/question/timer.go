package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start ticking...")
	timer := time.After(10 * time.Second)
	t := <-timer
	fmt.Println(t)
	fmt.Println("Zzzzzzz")
}
