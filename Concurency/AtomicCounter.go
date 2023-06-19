package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			//directly incrementing is now  work properly
			//counter++
			atomic.AddUint64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter is ", counter)
}
