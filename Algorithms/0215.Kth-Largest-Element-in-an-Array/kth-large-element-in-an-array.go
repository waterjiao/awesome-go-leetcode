package _215_Kth_Largest_Element_in_an_Array

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, start, end, k int) int {
	pivot := partition(nums, start, end)
	if pivot > k {
		return quickSelect(nums, start, pivot-1, k)
	} else if pivot < k {
		return quickSelect(nums, pivot+1, end, k)
	}
	return nums[pivot]
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
