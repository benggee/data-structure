package datastructure

import "fmt"

type mapNode struct {
	key string
	data interface{}
	next *mapNode
}

type mapLinkList struct {
	headNode *mapNode
	size int
}

type mapData struct {
	data *mapLinkList
}

func Mapcontains() *mapData {
	return &mapData{data: initMapLinkList()}
}

func (m mapData) Add(key string, data interface{}) {
	m.data.add(key, data)
}

func (m mapData) Del(key string) {
	m.data.remove(key)
}

func (m mapData) Size() int {
	return m.data.size
}

func (m mapData) Get(key string) interface{} {
	return m.data.get(key)
}

func (m mapData) Set(key string, data interface{}) {

}

func (m mapData) ToString() {
	m.data.toString()
}

func initMapLinkList() *mapLinkList {
	return &mapLinkList{headNode: &mapNode{},size:0}
}

func (ml *mapLinkList) add(key string, data interface{}) {
	newNode := mapNode{key: key, data: data}
	newNode.next = ml.headNode.next
	ml.headNode.next = &newNode
	ml.size++
}

func (ml *mapLinkList) remove(key string) {
	var delNode *mapNode
	pre := ml.headNode
	for pre != nil {
		if pre.next != nil && pre.next.key == key {
			break
		}
		pre = pre.next
	}
	if pre.next != nil && pre.next.data != nil {
		delNode = pre.next
		pre.next = delNode.next
		delNode.next = nil
	}
	ml.size--
}

func (ml *mapLinkList) get(key string) interface{} {
	pre := ml.headNode
	for pre != nil {
		if pre.key == key {
			return pre.data
		}
		pre = pre.next
	}
	return nil
}

func (ml *mapLinkList) update(key string, data interface{}) {
	pre := ml.headNode
	for pre != nil {
		if pre.key == key {
			pre.data = data
		}
		pre = pre.next
	}
}

func (ml *mapLinkList) toString() {
	pre := ml.headNode.next
	num := 0
	for pre != nil {
		if num > 0 && num < ml.size {
			fmt.Print("->")
		}
		fmt.Print(pre.data)
		num++
		pre = pre.next
	}
	fmt.Println()
}