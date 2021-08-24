package sort

func SelectSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 0; i < n; i++ {
		tmp := a[i]
		j := i
		min := a[j]
		index := j
		for ; j < n; j++ {
			if a[j] < min {
				min = a[j]
				index = j
			}
		}
		a[i] = min
		a[index] = tmp
	}
	return a
}
