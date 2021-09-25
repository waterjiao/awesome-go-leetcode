package _322_Coin_Change

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11

	result := coinChange(coins, amount)
	fmt.Print(result)
}
