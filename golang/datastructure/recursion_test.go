package datastructure

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	tmpArr := []int{}

	for i := 0; i < 10; i++ {
		tmpArr = append(tmpArr, i)
	}

	r := Recursion()
	fmt.Println(r.Sum(tmpArr))
}