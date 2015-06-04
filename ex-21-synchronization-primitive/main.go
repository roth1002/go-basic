package main

import (
	"fmt"
	"sync"
	"time"
)

/*
When the mutex (m) is locked any other attempt to lock it will block until it is unlocked.
Great care should be taken when using mutexes or the synchronization primitives provided in the sync/atomic package.

Traditional multithreaded programming is difficult;
it's easy to make mistakes and those mistakes are hard to find, since they may depend on a very specific,
relatively rare, and difficult to reproduce set of circumstances.
One of Go's biggest strengths is that the concurrency features it provides are much easier to understand and use properly than threads and locks.
*/

func main() {
	mutex := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			mutex.Lock()
			fmt.Println(i, "Lock Start")

			time.Sleep(time.Second * 1)

			fmt.Println(i, "Lock End")
			mutex.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
