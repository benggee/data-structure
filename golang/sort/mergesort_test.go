package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	testData := []int64{}
	dataLen := 50000000

	for i := 0; i < dataLen; i++ {
		testData = append(testData, rand.Int63n(100000))
	}

	startTime := time.Now().UnixNano() / 1e6
	fmt.Println("START TIME:", startTime)
	ms := NewMerge()
	ms.Sort(testData)
	endTime := time.Now().UnixNano()/1e6 - startTime
	fmt.Println("END TIME:", endTime)

	// 验证是否正确
	for i := 0; i < dataLen-1; i++ {
		if testData[i] > testData[i+1] {
			t.Fail()
		}
	}
}

func TestDownUpMergeSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	testData := []int64{}
	dataLen := 500000000

	for i := 0; i < dataLen; i++ {
		testData = append(testData, rand.Int63n(1000000000))
	}

	startTime := time.Now().UnixNano() / 1e6
	fmt.Println("START TIME:", startTime)
	ms := NewMerge()
	ms.DownUpSort(testData)
	endTime := time.Now().UnixNano()/1e6 - startTime
	fmt.Println("END TIME:", endTime)

	// 验证是否正确
	for i := 0; i < dataLen-1; i++ {
		if testData[i] > testData[i+1] {
			t.Fail()
		}
	}
}
