package sort

import (
	"data-structure/sort/tool"
	"math/rand"
)

// 快速排序
// 快排的原理是找到一个标点，一般使用数组最左边的元素，假定为P
// 然后指定一个中间flag的下标，假定为j，接着就可以开始扫描元素了
// 如果值>p则继续往下遍历，什么都不做
// 如果值<p则交换当前元素与j所在的元素值，接着j++, 这样一轮扫描完了之后，j左边都是小于p的，右边都是大于p的值
// 最后将l的值和j所在的值进行交换
// 我们把上面的过程当作一个整体的子过程，进行递归调用
func QuickSort(arr []int64) {
	_quickSort(arr, 0, len(arr)-1)
}

// 双路快排
// 双路快排实际上是对上面简化版的优化，假如数据中有大量相同的元素，j的左边或者右边包含的元素会非常不平均
// 双路快排实际上是分成两路进行操作
// 第一路，是从前往后，如果小于等于标点值就什么都不做，继续往后扫描，直到当前元素大于标点值
// 第二路，是从后往前，如果大于等于标点值就什么都不做，继续往前扫描，直到当元素小于标点值
// 两路遍历完，如果第一路的下标大于第二路则结束扫描，否则交换i和j所在位置的值
// 最后交换最左边和第一路最后下标的值
func QuickSort2Way(arr []int64) {
	_quickSort2Way(arr, 0, len(arr)-1)
}

// 三路快排
func QuickSort3Way(arr []int64) {
	_quickSort3Way(arr, 0, len(arr)-1)
}

func _quickSort(arr []int64, l int, r int) {
	if l >= r {
		return
	}

	p := partition(arr, l, r)
	_quickSort(arr, l, p-1)
	_quickSort(arr, p+1, r)
}

func partition(arr []int64, l int, r int) int {
	// 优化，随机选择p,防止数据在接近有序时退化成n^2的时间复杂度
	// rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(r-l) + l
	// 进行一次交换
	tool.ArrSwap(arr, randIndex, l)

	// 以最左右的位置为基准，记录下最左边的值
	p := arr[l]
	j := l  // 表示的是<p和>p的中间位置，起始从最左边开始
	for i := l + 1; i <= r; i++ {
		// 如果当前的位置<p则和j+1的位置进行交换，同时j++，那时j的位置就是交换过后的当前值
		if arr[i] < p {
			tool.ArrSwap(arr, i, j+1)
			j++
		}
	}

	// 最后，要将基准位置和j的位置进行交换，此时，j的位置左边是小于p，右边是大于p的元素了
	// 直到最后一次递归调用，整个数组就是有序的了
	tool.ArrSwap(arr, l, j)

	return j
}

func _quickSort2Way(arr []int64, l int, r int) {
	if l >= r {
		return
	}

	p := _partition(arr, l, r)
	_quickSort2Way(arr, l, p - 1)
	_quickSort2Way(arr, p + 1, r)
}

func _partition(arr []int64, l int, r int) int {
	// 随机取一个标定点，为了避免数据在接近有序的情况下退化成链表
	randIndex := rand.Intn(r-l) + l
	tool.ArrSwap(arr, randIndex, l)

	// 标定点元素的值
	v := arr[l]

	i := l + 1  // l所在位置的值是当前标点，跳过
	j := r
	for  {
		// 如果当前i所在的元素小于标点元素，则继续往后扫描
		for i <= r && arr[i] < v {
			i++
		}
		// 如果当前j所在元素小于标点元素，则继续向前扫摸黑
		for j >= l + 1 && arr[j] > v {
			j--
		}
		// 如果i 已经大于j了直接break
		if i > j {
			break
		}

		// 到此，交换j和i位置的值
		tool.ArrSwap(arr, i, j)
		i++
		j--
	}
	// 最后要交换l和j所在位置的值
	tool.ArrSwap(arr, l, j)

	// 将j返回用于下次递归
	return j
}

func _quickSort3Way(arr []int64, l int, r int) {
	// 递归结束
	if l >= r {
		return
	}

	// 随机取一个标定点，为了避免数据在接近有序的情况下退化成链表
	randIndex := rand.Intn(r-l) + l
	tool.ArrSwap(arr, randIndex, l)

	// 标定点元素的值
	v := arr[l]

	lt := l  // arr[l+1 ... lt] < v
	gt := r + 1 // arr[gt ...r] > v
	i := l + 1
	for i < gt {
		if arr[i] < v {
			tool.ArrSwap(arr, i, lt+1) // lt+1 是已排好的，比v小的区间最后一个元素的下一个元素
			i++
			lt++
		} else if arr[i] > v {
			tool.ArrSwap(arr, i, gt-1) // 这里的gt-1 是已排好的，比v大的区间最前面一个元素的前一个元素， 此时i不需要维护
			gt--
		} else { // arr[i] == v 如果相等，什么都不用做，继续往后扫描
			i++
		}
	}
	// 最后将标点l位置的值进行交换
	tool.ArrSwap(arr, l, lt)

	_quickSort3Way(arr, l, lt-1)
	_quickSort3Way(arr, gt, r)
}