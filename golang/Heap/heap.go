package main

import (
	"fmt"
)

// MaxHeap 最大堆实现
type MaxHeap[E any] struct {
	data      []E
	size      int
	lessFunc  func(E, E) bool
}

// NewMaxHeap 创建一个新的最大堆
// lessFunc 用于比较两个元素，返回 true 表示 a < b
func NewMaxHeap[E any](lessFunc func(E, E) bool, capacity int) *MaxHeap[E] {
	return &MaxHeap[E]{
		data:     make([]E, capacity),
		size:     0,
		lessFunc: lessFunc,
	}
}

// NewMaxHeapDefault 创建默认容量的最大堆
func NewMaxHeapDefault[E any](lessFunc func(E, E) bool) *MaxHeap[E] {
	return NewMaxHeap[E](lessFunc, 10)
}

func (h *MaxHeap[E]) Size() int {
	return h.size
}

func (h *MaxHeap[E]) IsEmpty() bool {
	return h.size == 0
}

func (h *MaxHeap[E]) Add(e E) {
	if h.size == len(h.data) {
		h.resize(2 * len(h.data))
	}
	h.data[h.size] = e
	h.siftUp(h.size)
	h.size++
}

func (h *MaxHeap[E]) Max() E {
	if h.IsEmpty() {
		panic("Heap is empty")
	}
	return h.data[0]
}

func (h *MaxHeap[E]) ExtractMax() E {
	if h.IsEmpty() {
		panic("Heap is empty")
	}

	ret := h.data[0]
	h.swap(0, h.size-1)
	h.size--
	h.siftDown(0)

	if h.size <= len(h.data)/4 && len(h.data)/2 != 0 {
		h.resize(len(h.data) / 2)
	}

	return ret
}

func (h *MaxHeap[E]) Replace(e E) E {
	if h.IsEmpty() {
		panic("Heap is empty")
	}

	ret := h.data[0]
	h.data[0] = e
	h.siftDown(0)
	return ret
}

// siftUp 元素上浮
func (h *MaxHeap[E]) siftUp(index int) {
	for index > 0 && h.lessFunc(h.data[h.parent(index)], h.data[index]) {
		h.swap(h.parent(index), index)
		index = h.parent(index)
	}
}

// siftDown 元素下沉
func (h *MaxHeap[E]) siftDown(index int) {
	for h.left(index) < h.size {
		j := h.left(index)
		// 找到左右子节点中较大的那个
		if j+1 < h.size && h.lessFunc(h.data[j], h.data[j+1]) {
			j = h.right(index)
		}

		// 如果父节点已经大于等于较大的子节点，停止
		if h.lessFunc(h.data[j], h.data[index]) {
			break
		}

		h.swap(index, j)
		index = j
	}
}

func (h *MaxHeap[E]) parent(index int) int {
	if index <= 0 {
		panic("Index has no parent")
	}
	return (index - 1) / 2
}

func (h *MaxHeap[E]) left(index int) int {
	return index*2 + 1
}

func (h *MaxHeap[E]) right(index int) int {
	return index*2 + 2
}

func (h *MaxHeap[E]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap[E]) resize(newCapacity int) {
	newData := make([]E, newCapacity)
	for i := 0; i < h.size; i++ {
		newData[i] = h.data[i]
	}
	h.data = newData
}

func (h *MaxHeap[E]) ToSlice() []E {
	result := make([]E, h.size)
	for i := 0; i < h.size; i++ {
		result[i] = h.data[i]
	}
	return result
}

func (h *MaxHeap[E]) String() string {
	return fmt.Sprintf("MaxHeap%v", h.ToSlice())
}

// PriorityQueue 基于最大堆的优先队列
type PriorityQueue[E any] struct {
	heap *MaxHeap[E]
}

// NewPriorityQueue 创建一个新的优先队列
func NewPriorityQueue[E any](lessFunc func(E, E) bool, capacity int) *PriorityQueue[E] {
	return &PriorityQueue[E]{
		heap: NewMaxHeap[E](lessFunc, capacity),
	}
}

// NewPriorityQueueDefault 创建默认容量的优先队列
func NewPriorityQueueDefault[E any](lessFunc func(E, E) bool) *PriorityQueue[E] {
	return NewPriorityQueue[E](lessFunc, 10)
}

func (pq *PriorityQueue[E]) Enqueue(e E) {
	pq.heap.Add(e)
}

func (pq *PriorityQueue[E]) Dequeue() E {
	return pq.heap.ExtractMax()
}

func (pq *PriorityQueue[E]) Front() E {
	return pq.heap.Max()
}

func (pq *PriorityQueue[E]) Size() int {
	return pq.heap.Size()
}

func (pq *PriorityQueue[E]) IsEmpty() bool {
	return pq.heap.IsEmpty()
}

func (pq *PriorityQueue[E]) String() string {
	return fmt.Sprintf("PriorityQueue%v", pq.heap.ToSlice())
}

func main() {
	fmt.Println("=== Go语言堆和优先队列示例 ===")

	// 创建用于整数的比较函数
	intLess := func(a, b int) bool {
		return a < b
	}

	// 测试最大堆
	fmt.Println("\n--- 最大堆 ---")
	heap := NewMaxHeapDefault[intLess](intLess)
	for i := 0; i < 10; i++ {
		heap.Add(i)
	}
	fmt.Printf("堆: %s\n", heap.String())
	fmt.Printf("最大值: %d\n", heap.Max())

	fmt.Println("\n取出5个最大值:")
	for i := 0; i < 5; i++ {
		fmt.Printf("ExtractMax: %d\n", heap.ExtractMax())
	}
	fmt.Printf("剩余: %s\n", heap.String())

	// 测试优先队列
	fmt.Println("\n--- 优先队列 ---")
	pq := NewPriorityQueueDefault[intLess](intLess)
	pq.Enqueue(3)
	pq.Enqueue(1)
	pq.Enqueue(4)
	pq.Enqueue(2)
	fmt.Printf("优先队列: %s\n", pq.String())

	fmt.Println("按优先级出队:")
	for !pq.IsEmpty() {
		fmt.Printf("Dequeue: %d\n", pq.Dequeue())
	}

	fmt.Println("\n=== 示例程序结束 ===")
}
