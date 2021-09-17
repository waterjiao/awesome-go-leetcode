package _509_Fibonacci_Number

func fib(n int) int {
	mano := make([]int, n+1)
	return fibNum(n, mano)
}

func fibNum(n int, mano []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if mano[n] != 0 {
		return mano[n]
	}
	mano[n] = fibNum(n-1, mano) + fibNum(n-2, mano)
	return mano[n]
}
