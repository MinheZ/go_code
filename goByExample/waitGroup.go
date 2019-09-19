package main

import (
	"fmt"
	"sync"
	"time"
)

/*
To wait for multiple goroutines to finish, we can use a wait group.
Compare to CountDownLatch in Java.
*/
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker1(i, &wg)
	}
	wg.Wait()
}

func worker1(id int, waitGroup *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	waitGroup.Done()
}
