package problem0035

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for {
		if low < high {
			mid := (low + high) / 2
			if nums[mid] == target {
				return mid
			}
		}
	}
}
