package bsearch

import "testing"

func TestBSearch7(t *testing.T) {
	a := []int{3, 5, 6, 8, 9, 10}
	value := bsearch7(a, 7)
	if value != 6 {
		t.Error("bsearch error")
	}
}
