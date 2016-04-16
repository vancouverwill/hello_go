package concurrency

import (
	"fmt"
	"strconv"
	"time"
)

func UseLoopToClearChannel() {
	fmt.Println("RUN UseLoopToClearChannel")

	aChan := make(chan string, 10)
	done := make(chan bool)

	go func() {
		for {
			aInt, ok := <-aChan
			if !ok {
				done <- true
				break
			}
			fmt.Println(aInt)
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		aChan <- strconv.Itoa(i)
	}
	shortTicker := time.NewTicker(time.Millisecond * 300)
	longTicker := time.NewTicker(time.Millisecond * 600)
	go func() {
		for ts := range shortTicker.C {
			aChan <- "shortTicker at" + ts.String()

		}
	}()
	go func() {
		for tl := range longTicker.C {
			aChan <- "longTicker at" + tl.String()
		}
	}()
	time.Sleep(time.Second * 4)
	shortTicker.Stop()
	longTicker.Stop()

	close(aChan)

	_, ok := <-done
	if ok {
		fmt.Println("all done")
	}
}
