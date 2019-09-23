package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func main() {
	a := []int{1, 2, 3, -4}
	sort.Ints(a)
	fmt.Println(a)
}

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
