package datastructure

import "fmt"

type array struct {
	dataSize int
	data []interface{}
}

func Array(size int) *array {
	arr := &array{
		dataSize: 0,
		data: []interface{}{},
	}
	arr.reSize(size)
	return arr
}

// 在任意位置插入元素
func (a *array)	Add(index int, data interface{}) {
	if index < 0 || index >= len(a.data) {
		panic("Index out of range.")
	}
	if a.dataSize+1 == len(a.data) {
		a.reSize(len(a.data) * 2)
	}

	arrayLen := len(a.data)
	for i := arrayLen - 2; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = data
	a.dataSize++
}

// 删除指定索引的元素
func (a *array) Remove(index int) {
	if a.data[index] != struct {}{} {
		a.dataSize--
	}
	a.data = append(a.data[0:index], a.data[index+1:len(a.data)]...)
}

// 删除第一个元素
func (a *array) RemoveFirst() {
	a.Remove(0)
}

// 删除最后一个
func (a *array) RemoveLast() {
	a.Remove(len(a.data)-1)
}

// 添加到数组开头
func (a *array) AddHead(data interface{}) {
	a.Add(0, data)
}

// 添加到数组开头
func (a *array) AddTail(data interface{}) {
	a.Add(len(a.data)-1, data)
}

// 获取元素个数
func (a *array) Size() int {
	return a.dataSize
}

func (a *array) reSize(newSize int) {
	if a.dataSize == newSize {
		return
	}
	tmpSlice := []interface{}{}
	for _, v := range a.data {
		if v != struct {}{} {
			tmpSlice = append(tmpSlice, v)
		}
	}
	a.dataSize = len(tmpSlice)

	for i := a.dataSize; i < newSize; i++ {
		tmpSlice = append(tmpSlice, struct {}{})
	}
	a.data = tmpSlice
}

// toString
func (a *array) ToString() {
	for k, v := range a.data {
		fmt.Println(k, "=>", v)
	}
}
