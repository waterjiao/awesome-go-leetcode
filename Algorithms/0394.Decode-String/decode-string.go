package _394_Decode_String

import (
	"strconv"
	"strings"
)

/**
解题思路：
遍历字符串s
1。当遇到数字时，尤其时连续数字时，要进行合并计算
2。当遇到字母时，尤其是连续字母时，要进行合并计算
3。当遇到左中括号时，要将当前字母和数字入栈，维护数字栈和字母栈
4。同时维护临时存放连续字母和连续数字的变量
5。当遇到右中括号时，要将当前临时和出栈的数字和字母合并计算
2[2[a3[c]2[bc]]]
2 ""
2 ""
临时a
临时3
3 "a"
临时c
计算 "a"+3*"c"
临时accc
临时2
2 "accc"
临时bc
计算 "accc"+2*"bc"
...

*/

func decodeString(s string) string {
	st := newDecodeStack()
	ans := ""
	multi := 0
	for _, c := range s {
		if i, err := strconv.Atoi(string(c)); err == nil {
			multi = multi*10 + i
		} else if string(c) == "[" {
			st.stack = append(st.stack, ans)
			st.num = append(st.num, multi)
			ans = ""
			multi = 0
		} else if string(c) == "]" {
			ansTmp := st.pop()
			tmp := st.popNum()
			ans = ansTmp + strings.Repeat(ans, tmp)
		} else {
			ans = ans + string(c)
		}
	}
	return ans
}

type decodeStack struct {
	stack []string
	num   []int
}

func newDecodeStack() *decodeStack {
	return &decodeStack{}
}

func (s *decodeStack) push(item string) {
	s.stack = append(s.stack, item)

}

func (s *decodeStack) pushNum(item int) {
	s.num = append(s.num, item)

}

func (s *decodeStack) pop() string {
	tmp := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return tmp
}

func (s *decodeStack) popNum() int {
	tmp := s.num[len(s.num)-1]
	s.num = s.num[:len(s.num)-1]
	return tmp
}
