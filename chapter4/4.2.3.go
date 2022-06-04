package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		defer close(result)

		result <- 2
		fmt.Sprintf("initial result: %d", result)
		for i := 3; i < 1000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			fmt.Printf("l is %d \n", l)
			found := false
			for j := 3; j < l+1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
	}()
	fmt.Printf("final result: %d", result)
	return result
}

func main() {
	pn := primeNumber()
	for n := range pn {
		fmt.Println(n)
	}
}
