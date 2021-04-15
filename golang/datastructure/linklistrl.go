package datastructure

import "fmt"

// 双向链表

type rlNode struct {
	data interface{}
	next *rlNode
	prev *rlNode
}

type rlLinklist struct {
	headNode *rlNode
	tailNode *rlNode
	size int
}

func RlLinklist() *rlLinklist {
	return &rlLinklist{
		headNode: &rlNode{},
		tailNode: &rlNode{},
		size: 0,
	}
}

func (ll *rlLinklist) Add(index int, data interface{}) {
	if index < 0 || index > ll.size {
		panic("Index out of linklist range.")
	}
	if ll.size == 0 {
		ll.headNode.next = ll.tailNode
		ll.tailNode.prev = ll.headNode
	}
	pre := ll.headNode
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	newNode := rlNode{data: data}

	newNode.next = pre.next
	newNode.prev = pre

	pre.next.prev = &newNode
	pre.next = &newNode
	ll.size++
}

func (ll *rlLinklist) AddFirst(data interface{}) {
	ll.Add(0, data)
}

func (ll *rlLinklist) AddLast(data interface{}) {
	ll.Add(ll.size, data)
}

func (ll *rlLinklist) Remove(index int) interface{} {
	if index < 0 || index > ll.size {
		panic("Index out of linklist range.")
	}
	pre := ll.headNode
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	delNode := pre.next
	delNode.next.prev = pre

	pre.next = delNode.next
	delNode.next = nil
	delNode.prev = nil
	ll.size--
	return delNode.data
}

func (ll *rlLinklist) RemoveFirst() interface{} {
	return ll.Remove(0)
}

func (ll *rlLinklist) RemoveLast() interface{} {
	return ll.Remove(ll.size)
}

func (ll *rlLinklist) RlToString() {
	pre := ll.headNode.next
	num := 0
	for pre != nil {
		if num > 0 && num < ll.size {
			fmt.Print("->")
		}
		fmt.Print(pre.data)
		num++
		pre = pre.next
	}
	fmt.Println()

	num1 := 0
	pre2 := ll.tailNode.prev
	for pre2 != nil {
		if num1 > 0 && num1 < ll.size {
			fmt.Print("->")
		}
		fmt.Print(pre2.data)
		num1++
		pre2 = pre2.prev
	}
	fmt.Println()
}