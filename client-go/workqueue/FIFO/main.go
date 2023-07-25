package main

import (
	"fmt"
	"k8s.io/client-go/util/workqueue"
)

func FIFO() {
	queue := workqueue.New()
	fmt.Println(queue.Len())

	queue.Add("a") // queue: [a] dirty: [a] processing: []

	// a is processing
	queue.Get() // queue: [] dirty: [] processing: [a]

	// enqueue a while a is processing
	queue.Add("a") // queue: [] dirty: [a] processing: [a]

	queue.Done("a") // queue: [a] dirty: [a] processing: []

	// get A again
	item, _ := queue.Get() //
	fmt.Println(item)      // queue: [] dirty: [] processing: [a]

}

func main() {
	FIFO()

}
