package main

import (
	"fmt"
)

func main() {
	buffer := make(chan string, 2)
	buffer <- "channel"
	buffer <- "buffer"
	//for b := range buffer {
	//	fmt.Println(b)
	//}
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
	close(buffer)
	_, ok := <-buffer
	fmt.Println(ok)
}
