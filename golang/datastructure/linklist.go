package datastructure

import "fmt"

type node struct {
	data interface{}
	next *node
}

type linkList struct {
	headNode *node
	size int
}

func Linklist() *linkList {
	return &linkList{
		headNode: &node{},
		size: 0,
	}
}

func (ll *linkList) Add(index int, data interface{}) {
	if index < 0 || index > ll.size {
		panic("Index out of linklist range.")
	}
	pre := ll.headNode
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	newNode := node{
		data: data,
	}
	newNode.next = pre.next
	pre.next = &newNode
	ll.size++
}

func (ll *linkList) AddTail(data interface{}) {
	ll.Add(ll.size, data)
}

func (ll *linkList) AddHead(data interface{}) {
	ll.Add(0, data)
}

func (ll *linkList) Remove(index int) *node {
	if index < 0 || index > ll.size {
		panic("Index out linklist range.")
	}
	pre := ll.headNode
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	delNode := pre.next
	pre.next = delNode.next
	delNode.next = nil
	ll.size--
	return delNode
}

func (ll *linkList) RemoveByValue(data interface{}) {
	var delNode *node
	pre := ll.headNode
	for pre != nil {
		if pre.next != nil && pre.next.data == data {
			delNode = pre.next
			break
		}
		pre = pre.next
	}
	if delNode != nil {
		delNode := pre.next
		pre.next = delNode.next
		delNode.next = nil
	}
	ll.size--
}

func (ll *linkList) Contains(data interface{}) bool {
	pre := ll.headNode
	for pre != nil {
		if pre.data == data {
			return true
		}
		pre = pre.next
	}
	return false
}

func (ll *linkList) RemoveFirst() *node {
	return ll.Remove(0)
}

func (ll *linkList) RemoveLast() *node {
	return ll.Remove(ll.size-1)
}

func (ll *linkList) Size() int {
	return ll.size
}

func (ll *linkList) ToString() {
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
}