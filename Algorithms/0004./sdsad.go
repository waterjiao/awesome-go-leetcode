package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) > 0 {
		if len(nums2)%2 == 0 {
			return float64(nums2[len(nums2)/2-1]+nums2[len(nums2)/2]) / 2
		}
		return float64(nums2[(len(nums2)-1)/2])
	}
	if len(nums2) == 0 && len(nums1) > 0 {
		if len(nums1)%2 == 0 {
			return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2.0
		}
		return float64(nums1[(len(nums1)-1)/2])
	}
	var lastNums = make([]int, 0)
	var nums2Index = 0
	for i := 0; i < len(nums1); i++ {
		for j := nums2Index; j < len(nums2); j++ {
			if nums2[j] <= nums1[i] {
				lastNums = append(lastNums, nums2[j])
				nums2Index = j
			} else {
				nums2Index = j
				break
			}
		}
		lastNums = append(lastNums, nums1[i])
	}
	if nums2Index<len(nums2) {
		lastNums = append(lastNums, nums2[nums2Index:]...)
	}
	fmt.Println(lastNums)
	if len(lastNums)%2 == 0 {
		return float64(lastNums[len(lastNums)/2-1]+lastNums[len(lastNums)/2]) / 2.0
	}
	return float64(lastNums[(len(lastNums)-1)/2])
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	num := findMedianSortedArrays(nums1, nums2)
	fmt.Println(num)
}
