package main

import (
	"fmt"
	"math"

	"github.com/facebookgo/pqueue"
)

func main() {
	c := 100
	pq := pqueue.New(c)

	for i := 0; i < c; i++ {
		item := &pqueue.Item{
			Value:    i,
			Priority: math.MaxInt64,
			Index:    i,
		}
		pq.Push(item)

	}

	fmt.Println(pq.Pop())

}
