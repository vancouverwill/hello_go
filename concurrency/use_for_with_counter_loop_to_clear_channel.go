package concurrency

import (
	"fmt"
	"time"
)

func UseLoopWithCounterToClearChannel() {
	fmt.Println("RUN UseLoopWithCounterToClearChannel")

	aChan := make(chan int, 10)
	
    go func() {
        for i := 0; i < 10; i++ {
            time.Sleep(100 * time.Millisecond)
            aChan <- i
        }
        close(aChan)
    }()
    
    for i := 0; i < 10; i++ {
        aInt := <-aChan
        fmt.Println(aInt)
    }
	

}
