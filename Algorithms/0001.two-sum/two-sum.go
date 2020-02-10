package problem0001

func twoSum(nums []int, target int) []int {
	//创建map，保存键值对，键为nums中的值，值为对应的下标
	index := make(map[int]int, len(nums))
	for i, v := range nums {
		//当有值存在于index中时，此时的值为target-v，返回时j在i前
		if j, ok := index[target-v]; ok {
			return []int{j, i}
		}
		//相同值保存最小的下标
		if _, ok := index[v]; !ok {
			index[v] = i
		}
	}
	return nil
}
