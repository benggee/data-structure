package sort

import (
	"data-structure/sort/tool"
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := tool.GenarageDisorderInt64(500000)
	QuickSort(arr)

	if tool.VerifyOrder(arr) == false {
		t.Fail()
	}
}

func TestQuickSortByOrderly(t *testing.T) {
	arr := tool.GenarageOrderlyInt64(50000000)
	t1 := tool.EchoCurrenNaTime()
	QuickSort(arr)
	t2 := tool.EchoCurrenNaTime()
	fmt.Println("TIMES:", t2-t1, " ms ")
	if tool.VerifyOrder(arr) == false {
		t.Fail()
	}
}

// 双路快排
func TestQuickSort2Way(t *testing.T) {
	arr := tool.GenarageDisorderInt64(500000)
	t1 := tool.EchoCurrenNaTime()
	QuickSort2Way(arr)
	t2 := tool.EchoCurrenNaTime()
	fmt.Println("TIMES:", t2-t1, " ms ")
	if tool.VerifyOrder(arr) == false {
		t.Fail()
	}
}

// 三路快排
func TestQuickSort3Way(t *testing.T) {
	arr := tool.GenarageDisorderInt64(500000)
	t1 := tool.EchoCurrenNaTime()
	QuickSort3Way(arr)
	t2 := tool.EchoCurrenNaTime()
	fmt.Println("TIMES:", t2-t1, " ms ")
	if tool.VerifyOrder(arr) == false {
		t.Fail()
	}
}