package _394_Decode_String

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	a := "10[leetcode]"
	v := decodeString(a)
	fmt.Println(v)
}
