package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{4,5,6,1,2,3}
	result := MergeSort(a)
	fmt.Println(result)
}
