package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for i := 1; i <= 5; i++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker ", id, "start job ", job)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job ", job)
		results <- job * 2
	}
}
