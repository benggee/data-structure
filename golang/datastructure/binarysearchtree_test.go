package datastructure

import (
	"fmt"
	"testing"

	"data-structure/sort/tool"
)

func TestBinarySearchTreeAdd(t *testing.T) {
	tmpArr := tool.GenarageDisorderInt64(10)
	bt := BinarySearchTree()
	for _, v := range tmpArr {
		bt.Add(v, v)
	}

	bt.PreOrder()
}

func TestBinarySearchTreeOrder(t *testing.T) {
	testData := []int64{2,4,10,5,3,1}
	bt := BinarySearchTree()
	for _,v := range testData {
		bt.Add(v,v)
	}
	bt.PreOrder()
	fmt.Println("pre order end.")
	bt.InOrder()
	fmt.Println("in order end.")
	bt.PostOrder()
	fmt.Println("last order end.")
	bt.LevelOrder()
	fmt.Println("level order end.")
	bt.PreOrderNR()
	fmt.Println("pre order nr end.")
}

func TestBinarySearchTreeMaxMin(t *testing.T) {
	testData := []int64{2,4,10,5,3,1}
	bt := BinarySearchTree()
	for _,v := range testData {
		bt.Add(v,v)
	}
	fmt.Println("MAX-", bt.Max())
	fmt.Println("MIN-", bt.Min())
}

func TestBinarySearchTreeDelMaxAndMin(t *testing.T) {
	testData := []int64{2,4,10,5,3,1}
	bt := BinarySearchTree()
	for _,v := range testData {
		bt.Add(v,v)
	}
	bt.PreOrderNR()
	bt.DelMin()
	fmt.Println("after del min")
	bt.PreOrderNR()
	bt.DelMax()
	fmt.Println("after del max")
	bt.PreOrderNR()
}

func TestBinarySearchTreeDelAny(t *testing.T) {
	testData := []int64{2,4,10,5,3,1}
	bt := BinarySearchTree()
	for _,v := range testData {
		bt.Add(v,v)
	}
	bt.PreOrderNR()
	bt.Del(2)
	fmt.Println("after del 2")
	bt.PreOrderNR()
	bt.Del(4)
	fmt.Println("after del 4")
	bt.PreOrderNR()
}