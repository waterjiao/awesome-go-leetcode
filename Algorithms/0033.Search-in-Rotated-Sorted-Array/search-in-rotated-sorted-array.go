package _033_Search_in_Rotated_Sorted_Array

// 变体二分查找需要考虑的
// 1. 必须考虑全边界，如两种情况，必然是一个大于等于，一个小于；或者是一个小于等于，一个大于
// 2. 要注意二分的使用场景，有序，即使是当前题目中旋转，也必然是二分之后，一半有序，一半可能有序（可能有序很关键，有可能是有序的，也有可能是无序的）
// 3. 再考虑有序中的处理和无序中的处理
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid, index := -1, -1
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[l] <= nums[mid] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target <= nums[r] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return index
}
