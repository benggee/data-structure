package datastructure

import "fmt"

type loopNode struct {
	data interface{}
	next *loopNode
	prev *loopNode
}

type loopLinkList struct {
	headNode *loopNode
	size int
	listLen int
}

func LoopLinkList(listLen int) *loopLinkList {
	return &loopLinkList{
		headNode: &loopNode{},
		size: 0,
		listLen: listLen,
	}
}

func (l *loopLinkList) Add(index int,data interface{}) {
	if index < 0 || index > l.size {
		panic("Index out of link list range.")
	}
	if index > l.listLen {
		panic("Index out of link list max length.")
	}
	pre := l.headNode
	for i := 0; i < index; i++ {
		pre = pre.next
	}

	newNode := &loopNode{
		data: data,
	}

	if l.size == 0 {
		newNode.next = pre
		pre.prev = newNode
	} else {
		newNode.next = pre.next
		pre.next.prev = newNode
	}
	newNode.prev = pre
	pre.next = newNode

	l.size++
}

func (l *loopLinkList) AddHead(data interface{}) {
	l.Add(0, data)
}

func (l *loopLinkList) AddEnd(data interface{}) {
	l.Add(l.size, data)
}

func (l *loopLinkList) Remove(index int) {
	if index < 0 || index > l.size {
		panic("Index out of link list range.")
	}
	pre := l.headNode
	for i := 1; i < index; i++ {
		pre = pre.next
	}
	delNode := pre.next
	delNode.next.prev = pre
	pre.next = delNode.next

	l.size--
}

func (l *loopLinkList) RemoveFirst() {
	l.Remove(0)
}

func (l *loopLinkList) RemoveLast() {
	l.Remove(l.size)
}

func (l *loopLinkList) ToString() {
	pre := l.headNode.next
	num := 0
	for pre != nil {
		if num > 0 && num < l.size {
			fmt.Print("->")
		}
		fmt.Print(pre.data)
		num++
		pre = pre.next
		if pre.data == nil {
			break
		}
	}
	fmt.Println()

	num1 := 0
	pre2 := l.headNode.prev
	for pre2 != nil {
		if num1 > 0 && num1 < l.size {
			fmt.Print("->")
		}
		fmt.Print(pre2.data)
		num1++
		pre2 = pre2.prev
		if pre2.data == nil {
			break
		}
	}
	fmt.Println()
}