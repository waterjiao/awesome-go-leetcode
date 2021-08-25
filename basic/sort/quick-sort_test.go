package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{4, 5, 6, 1, 2, 3}
	result := QuickSort(a)
	fmt.Println(result)
}
