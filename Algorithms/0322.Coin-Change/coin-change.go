package _322_Coin_Change

//import "math"
const IntMax = int(^uint(0) >> 1)
var mamo []int

func coinChange(coins []int, amount int) int {
	mamo = make([]int, amount+1)
	for i := range mamo {
		mamo[i] = -2
	}
	return bp(coins, amount)
}

func bp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if mamo[amount] != -2 {
		return mamo[amount]
	}
	result := IntMax
	for _, coin := range coins {
		subProblem := bp(coins, amount-coin)
		if subProblem < 0 {
			continue
		}
		result = min(result, subProblem+1)
	}
	if result == IntMax {
		result = -1
	}
	mamo[amount] = result
	return mamo[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
