package main

import (
	"fmt"
	"strings"
)

// LinkList 链表实现
// 使用虚拟头节点（dummy head）来简化操作
type LinkList[E any] struct {
	dummyHead *Node[E] // 虚拟头节点
	size      int      // 链表大小
}

// Node 链表节点
type Node[E any] struct {
	e    E        // 节点数据
	next *Node[E] // 下一个节点指针
}

// NewLinkList 创建一个新的链表
func NewLinkList[E any]() *LinkList[E] {
	return &LinkList[E]{
		dummyHead: &Node[E]{}, // 创建虚拟头节点
		size:      0,
	}
}

// newNode 创建一个新的节点
func newNode[E any](e E) *Node[E] {
	return &Node[E]{
		e:    e,
		next: nil,
	}
}

// GetSize 返回链表长度
func (l *LinkList[E]) GetSize() int {
	return l.size
}

// IsEmpty 判断链表是否为空
func (l *LinkList[E]) IsEmpty() bool {
	return l.size == 0
}

// Add 在指定索引位置添加元素
// index: 要插入的位置（0-based）
// e: 要插入的元素
func (l *LinkList[E]) Add(index int, e E) {
	if index < 0 || index > l.size {
		panic("Index out of range.")
	}

	// 找到插入位置的前一个节点
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}

	// 创建新节点并插入
	node := newNode(e)
	node.next = pre.next
	pre.next = node
	l.size++
}

// AddFirst 在链表头部添加元素
func (l *LinkList[E]) AddFirst(e E) {
	l.Add(0, e)
}

// AddLast 在链表尾部添加元素
func (l *LinkList[E]) AddLast(e E) {
	l.Add(l.size, e)
}

// Get 获取指定索引位置的元素
func (l *LinkList[E]) Get(index int) E {
	if index < 0 || index >= l.size {
		panic("Index out of range.")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

// GetFirst 获取第一个元素
func (l *LinkList[E]) GetFirst() E {
	return l.Get(0)
}

// GetLast 获取最后一个元素
func (l *LinkList[E]) GetLast() E {
	return l.Get(l.size - 1)
}

// Set 设置指定索引位置的元素
func (l *LinkList[E]) Set(index int, e E) {
	if index < 0 || index >= l.size {
		panic("Index out of range.")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

// Find 查找元素是否在链表中
// 注意：这里使用类型断言来比较元素，对于复杂类型需要实现自定义比较方法
func (l *LinkList[E]) Find(e E) bool {
	cur := l.dummyHead.next
	for cur != nil {
		// 使用类型断言进行相等性比较
		if l.equal(cur.e, e) {
			return true
		}
		cur = cur.next
	}
	return false
}

// Del 删除指定索引位置的元素
func (l *LinkList[E]) Del(index int) E {
	if index < 0 || index >= l.size {
		panic("Index out of range.")
	}

	// 找到要删除节点的前一个节点
	pre := l.dummyHead
	for i := 0; i < index; i++ {
		pre = pre.next
	}

	// 删除节点
	delNode := pre.next
	pre.next = delNode.next
	delNode.next = nil
	l.size--

	return delNode.e
}

// DelFirst 删除第一个元素
func (l *LinkList[E]) DelFirst() E {
	return l.Del(0)
}

// DelLast 删除最后一个元素
func (l *LinkList[E]) DelLast() E {
	return l.Del(l.size - 1)
}

// Remove 删除指定元素（如果存在）
func (l *LinkList[E]) Remove(e E) {
	// 找到要删除元素的前一个节点
	pre := l.dummyHead
	for pre.next != nil {
		if l.equal(pre.next.e, e) {
			// 删除节点
			delNode := pre.next
			pre.next = delNode.next
			delNode.next = nil
			l.size--
			return
		}
		pre = pre.next
	}
}

// String 返回链表的字符串表示
func (l *LinkList[E]) String() string {
	var res strings.Builder
	cur := l.dummyHead.next
	for cur != nil {
		res.WriteString(fmt.Sprintf("%v->", cur.e))
		cur = cur.next
	}
	res.WriteString("NULL")
	return res.String()
}

// equal 比较两个元素是否相等
// 这里使用类型断言，对于复杂类型需要实现自定义比较方法
func (l *LinkList[E]) equal(a, b E) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

// ToSlice 将链表转换为切片
func (l *LinkList[E]) ToSlice() []E {
	result := make([]E, l.size)
	cur := l.dummyHead.next
	for i := 0; i < l.size; i++ {
		result[i] = cur.e
		cur = cur.next
	}
	return result
}

// Clear 清空链表
func (l *LinkList[E]) Clear() {
	l.dummyHead.next = nil
	l.size = 0
}

// Reverse 反转链表
func (l *LinkList[E]) Reverse() {
	if l.size <= 1 {
		return
	}

	var prev *Node[E] = nil
	cur := l.dummyHead.next
	var next *Node[E]

	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	l.dummyHead.next = prev
}

// HasCycle 检测链表是否有环
func (l *LinkList[E]) HasCycle() bool {
	if l.size <= 1 {
		return false
	}

	slow := l.dummyHead.next
	fast := l.dummyHead.next.next

	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}

	return false
}
