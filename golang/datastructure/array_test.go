package datastructure

import (
	"fmt"
	"testing"
)

func TestArrayAdd(t *testing.T) {
	array := Array(2)
	fmt.Println(array)
	array.Add(0, 10)
	array.Add(1, "tetetet")
	array.ToString()
}


func TestAddHead(t *testing.T) {
	array := Array(10)
	array.AddHead("First")
	array.AddHead("Second")
	array.AddTail("Third")
	array.ToString()
}


func TestRemove(t *testing.T) {
	array := Array(10)
	array.AddHead("First")
	array.AddHead("Second")
	array.AddTail("Third")
	fmt.Println(array.Size())
	array.RemoveLast()
	array.RemoveFirst()
	fmt.Println(array.Size())
	array.ToString()
}

func TestType(t *testing.T) {
	s1 := []int { 1,2 ,3,5}

	fmt.Println(s1[2:3])


	fmt.Println(struct {}{} == struct{}{})
}
