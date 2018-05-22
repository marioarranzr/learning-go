package main

import (
	"fmt"
	"go/go-web-dev/02-queue/queue"
)

func main() {

	_queue := queue.NewQueue(1)
	_queue.Push(&queue.Node{Value: 1})
	_queue.Push(&queue.Node{Value: 2})
	_queue.Push(&queue.Node{Value: 5})
	_queue.Push(&queue.Node{Value: 6})
	_queue.Push(&queue.Node{Value: 7})
	_queue.Push(&queue.Node{Value: 8})
	_queue.Push(&queue.Node{Value: 9})
	_queue.Push(&queue.Node{Value: 10})

	fmt.Println(_queue.Pop(), _queue.Pop(), _queue.Pop(), _queue.Pop(), _queue.Pop(),
		_queue.Pop(), _queue.Pop(), _queue.Pop(), _queue.Pop(), _queue.Pop())

}
