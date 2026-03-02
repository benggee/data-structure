package main

import (
	"fmt"
	"strings"
)

// Queue 队列接口
type Queue[E any] interface {
	Size() int
	IsEmpty() bool
	Enqueue(e E)
	Dequeue() E
	Front() E
}

// ArrayQueue 基于动态数组的队列实现
type ArrayQueue[E any] struct {
	array *Array[E]
}

// NewArrayQueue 创建一个新的数组队列
func NewArrayQueue[E any](capacity int) *ArrayQueue[E] {
	return &ArrayQueue[E]{
		array: NewArray[E](capacity),
	}
}

// NewArrayQueueDefault 创建默认容量的数组队列
func NewArrayQueueDefault[E any]() *ArrayQueue[E] {
	return &ArrayQueue[E]{
		array: NewArrayDefault[E](),
	}
}

func (q *ArrayQueue[E]) Size() int {
	return q.array.Size()
}

func (q *ArrayQueue[E]) IsEmpty() bool {
	return q.array.IsEmpty()
}

func (q *ArrayQueue[E]) Enqueue(e E) {
	q.array.AddLast(e)
}

func (q *ArrayQueue[E]) Dequeue() E {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	return q.array.RemoveFirst()
}

func (q *ArrayQueue[E]) Front() E {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	return q.array.Get(0)
}

func (q *ArrayQueue[E]) String() string {
	var builder strings.Builder
	builder.WriteString("front:[")
	for i := 0; i < q.array.Size(); i++ {
		builder.WriteString(fmt.Sprintf("%v", q.array.Get(i)))
		if i < q.array.Size()-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]tail")
	return builder.String()
}

// LoopQueue 循环队列实现
type LoopQueue[E any] struct {
	data  []E
	front int
	tail  int
	size  int
}

// NewLoopQueue 创建一个新的循环队列
func NewLoopQueue[E any](capacity int) *LoopQueue[E] {
	return &LoopQueue[E]{
		data:  make([]E, capacity+1), // +1 for distinguishing full/empty
		front: 0,
		tail:  0,
		size:  0,
	}
}

// NewLoopQueueDefault 创建默认容量的循环队列
func NewLoopQueueDefault[E any]() *LoopQueue[E] {
	return NewLoopQueue[E](10)
}

func (q *LoopQueue[E]) Capacity() int {
	return len(q.data) - 1
}

func (q *LoopQueue[E]) Size() int {
	return q.size
}

func (q *LoopQueue[E]) IsEmpty() bool {
	return q.front == q.tail
}

func (q *LoopQueue[E]) Enqueue(e E) {
	// Check if queue is full
	if (q.tail+1)%len(q.data) == q.front {
		q.resize(q.Capacity() * 2)
	}

	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

func (q *LoopQueue[E]) Dequeue() E {
	if q.IsEmpty() {
		panic("Queue is empty")
	}

	ret := q.data[q.front]
	q.data[q.front] = *new(E) // Reset to zero value
	q.front = (q.front + 1) % len(q.data)
	q.size--

	// Shrink if needed
	if q.size <= q.Capacity()/4 && q.Capacity()/2 != 0 {
		q.resize(q.Capacity() / 2)
	}

	return ret
}

func (q *LoopQueue[E]) Front() E {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	return q.data[q.front]
}

func (q *LoopQueue[E]) resize(newCapacity int) {
	newData := make([]E, newCapacity+1)
	for i := 0; i < q.size; i++ {
		newData[i] = q.data[(q.front+i)%len(q.data)]
	}
	q.data = newData
	q.front = 0
	q.tail = q.size
}

func (q *LoopQueue[E]) String() string {
	var builder strings.Builder
	builder.WriteString("front:[")

	if q.tail >= q.front {
		for i := q.front; i < q.tail; i++ {
			builder.WriteString(fmt.Sprintf("%v", q.data[i]))
			if i < q.tail-1 {
				builder.WriteString(", ")
			}
		}
	} else {
		for i := q.front; i < len(q.data); i++ {
			builder.WriteString(fmt.Sprintf("%v", q.data[i]))
			builder.WriteString(", ")
		}
		for i := 0; i < q.tail; i++ {
			builder.WriteString(fmt.Sprintf("%v", q.data[i]))
			if i < q.tail-1 {
				builder.WriteString(", ")
			}
		}
	}

	builder.WriteString("]tail")
	return builder.String()
}

func main() {
	fmt.Println("=== Go语言队列示例 ===")

	// 测试数组队列
	fmt.Println("\n--- 数组队列 ---")
	aq := NewArrayQueueDefault[int]()
	for i := 0; i < 5; i++ {
		aq.Enqueue(i)
		fmt.Printf("入队 %d: %s\n", i, aq.String())
	}

	fmt.Printf("队首元素: %d\n", aq.Front())
	fmt.Printf("出队: %d\n", aq.Dequeue())
	fmt.Printf("出队后: %s\n", aq.String())

	// 测试循环队列
	fmt.Println("\n--- 循环队列 ---")
	lq := NewLoopQueueDefault[int]()
	for i := 0; i < 10; i++ {
		lq.Enqueue(i)
	}
	fmt.Printf("入队10个元素后: %s\n", lq.String())
	fmt.Printf("容量: %d, 大小: %d\n", lq.Capacity(), lq.Size())

	for i := 0; i < 5; i++ {
		lq.Dequeue()
	}
	fmt.Printf("出队5个元素后: %s\n", lq.String())
	fmt.Printf("容量: %d, 大小: %d\n", lq.Capacity(), lq.Size())

	fmt.Println("\n=== 示例程序结束 ===")
}
