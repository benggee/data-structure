package tool

import (
	"math/rand"
	"time"
)

// 随机生成无序数组
func GenarageDisorderInt64(len int) []int64 {
	rand.Seed(time.Now().UnixNano())

	tmpData := []int64{}
	for i := 0; i < len; i++ {
		tmpData = append(tmpData, rand.Int63n(100000))
	}
	return tmpData
}

// 随机生成近乎有序的数组
func GenarageOrderlyInt64(len int) []int64 {
	tmpArr := []int64{}
	for i := 0; i < len; i++ {
		tmpArr = append(tmpArr, int64(i + 100))
	}

	// 随机进行100次交换，让数组变成无序的
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100 ; i++ {
		randIndex := rand.Intn(len-1)
		randIndex2 := rand.Intn(len-1)
		tmp := tmpArr[randIndex]
		tmpArr[randIndex] = tmpArr[randIndex2]
		tmpArr[randIndex2] = tmp
	}
	return tmpArr
}