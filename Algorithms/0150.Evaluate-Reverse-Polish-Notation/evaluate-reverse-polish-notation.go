package main

import "strconv"

var m = map[string]struct{}{"+": {}, "-": {}, "*": {}, "/": {}}

func evalRPN(tokens []string) int {
	s := newStack()
	for _, token := range tokens {
		if _, ok := m[token]; ok {
			v := compute(s.pop(), s.pop(), token)
			s.push(v)
			continue
		}
		t, _ := strconv.Atoi(token)
		s.push(t)
	}
	return s.pop()
}

func compute(x, y int, flag string) int {
	switch flag {
	case "+":
		return x + y
	case "-":
		return y - x
	case "*":
		return x * y
	case "/":
		return y / x
	}
	return 0
}

type stack struct {
	items []int
}

func newStack() *stack {
	return &stack{make([]int, 0)}
}

func (s *stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *stack) pop() int {
	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return tmp
}
