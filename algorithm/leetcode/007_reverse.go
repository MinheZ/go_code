package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x = x / 10

	}
	limit := 2147483648
	if res > limit-1 || res < -limit {
		res = 0
	}
	return res
}
