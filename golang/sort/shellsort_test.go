package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestShellSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	testData := []int64{}
	dataLen := 5000000

	for i := 0; i < dataLen; i++ {
		testData = append(testData, rand.Int63n(100000))
	}

	startTime := time.Now().UnixNano() / 1e6
	fmt.Println("START TIME:", startTime)
	ShellSort(testData)
	endTime := time.Now().UnixNano()/1e6 - startTime
	fmt.Println("END TIME:", endTime)

	// 验证是否正确
	for i := 0; i < dataLen-1; i++ {
		if testData[i] > testData[i+1] {
			t.Fail()
		}
	}
}