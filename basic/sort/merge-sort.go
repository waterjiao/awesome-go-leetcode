package sort

func MergeSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	return mergeSortC(a, 0, n-1)
}

func mergeSortC(a []int, p, r int) []int {
	if p > r {
		return a
	}
	q := (p + r)/2
	mergeSortC(a, p, q)
	mergeSortC(a, q+1, r)
	merge(a, a[p:q], a[q+1:r])
	return a
}

func merge(a, a1, a2 []int) []int {
	
}