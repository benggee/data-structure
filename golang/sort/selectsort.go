package sort

func SelectSort(data []int64) {
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		for j := i + 1; j < dataLen; j++ {
			if data[j] < data[i] {
				// 交换位置
				tmp := data[i]
				data[i] = data[j]
				data[j] = tmp
			}
		}
	}
}
