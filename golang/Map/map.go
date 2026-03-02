package main

import (
	"fmt"
	"strings"
)

// Map 映射接口
type Map[K any, V any] interface {
	Remove(key K) V
	Contains(key K) bool
	Get(key K) V
	Set(key K, value V)
	Size() int
	IsEmpty() bool
}

// MapNode 映射节点
type MapNode[K any, V any] struct {
	key   K
	value V
	next  *MapNode[K, V]
}

// LinkListMap 基于链表的映射实现
type LinkListMap[K any, V any] struct {
	dummyHead *MapNode[K, V]
	size      int
}

// NewLinkListMap 创建一个新的链表映射
func NewLinkListMap[K any, V any]() *LinkListMap[K, V] {
	return &LinkListMap[K, V]{
		dummyHead: &MapNode[K, V]{},
		size:      0,
	}
}

// getNode 根据键查找节点
func (m *LinkListMap[K, V]) getNode(key K) *MapNode[K, V] {
	cur := m.dummyHead.next
	for cur != nil {
		if fmt.Sprintf("%v", cur.key) == fmt.Sprintf("%v", key) {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func (m *LinkListMap[K, V]) Size() int {
	return m.size
}

func (m *LinkListMap[K, V]) IsEmpty() bool {
	return m.size == 0
}

func (m *LinkListMap[K, V]) Contains(key K) bool {
	return m.getNode(key) != nil
}

func (m *LinkListMap[K, V]) Get(key K) V {
	var zero V
	node := m.getNode(key)
	if node == nil {
		return zero
	}
	return node.value
}

func (m *LinkListMap[K, V]) Set(key K, value V) {
	node := m.getNode(key)
	if node != nil {
		// 键已存在，更新值
		node.value = value
	} else {
		// 键不存在，添加新节点
		newNode := &MapNode[K, V]{
			key:   key,
			value: value,
			next:  m.dummyHead.next,
		}
		m.dummyHead.next = newNode
		m.size++
	}
}

func (m *LinkListMap[K, V]) Remove(key K) V {
	var zero V
	prev := m.dummyHead
	for prev.next != nil {
		if fmt.Sprintf("%v", prev.next.key) == fmt.Sprintf("%v", key) {
			delNode := prev.next
			prev.next = delNode.next
			delNode.next = nil
			m.size--
			return delNode.value
		}
		prev = prev.next
	}
	return zero
}

func (m *LinkListMap[K, V]) String() string {
	var builder strings.Builder
	builder.WriteString("LinkListMap{")
	cur := m.dummyHead.next
	for cur != nil {
		builder.WriteString(fmt.Sprintf("%v:%v", cur.key, cur.value))
		if cur.next != nil {
			builder.WriteString(", ")
		}
		cur = cur.next
	}
	builder.WriteString("}")
	return builder.String()
}

func main() {
	fmt.Println("=== Go语言映射示例 ===")

	m := NewLinkListMap[string, int]()

	// 添加键值对
	fmt.Println("--- 添加键值对 ---")
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)
	fmt.Printf("映射: %s\n", m.String())
	fmt.Printf("大小: %d\n", m.Size())

	// 获取值
	fmt.Println("\n--- 获取值 ---")
	fmt.Printf("Get('two'): %d\n", m.Get("two"))
	fmt.Printf("Get('four'): %d\n", m.Get("four"))

	// 检查键是否存在
	fmt.Println("\n--- 检查键 ---")
	fmt.Printf("Contains('one'): %t\n", m.Contains("one"))
	fmt.Printf("Contains('four'): %t\n", m.Contains("four"))

	// 更新值
	fmt.Println("\n--- 更新值 ---")
	m.Set("two", 22)
	fmt.Printf("更新后: %s\n", m.String())

	// 删除键值对
	fmt.Println("\n--- 删除键值对 ---")
	removed := m.Remove("one")
	fmt.Printf("删除'one'，值为: %d\n", removed)
	fmt.Printf("删除后: %s\n", m.String())

	fmt.Println("\n=== 示例程序结束 ===")
}
