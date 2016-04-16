package concurrency

import (
    "fmt"
)

func UseRangeToClearChannel() {
    fmt.Println("RUN UseLoopToClearChannel")

    aChan := make(chan int, 10)
    
     for i := 0; i < 10; i++ {
        aChan <- i
    }
    close(aChan)

    for aInt := range aChan {
        fmt.Println(aInt)
    }
    
}


