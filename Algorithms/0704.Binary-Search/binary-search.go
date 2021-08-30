package _704_Binary_Search

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid, index := -1, -1
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return index
}
