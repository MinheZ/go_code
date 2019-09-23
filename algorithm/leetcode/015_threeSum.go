package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{-2, 1, 1, 1, 1}
	fmt.Println(threeSum(arr))
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-1; i++ {
		left, right := i+1, len(nums)-1
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		for left < right {
			//if right < len(nums)-1 && nums[right] == nums[right+1] {
			//	right--
			//	continue
			//}
			twoSum := nums[left] + nums[right]
			if twoSum == target {
				var ans []int
				ans = append(ans, nums[i], nums[left], nums[right])
				ret = append(ret, ans)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if twoSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}
