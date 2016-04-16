package concurrency

import "fmt"
import "sync/atomic"

func AtomicExample() {
	fmt.Println("Atomic_example")
	var ops uint64

	for i := 0; i < 10; i++ {
		atomic.AddUint64(&ops, 1)
		fmt.Println(ops)
	}

	for i := 0; i < 10; i++ {
		atomic.AddUint64(&ops, ^uint64(0))
		fmt.Println(ops)
	}
}
