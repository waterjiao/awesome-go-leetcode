package _069_Sqrt_x

func mySqrt(x int) int {
	l, r := 0, x
	var result int
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return result
}
