package main

import (
	"fmt"
	"strings"
)

// Set 集合接口
type Set[E any] interface {
	Add(e E)
	Contains(e E) bool
	Remove(e E)
	GetSize() int
	IsEmpty() bool
}

// LinkListSet 基于链表的集合实现
type LinkListSet[E any] struct {
	list *LinkList[E]
}

// NewLinkListSet 创建一个新的链表集合
func NewLinkListSet[E any]() *LinkListSet[E] {
	return &LinkListSet[E]{
		list: NewLinkList[E](),
	}
}

func (s *LinkListSet[E]) GetSize() int {
	return s.list.GetSize()
}

func (s *LinkListSet[E]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *LinkListSet[E]) Add(e E) {
	// 只有当元素不存在时才添加
	if !s.list.Find(e) {
		s.list.AddFirst(e)
	}
}

func (s *LinkListSet[E]) Contains(e E) bool {
	return s.list.Find(e)
}

func (s *LinkListSet[E]) Remove(e E) {
	s.list.Remove(e)
}

func (s *LinkListSet[E]) String() string {
	var builder strings.Builder
	builder.WriteString("LinkListSet{")
	builder.WriteString(s.list.String())
	builder.WriteString("}")
	return builder.String()
}

func main() {
	fmt.Println("=== Go语言集合示例 ===")

	set := NewLinkListSet[int]()

	// 添加元素
	fmt.Println("--- 添加元素 ---")
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(1) // 重复元素，不会被添加
	fmt.Printf("集合: %s\n", set.String())
	fmt.Printf("大小: %d\n", set.GetSize())

	// 查找元素
	fmt.Println("\n--- 查找元素 ---")
	fmt.Printf("包含2: %t\n", set.Contains(2))
	fmt.Printf("包含5: %t\n", set.Contains(5))

	// 删除元素
	fmt.Println("\n--- 删除元素 ---")
	set.Remove(2)
	fmt.Printf("删除2后: %s\n", set.String())
	fmt.Printf("大小: %d\n", set.GetSize())

	// 添加字符串
	fmt.Println("\n--- 字符串集合 ---")
	strSet := NewLinkListSet[string]()
	strSet.Add("Hello")
	strSet.Add("World")
	strSet.Add("Go")
	fmt.Printf("字符串集合: %s\n", strSet.String())

	fmt.Println("\n=== 示例程序结束 ===")
}
