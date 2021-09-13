package _402_Remove_K_Digits

import "strings"

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	num = strings.TrimLeft(string(stack), "0")
	if num == "" {
		return "0"
	}
	return num
}

/**
	以下思路有误
按照贪心算法，不能求解最正确的解
 */
func removeKdigits1(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	i := 1
	for i < len(num) && k > 0 {
		if num[i] < num[i-1] {
			num = num[:i-1] + num[i:]
			k--
		} else {
			i++
		}
	}
	if k > 0 {
		num = num[k:]
	}
	num = strings.TrimLeft(num, "0")
	if num == "" {
		return "0"
	}
	return num
}
