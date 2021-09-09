package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	a := []int{4, 5, 6, 3, 2, 1}
	n := 6
	a = SelectSort(a, n)
	fmt.Println(a)
}
