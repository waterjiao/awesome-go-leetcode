package _046_Permutations

var results [][]int

func permute(nums []int) [][]int {
	results = make([][]int, 0)
	track := make([]int, 0)
	backtrack(nums, track)
	return results
}

func backtrack(nums []int, track []int) {
	if len(track) == len(nums) {
		result := make([]int, len(track))
		copy(result, track)
		results = append(results, result)
	}
	for i := 0; i < len(nums); i++ {
		if isTrack(nums[i], track) {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, track)
		track = track[:len(track)-1]
	}
}

func isTrack(num int, track []int) bool {
	for _, r := range track {
		if r == num {
			return true
		}
	}
	return false
}