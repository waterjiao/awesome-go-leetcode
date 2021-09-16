package dymanic

var results [][]int
var max int

func count(num []int, weight int) [][]int {
	results = make([][]int, 0)
	max = weight
	result := make([]int, 0)
	backtrack(num, 0, 0, result)
	return results
}

func backtrack(nums []int, index, w int, result []int) {
	if w == max || index == len(nums)-1 {
		if w <= max {
			r := make([]int, len(result))
			copy(r, result)
			results = append(results, r)
		}
		return
	}
	if w + nums[index] <= max {
		w = w + nums[index]
		result = append(result, nums[index])
		backtrack(nums, index+1, w, result)
		w = w - nums[index]
		result = result[:len(result)-1]
	}
}
