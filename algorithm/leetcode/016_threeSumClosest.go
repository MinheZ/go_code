package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	var target = 1
	res := threeSumClosest(nums, target)
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	var closestNum = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for {
			if left >= right {
				break
			}
			threeSum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(threeSum-target)) < math.Abs(float64(closestNum-target)) {
				closestNum = threeSum
			}
			if threeSum > target {
				right--
			} else if threeSum < target {
				left++
			} else {
				return target
			}
		}
	}
	return closestNum
}
