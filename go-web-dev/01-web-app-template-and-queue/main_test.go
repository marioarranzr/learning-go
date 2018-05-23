package main

import (
	"testing"
)

func TestLenQueue(t *testing.T) {
	numElementsSlice := len(l)
	numElementsQueue := q.Len()
	if numElementsSlice != numElementsQueue {
		t.Errorf("Expected queue length of %v but got %v", numElementsSlice, numElementsQueue)
	}
}
