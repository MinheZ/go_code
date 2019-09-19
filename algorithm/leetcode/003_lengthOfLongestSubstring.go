package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	positions := make([]int, 256)

	for i := 0; i < len(positions); i++ {
		positions[i] = -1
	}
	left, maxSize := 0, 0
	for i := 0; i < len(s); i++ {
		if positions[s[i]] >= left {
			left = positions[s[i]] + 1
		} else if i+1-left > maxSize {
			maxSize = i + 1 - left
		}
		positions[s[i]] = i
	}
	return maxSize
}
