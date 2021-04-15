package sort

type Merge struct {
}

func NewMerge() *Merge {
	return &Merge{}
}

// 自顶向下排序
// 递归
func (m *Merge) Sort(data []int64) {
	m.mergeSort(data, 0, len(data)-1)
}

// 自底向上排序
// 循环
// 这种方式实际意义可以对链表进行排序
// TODO: sort for link
// TODO: 需花大力气研究的代码
func (m *Merge) DownUpSort(data []int64) {
	dataLen := len(data)
	for i := 1; i <= dataLen; i = i + i {
		for j := 0; j+i < dataLen; j = j + i*2 {
			r := j + i*2 - 1
			if r > dataLen-1 {
				r = dataLen - 1
			}
			m.doMerge(data, j, j+i-1, r)
		}
	}
}

func (m *Merge) mergeSort(data []int64, l int, r int) {
	// 这是没有使用插入排序优化的方式
	// if l >= r {
	// 	return
	// }

	// 在数量比较小的情况下，我们可以使用效率更高的插入排序来进一步优化
	if (r - l) < 20 {
		m.insertSortLR(data, l, r)
		return
	}
	var mid int = (l + r) / 2 // 注意，这里可能会出现l+r越界的问题
	// 左区间
	m.mergeSort(data, l, mid)
	// 右区间
	m.mergeSort(data, mid+1, r)
	// 如果data[mid]>data[mid+1]说明不是有序
	if data[mid] > data[mid+1] {
		m.doMerge(data, l, mid, r)
	}
}

func (m *Merge) doMerge(data []int64, l int, mid int, r int) {
	// 创建一个副本
	tmp := make([]int64, r-l+1)
	for i := l; i <= r; i++ {
		tmp[i-l] = data[i]
	}
	// 排序
	var i int = l
	var j int = mid + 1
	for k := l; k <= r; k++ {
		// 如果i越过了中间量，说明已经到了右区间，直接将右区间当前指向的元素赋值给当前数组的元素
		if i > mid {
			data[k] = tmp[j-l]
			j = j + 1
		} else if j > r {
			data[k] = tmp[i-l]
			i = i + 1
		} else if tmp[i-l] < tmp[j-l] {
			data[k] = tmp[i-l]
			i = i + 1
		} else {
			data[k] = tmp[j-l]
			j = j + 1
		}
	}
}

func (m *Merge) insertSortLR(data []int64, l int, r int) {
	dataLen := r - l
	for i := l; i < dataLen; i++ {
		tmp := data[i]
		var j int
		for j = i; j > 0 && data[j+1] > tmp; j++ {
			data[j] = data[j+1]
		}
		data[j] = tmp
	}
}
