package _015_3Sum

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			r := make([]int, 0)
			if nums[left]+nums[right] == target {
				r = append(r, nums[i], nums[left], nums[right])
				result = append(result, r)
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
