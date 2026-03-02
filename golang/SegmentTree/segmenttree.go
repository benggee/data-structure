package main

import (
	"fmt"
)

// Merger 合并两个值的函数类型
type Merger[T any] func(a, b T) T

// SegmentTree 线段树实现
type SegmentTree[T any] struct {
	tree    []T
	data    []T
	merger  Merger[T]
}

// NewSegmentTree 创建一个新的线段树
func NewSegmentTree[T any](arr []T, merger Merger[T]) *SegmentTree[T] {
	tree := make([]T, 4*len(arr))
	data := make([]T, len(arr))
	copy(data, arr)

	st := &SegmentTree[T]{
		tree:   tree,
		data:   data,
		merger: merger,
	}

	if len(arr) > 0 {
		st.buildSegmentTree(0, 0, len(arr)-1)
	}

	return st
}

// buildSegmentTree 构建线段树
func (st *SegmentTree[T]) buildSegmentTree(treeIndex, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}

	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)

	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// Query 查询区间 [queryL, queryR]
func (st *SegmentTree[T]) Query(queryL, queryR int) T {
	if queryL < 0 || queryL >= len(st.data) ||
		queryR < 0 || queryR >= len(st.data) || queryL > queryR {
		panic("Query range is invalid")
	}

	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// query 递归查询
func (st *SegmentTree[T]) query(treeIndex, l, r, queryL, queryR int) T {
	if l == queryL && r == queryR {
		return st.tree[treeIndex]
	}

	mid := l + (r-l)/2
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	if queryL > mid {
		// 查询区间完全在右子树
		return st.query(rightTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		// 查询区间完全在左子树
		return st.query(leftTreeIndex, l, mid, queryL, queryR)
	} else {
		// 查询区间跨越左右子树
		leftResult := st.query(leftTreeIndex, l, mid, queryL, mid)
		rightResult := st.query(rightTreeIndex, mid+1, r, mid+1, queryR)
		return st.merger(leftResult, rightResult)
	}
}

// Set 设置索引位置的值
func (st *SegmentTree[T]) Set(index int, value T) {
	if index < 0 || index >= len(st.data) {
		panic("Index out of range")
	}

	st.data[index] = value
	st.set(0, 0, len(st.data)-1, index, value)
}

// set 递归设置值
func (st *SegmentTree[T]) set(treeIndex, l, r, index int, value T) {
	if l == r {
		st.tree[treeIndex] = value
		return
	}

	mid := l + (r-l) / 2
	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	if index <= mid {
		st.set(leftTreeIndex, l, mid, index, value)
	} else {
		st.set(rightTreeIndex, mid+1, r, index, value)
	}

	// 重新计算父节点的值
	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

// GetSize 获取数据大小
func (st *SegmentTree[T]) GetSize() int {
	return len(st.data)
}

// Get 获取索引位置的值
func (st *SegmentTree[T]) Get(index int) T {
	if index < 0 || index >= len(st.data) {
		panic("Index out of range")
	}
	return st.data[index]
}

func (st *SegmentTree[T]) leftChild(index int) int {
	return 2*index + 1
}

func (st *SegmentTree[T]) rightChild(index int) int {
	return 2*index + 2
}

func (st *SegmentTree[T]) String() string {
	return fmt.Sprintf("SegmentTree(size=%d)", len(st.data))
}

func main() {
	fmt.Println("=== Go语言线段树示例 ===")

	// 测试求和线段树
	fmt.Println("\n--- 求和线段树 ---")
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	sumMerger := func(a, b int) int {
		return a + b
	}

	segTree := NewSegmentTree(arr, sumMerger)
	fmt.Printf("原始数组: %v\n", arr)

	fmt.Printf("Sum[0, 3]: %d\n", segTree.Query(0, 3))
	fmt.Printf("Sum[2, 5]: %d\n", segTree.Query(2, 5))
	fmt.Printf("Sum[0, 7]: %d\n", segTree.Query(0, 7))

	segTree.Set(3, 10)
	fmt.Printf("\n设置索引3为10后:\n")
	fmt.Printf("Sum[0, 3]: %d\n", segTree.Query(0, 3))

	// 测试最大值线段树
	fmt.Println("\n--- 最大值线段树 ---")
	arr2 := []int{1, 3, 5, 7, 9, 11, 13, 15}
	maxMerger := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxSegTree := NewSegmentTree(arr2, maxMerger)
	fmt.Printf("原始数组: %v\n", arr2)

	fmt.Printf("Max[0, 7]: %d\n", maxSegTree.Query(0, 7))
	fmt.Printf("Max[2, 5]: %d\n", maxSegTree.Query(2, 5))

	maxSegTree.Set(3, 20)
	fmt.Printf("\n设置索引3为20后:\n")
	fmt.Printf("Max[0, 7]: %d\n", maxSegTree.Query(0, 7))

	// 测试最小值线段树
	fmt.Println("\n--- 最小值线段树 ---")
	arr3 := []int{5, 3, 8, 1, 9, 2, 7, 4}
	minMerger := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	minSegTree := NewSegmentTree(arr3, minMerger)
	fmt.Printf("原始数组: %v\n", arr3)

	fmt.Printf("Min[0, 7]: %d\n", minSegTree.Query(0, 7))
	fmt.Printf("Min[2, 5]: %d\n", minSegTree.Query(2, 5))

	fmt.Println("\n=== 示例程序结束 ===")
}
