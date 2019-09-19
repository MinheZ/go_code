package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var y = x
	var reverse = 0

	for y != 0 {
		reverse = reverse*10 + y%10
		y /= 10
	}
	return reverse == x
}
