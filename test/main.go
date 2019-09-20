package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func main() {
	// var str = []string{}
	strs := sort.StringSlice{"abc", "d"}
	fmt.Println(strs)
}

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return len(p[i]) < len(p[j]) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
