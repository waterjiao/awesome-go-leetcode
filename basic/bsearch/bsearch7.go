package bsearch

func bsearch7(a []int, value int) int {
	if len(a) < 1 {
		return -1
	}
	l, r := 0, len(a)-1
	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if a[mid] <= value {
			if mid == len(a)-1 || a[mid+1] > value {
				break
			} else {
				l = mid + 1
			}
		} else {
			r = mid - 1
		}
	}
	return a[mid]
}
