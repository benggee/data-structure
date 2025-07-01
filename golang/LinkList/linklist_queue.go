package main

import (
	"fmt"
	"strings"
	"time"
)

// LinkListQueue 基于链表实现的队列
// 队列是一种先进先出（FIFO）的数据结构
type LinkListQueue[E any] struct {
	head *Node[E] // 队列头部
	tail *Node[E] // 队列尾部
	size int      // 队列大小
}

// NewLinkListQueue 创建一个新的链表队列
func NewLinkListQueue[E any]() *LinkListQueue[E] {
	return &LinkListQueue[E]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Enqueue 将元素加入队列尾部
func (q *LinkListQueue[E]) Enqueue(e E) {
	if q.tail == nil {
		// 队列为空，创建第一个节点
		q.tail = newNode(e)
		q.head = q.tail
	} else {
		// 在尾部添加新节点
		q.tail.next = newNode(e)
		q.tail = q.tail.next
	}
	q.size++
}

// Dequeue 从队列头部取出元素
func (q *LinkListQueue[E]) Dequeue() E {
	if q.size <= 0 {
		panic("The Queue is empty")
	}

	tmpNode := q.head
	q.head = q.head.next

	// 注意：这里使用 tmpNode.next = nil 来断开节点连接
	// 这是因为，tmpNode是作用域仅在函数里，这样赋值只是表示tmpNode指向的位置是一个空地址
	// 使用tmpNode.next = nil 实际上是将原head指向的next那个地方清空，当函数结束时，tmpNode也将被销毁，达到释放内存的目的
	tmpNode.next = nil

	// 如果队列变为空，需要更新tail指针
	if q.head == nil {
		q.tail = nil
	}

	q.size--
	return tmpNode.e
}

// GetFront 获取队列头部元素（不取出）
func (q *LinkListQueue[E]) GetFront() E {
	if q.size <= 0 {
		panic("The queue is empty")
	}
	return q.head.e
}

// GetSize 返回队列的大小
func (q *LinkListQueue[E]) GetSize() int {
	return q.size
}

// IsEmpty 判断队列是否为空
func (q *LinkListQueue[E]) IsEmpty() bool {
	return q.size == 0
}

// String 返回队列的字符串表示
func (q *LinkListQueue[E]) String() string {
	var res strings.Builder
	res.WriteString("Queue: ")
	cur := q.head
	for cur != nil {
		res.WriteString(fmt.Sprintf("%v->", cur.e))
		cur = cur.next
	}
	res.WriteString("NULL")
	return res.String()
}

// Clear 清空队列
func (q *LinkListQueue[E]) Clear() {
	q.head = nil
	q.tail = nil
	q.size = 0
}

// ToSlice 将队列转换为切片
func (q *LinkListQueue[E]) ToSlice() []E {
	result := make([]E, q.size)
	cur := q.head
	for i := 0; i < q.size; i++ {
		result[i] = cur.e
		cur = cur.next
	}
	return result
}

// TestQueueOperations 测试队列的基本操作
func TestQueueOperations() {
	fmt.Println("=== 测试链表队列基本操作 ===")

	queue := NewLinkListQueue[int]()

	// 测试入队操作
	fmt.Println("测试入队操作...")
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Printf("入队 %d, 队列: %s\n", i, queue.String())
	}

	// 测试查看队首元素
	fmt.Printf("队首元素: %d\n", queue.GetFront())

	// 测试出队操作
	fmt.Println("测试出队操作...")
	for i := 0; i < 4; i++ {
		dequeued := queue.Dequeue()
		fmt.Printf("出队 %d, 队列: %s\n", dequeued, queue.String())
	}

	// 测试队列的大小
	fmt.Printf("队列的大小: %d\n", queue.GetSize())
	fmt.Printf("队列是否为空: %t\n", queue.IsEmpty())
}

// TestQueuePerformance 测试队列的性能
func TestQueuePerformance() {
	fmt.Println("\n=== 测试链表队列性能 ===")

	queue := NewLinkListQueue[int]()
	opCount := 100000

	startTime := time.Now()

	// 执行入队操作
	for i := 0; i < opCount; i++ {
		queue.Enqueue(i)
	}

	// 执行出队操作
	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("执行 %d 次操作耗时: %.6f 秒\n", opCount*2, duration.Seconds())
}

// TestQueueEdgeCases 测试队列的边界情况
func TestQueueEdgeCases() {
	fmt.Println("\n=== 测试队列边界情况 ===")

	queue := NewLinkListQueue[int]()

	// 测试空队列
	fmt.Printf("空队列大小: %d\n", queue.GetSize())
	fmt.Printf("空队列是否为空: %t\n", queue.IsEmpty())

	// 测试单个元素
	queue.Enqueue(1)
	fmt.Printf("单个元素队列: %s\n", queue.String())

	// 测试出队后变为空
	queue.Dequeue()
	fmt.Printf("出队后队列: %s\n", queue.String())
	fmt.Printf("出队后是否为空: %t\n", queue.IsEmpty())

	// 测试多个元素入队出队
	for i := 0; i < 5; i++ {
		queue.Enqueue(i)
	}
	fmt.Printf("入队5个元素后: %s\n", queue.String())

	for i := 0; i < 3; i++ {
		queue.Dequeue()
	}
	fmt.Printf("出队3个元素后: %s\n", queue.String())
}
