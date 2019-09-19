package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			/*
				To simulate concurrent updates, we’ll start 50 goroutines
				that each increment the counter about once a millisecond.
			*/
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
