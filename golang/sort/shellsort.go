package sort

func ShellSort(data []int64) {
	dataLen := len(data)
	var d int = dataLen / 2  // 这里简单除以2得到增量
	for d >= 1 {
		// 这一层循环是处理每一组数据 假如数组长度为20，那么第一次是10，第二次是5，第三次是2
		for i := 0; i < d; i++ {
			// 这里是从第一个元素跳过d的长度遍历完同一组的数据
			for j := i+d; j < dataLen; j = j+d {
				tmp := data[j] 
				// 对每一组进行插入排序
				var x int  
				for x = j-d; x >= 0 && data[x] > tmp; x = x-d {
					data[x+d] = data[x]
				}
				data[x+d] = tmp 
			}
		}
		d = d / 2   // 增量除以2
	}
}