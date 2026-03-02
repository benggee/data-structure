package main

import (
	"fmt"
	"time"
)

// Main 主程序结构体
type Main struct{}

// TestLinkList 测试链表的基本功能
func (m *Main) TestLinkList() {
	fmt.Println("=== 测试链表基本功能 ===")

	// 创建链表
	list := NewLinkList[int]()

	// 测试添加元素
	fmt.Println("测试添加元素...")
	for i := 0; i < 6; i++ {
		list.AddFirst(i)
		fmt.Printf("添加 %d, 链表: %s\n", i, list.String())
	}

	// 测试在指定位置添加元素
	fmt.Println("测试在指定位置添加元素...")
	list.Add(3, 8888)
	fmt.Printf("在位置3添加8888, 链表: %s\n", list.String())

	// 测试查找元素
	fmt.Println("测试查找元素...")
	fmt.Printf("查找元素1: %t\n", list.Find(1))
	fmt.Printf("查找元素999: %t\n", list.Find(999))

	// 测试获取元素
	fmt.Printf("获取第一个元素: %d\n", list.GetFirst())
	fmt.Printf("获取最后一个元素: %d\n", list.GetLast())
	fmt.Printf("获取位置2的元素: %d\n", list.Get(2))

	// 测试删除元素
	fmt.Println("测试删除元素...")
	deleted := list.Del(3)
	fmt.Printf("删除位置3的元素 %d, 链表: %s\n", deleted, list.String())

	deleted = list.DelFirst()
	fmt.Printf("删除第一个元素 %d, 链表: %s\n", deleted, list.String())

	// 测试设置元素
	fmt.Println("测试设置元素...")
	list.Set(1, 100)
	fmt.Printf("设置位置1为100, 链表: %s\n", list.String())

	// 测试链表大小
	fmt.Printf("链表大小: %d\n", list.GetSize())
	fmt.Printf("链表是否为空: %t\n", list.IsEmpty())
}

// TestLinkListAdvanced 测试链表的高级功能
func (m *Main) TestLinkListAdvanced() {
	fmt.Println("\n=== 测试链表高级功能 ===")

	list := NewLinkList[int]()

	// 添加一些元素
	for i := 0; i < 5; i++ {
		list.AddLast(i)
	}
	fmt.Printf("原始链表: %s\n", list.String())

	// 测试反转链表
	fmt.Println("测试反转链表...")
	list.Reverse()
	fmt.Printf("反转后: %s\n", list.String())

	// 测试转换为切片
	fmt.Println("测试转换为切片...")
	slice := list.ToSlice()
	fmt.Printf("转换为切片: %v\n", slice)

	// 测试清空链表
	fmt.Println("测试清空链表...")
	list.Clear()
	fmt.Printf("清空后: %s\n", list.String())
	fmt.Printf("清空后大小: %d\n", list.GetSize())

	// 测试环检测
	fmt.Println("测试环检测...")
	fmt.Printf("是否有环: %t\n", list.HasCycle())

	// 创建一个有环的链表（仅用于测试）
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	// 注意：这里不实际创建环，因为会破坏链表结构
	fmt.Printf("正常链表是否有环: %t\n", list.HasCycle())
}

// TestStack 测试栈功能
func (m *Main) TestStack() {
	fmt.Println("\n=== 测试栈功能 ===")

	// 测试栈的基本操作
	TestStackOperations()

	// 测试栈的性能
	TestStackPerformance()

	// 测试栈的边界情况
	fmt.Println("\n=== 测试栈边界情况 ===")
	stack := NewLinkListStack[int]()

	fmt.Printf("空栈大小: %d\n", stack.GetSize())
	fmt.Printf("空栈是否为空: %t\n", stack.IsEmpty())

	// 测试单个元素
	stack.Push(1)
	fmt.Printf("压入1后: %s\n", stack.String())

	// 测试出栈后变为空
	stack.Pop()
	fmt.Printf("出栈后: %s\n", stack.String())
	fmt.Printf("出栈后是否为空: %t\n", stack.IsEmpty())
}

// TestQueue 测试队列功能
func (m *Main) TestQueue() {
	fmt.Println("\n=== 测试队列功能 ===")

	// 测试队列的基本操作
	TestQueueOperations()

	// 测试队列的性能
	TestQueuePerformance()

	// 测试队列的边界情况
	TestQueueEdgeCases()
}

// TestPerformanceComparison 测试性能对比
func (m *Main) TestPerformanceComparison() {
	fmt.Println("\n=== 性能对比测试 ===")

	// 测试链表栈性能
	linkStack := NewLinkListStack[int]()
	linkTime := TestStack(linkStack, 100000, func() int { return 0 })
	fmt.Printf("链表栈 100000 次操作耗时: %.6f 秒\n", linkTime)

	// 测试链表队列性能
	linkQueue := NewLinkListQueue[int]()
	opCount := 100000

	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		linkQueue.Enqueue(i)
	}
	for i := 0; i < opCount; i++ {
		linkQueue.Dequeue()
	}
	endTime := time.Now()
	queueTime := endTime.Sub(startTime).Seconds()

	fmt.Printf("链表队列 100000 次操作耗时: %.6f 秒\n", queueTime)
}

func main() {
	main := &Main{}

	// 测试链表基本功能
	main.TestLinkList()

	// 测试链表高级功能
	main.TestLinkListAdvanced()

	// 测试栈功能
	main.TestStack()

	// 测试队列功能
	main.TestQueue()

	// 测试性能对比
	main.TestPerformanceComparison()

	fmt.Println("\n=== 所有测试完成 ===")
}
