package sort

// 冒泡排序
func BubbleSort(data []int64) {
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		for j := 0; j < dataLen-i-1; j++ {
			if data[j] > data[j+1] {
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
	}
}