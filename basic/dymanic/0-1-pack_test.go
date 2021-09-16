package dymanic

import "testing"

func TestCount(t *testing.T) {
	nums := []int{2,2,4,6,3}
	num := 9

	result := count(nums, num)
	t.Logf("results is %v", result)
}
