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

func fib2(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	pre, curr := 0, 1
	for i := 2; i <= n; i++ {
		sum := pre + curr
		pre = curr
		curr = sum
	}
	return curr
}
