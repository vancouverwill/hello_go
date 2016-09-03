package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wgA sync.WaitGroup

func slowAsyncTask(depth, taskNumber int) {
	defer wgA.Done()
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("we are at depth:%d taskNumber:%d\n", depth, taskNumber)
	if depth > 0 {
		wgA.Add(1)
		go slowAsyncTask(depth-1, taskNumber)
	}
}

func RunAsynchronousTasks() int {
	fmt.Println("begin")
	for i := 0; i < 1000; i++ {
		depth := 30
		wgA.Add(1)
		go slowAsyncTask(depth, i)
	}
	wgA.Wait()
	fmt.Println("end")

	return 5
}
