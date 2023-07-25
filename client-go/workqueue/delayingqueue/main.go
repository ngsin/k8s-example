package main

import (
	"fmt"
	"k8s.io/client-go/util/workqueue"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	queue := workqueue.NewDelayingQueue()
	queue.AddAfter("a", time.Second*5)

	go func() {
		fmt.Println(queue.Get())
		wg.Done()
	}()

	wg.Wait()

}
