package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // 不关闭通道会产生死锁
	for ele := range queue {
		fmt.Println(ele)
	}
}
