package main

import (
	"fmt"
	"math/rand"
	"time"
)

// LinkListStack 基于链表实现的栈
// 栈是一种后进先出（LIFO）的数据结构
type LinkListStack[E any] struct {
	list *LinkList[E] // 使用链表作为底层存储
}

// NewLinkListStack 创建一个新的链表栈
func NewLinkListStack[E any]() *LinkListStack[E] {
	return &LinkListStack[E]{
		list: NewLinkList[E](),
	}
}

// GetSize 返回栈的大小
func (s *LinkListStack[E]) GetSize() int {
	return s.list.GetSize()
}

// IsEmpty 判断栈是否为空
func (s *LinkListStack[E]) IsEmpty() bool {
	return s.list.IsEmpty()
}

// Push 将元素压入栈顶
func (s *LinkListStack[E]) Push(e E) {
	s.list.AddFirst(e)
}

// Pop 弹出栈顶元素
func (s *LinkListStack[E]) Pop() E {
	return s.list.DelFirst()
}

// Peek 查看栈顶元素（不弹出）
func (s *LinkListStack[E]) Peek() E {
	return s.list.GetFirst()
}

// String 返回栈的字符串表示
func (s *LinkListStack[E]) String() string {
	var res string
	res = "Stack: top " + s.list.String()
	return res
}

// Clear 清空栈
func (s *LinkListStack[E]) Clear() {
	s.list.Clear()
}

// ToSlice 将栈转换为切片（从栈顶到栈底）
func (s *LinkListStack[E]) ToSlice() []E {
	return s.list.ToSlice()
}

// TestStack 测试栈的性能
// opCount: 操作次数
// 返回执行时间（秒）
func TestStack[E any](stack *LinkListStack[E], opCount int, gen func() E) float64 {
	startTime := time.Now()

	// 执行压栈操作
	for i := 0; i < opCount; i++ {
		stack.Push(gen())
	}

	// 执行出栈操作
	for i := 0; i < opCount; i++ {
		stack.Pop()
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	return duration.Seconds()
}

// TestStackOperations 测试栈的基本操作
func TestStackOperations() {
	fmt.Println("=== 测试链表栈基本操作 ===")

	stack := NewLinkListStack[int]()

	// 测试压栈操作
	fmt.Println("测试压栈操作...")
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Printf("压入 %d, 栈: %s\n", i, stack.String())
	}

	// 测试查看栈顶元素
	fmt.Printf("栈顶元素: %d\n", stack.Peek())

	// 测试出栈操作
	fmt.Println("测试出栈操作...")
	for i := 0; i < 3; i++ {
		popped := stack.Pop()
		fmt.Printf("弹出 %d, 栈: %s\n", popped, stack.String())
	}

	// 测试栈的大小
	fmt.Printf("栈的大小: %d\n", stack.GetSize())
	fmt.Printf("栈是否为空: %t\n", stack.IsEmpty())
}

// TestStackPerformance 测试栈的性能
func TestStackPerformance() {
	fmt.Println("\n=== 测试链表栈性能 ===")

	stack := NewLinkListStack[int]()
	opCount := 1000000
	rand.Seed(time.Now().UnixNano())
	time := TestStack[int](stack, opCount, func() int { return rand.Int() })
	fmt.Printf("执行 %d 次操作耗时: %.6f 秒\n", opCount*2, time)
}

// CompareStackPerformance 比较不同栈实现的性能
func CompareStackPerformance() {
	fmt.Println("\n=== 比较栈实现性能 ===")

	// 测试链表栈
	linkStack := NewLinkListStack[int]()
	rand.Seed(time.Now().UnixNano())
	linkTime := TestStack[int](linkStack, 100000, func() int { return rand.Int() })
	fmt.Printf("链表栈耗时: %.6f 秒\n", linkTime)

	// 这里可以添加数组栈的性能测试
	// arrayStack := NewArrayStack[int]()
	// arrayTime := TestArrayStack(arrayStack, 100000)
	// fmt.Printf("数组栈耗时: %.6f 秒\n", arrayTime)
}
