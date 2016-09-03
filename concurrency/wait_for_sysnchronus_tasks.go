package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func slowSyncTask(depth, taskNumber int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 400)
	fmt.Printf("we are at depth:%d taskNumber:%d\n", depth, taskNumber)
	if depth > 0 {
		wg.Add(1)
		slowSyncTask(depth-1, taskNumber)
	}
}

func RunSynchronousTasks() {
	fmt.Println("begin")
	for i := 0; i < 3; i++ {
		depth := 3
		wg.Add(1)
		slowSyncTask(depth, i)
	}
	wg.Wait()
	fmt.Println("end")

}
