package sort

// 插入排序
func InsertionSort(data []int64) {
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			tmp := data[j]
			data[j] = data[j-1]
			data[j-1] = tmp
		}
	}
}

// 优化版
func InsertionSortv2(data []int64) {
	var tmpData int64
	var j int
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		tmpData = data[i]
		for j = i; j > 0 && data[j-1] > tmpData; j-- {
			data[j] = data[j-1]
		}
		data[j] = tmpData
	}
}