package main

import (
	"fmt"
)

func main() {
	intSlice := make([]int, 11)
	for i := 0; i < len(intSlice); i++ {
		intSlice[i] = i
	}

	for _, number := range intSlice {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}
