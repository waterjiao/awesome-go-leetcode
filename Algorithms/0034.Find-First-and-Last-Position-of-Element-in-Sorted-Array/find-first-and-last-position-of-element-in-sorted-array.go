package _034_Find_First_and_Last_Position_of_Element_in_Sorted_Array

/**
思路：
通过二分查找先找到一个target
1、如果找到
则两次for循环，一次从mid往前找，一次从mid往后找
找到第一个和最后一个是target的数值
如果mid只有一个，返回正好就是[mid, mid]
2、如果没找到
返回[-1, -1]
3、如果nums是空
返回[-1, -1]
 */
func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	first, second := -1, -1
	result := make([]int, 0)
	l, r = mid, mid
	for ; l >= 0; l-- {
		if nums[l] == target {
			first = l
		}
	}
	for ; r <= len(nums)-1; r++ {
		if nums[r] == target {
			second = r
		}
	}
	return append(result, first, second)
}
