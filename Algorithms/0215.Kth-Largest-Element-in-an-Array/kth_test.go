package _215_Kth_Largest_Element_in_an_Array

import (
	"fmt"
	"testing"
)

func TestKth(t *testing.T) {
	a := []int{3,2,1,5,6,4}
	k := 2

	result := findKthLargest(a, k)
	fmt.Println(result)
}
