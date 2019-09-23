package main

import (
	"fmt"
	"sort"
)

type Integer int

func main() {
	a := []int{1, 2, 3, -4}
	sort.Ints(a)
	fmt.Println(a)
}
