package main

import (
	"fmt"
	"testing"
)

func TestCompute(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	c := evalRPN(tokens)
	fmt.Println(c)

	a := "3[a]"
	for _, b := range a {
		fmt.Print(string(b), "\n")
	}
}
