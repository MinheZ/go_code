package main

import (
	"fmt"
)

func main() {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	var max = 0
	var left, right = 0, len(height) - 1
	for left < right {
		var a int
		if height[left] < height[right] {
			a = height[left]
		} else {
			a = height[right]
		}
		var b = right - left
		var area = a * b
		if max < area {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
