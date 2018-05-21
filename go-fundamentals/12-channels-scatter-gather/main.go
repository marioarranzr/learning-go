package main

import (
	"fmt"
	"math/rand"
)

func main() {

	n := 100

	resultC := make(chan int, n)

	// Scatter
	for i := 0; i < n; i++ {
		go func() {
			resultC <- rand.Intn(100)
		}()
	}

	// Gather
	for i := 0; i < n; i++ {
		fmt.Println(<-resultC)
	}
}
