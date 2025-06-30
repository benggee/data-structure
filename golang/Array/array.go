package main

import (
	"fmt"
	"reflect"
)

// Array 动态数组结构体，使用泛型实现
// E 表示数组元素的类型
type Array[E any] struct {
	arr  []E // 底层数组
	size int // 当前元素数量
}

// NewArray 创建一个指定容量的新数组
// capacity: 数组的初始容量
func NewArray[E any](capacity int) *Array[E] {
	return &Array[E]{
		arr:  make([]E, capacity),
		size: 0,
	}
}

// NewArrayDefault 创建一个默认容量(10)的新数组
func NewArrayDefault[E any]() *Array[E] {
	return NewArray[E](10)
}

// NewArrayFromSlice 从切片创建新数组
// slice: 用于初始化数组的切片
func NewArrayFromSlice[E any](slice []E) *Array[E] {
	arr := NewArray[E](len(slice))
	for i, v := range slice {
		arr.arr[i] = v
	}
	arr.size = len(slice)
	return arr
}

// Capacity 获取数组容量
func (a *Array[E]) Capacity() int {
	return len(a.arr)
}

// Size 获取元素数量
func (a *Array[E]) Size() int {
	return a.size
}

// IsEmpty 判断数组是否为空
func (a *Array[E]) IsEmpty() bool {
	return a.size == 0
}

// Add 在指定索引位置插入元素
// index: 插入位置索引
// e: 要插入的元素
func (a *Array[E]) Add(index int, e E) {
	if index < 0 || index > a.size {
		panic("Add failed. index must be required index>=0 and index<=size")
	}

	// 如果数组已满，进行扩容
	if a.size == len(a.arr) {
		a.resize(2 * len(a.arr))
	}

	// 将index位置及之后的元素向后移动
	for i := a.size - 1; i >= index; i-- {
		a.arr[i+1] = a.arr[i]
	}

	a.arr[index] = e
	a.size++
}

// AddFirst 在数组开头添加一个元素
func (a *Array[E]) AddFirst(e E) {
	a.Add(0, e)
}

// AddLast 在数组结尾添加一个元素
func (a *Array[E]) AddLast(e E) {
	a.Add(a.size, e)
}

// Get 获取指定索引位置的元素
// index: 元素索引
func (a *Array[E]) Get(index int) E {
	if index < 0 || index >= a.size {
		panic("The index out of range size")
	}
	return a.arr[index]
}

// Set 设置指定索引位置的值
// index: 索引位置
// e: 新值
func (a *Array[E]) Set(index int, e E) {
	if index < 0 || index >= a.size {
		panic("The index out of range size")
	}
	a.arr[index] = e
}

// Swap 交换两个索引位置的元素
// i, j: 要交换的两个索引位置
func (a *Array[E]) Swap(i, j int) {
	if i < 0 || i >= a.size || j < 0 || j >= a.size {
		panic("The i or j out of range")
	}
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
}

// Find 查找元素所在索引位置，如果不存在返回-1
// e: 要查找的元素
func (a *Array[E]) Find(e E) int {
	for i := 0; i < a.size; i++ {
		if reflect.DeepEqual(a.arr[i], e) {
			return i
		}
	}
	return -1
}

// Contains 判断元素是否在数组内
// e: 要查找的元素
func (a *Array[E]) Contains(e E) bool {
	return a.Find(e) != -1
}

// Remove 删除指定索引位置的元素
// index: 要删除的索引位置
func (a *Array[E]) Remove(index int) E {
	if index < 0 || index >= a.size {
		panic("Index out of range")
	}

	ret := a.arr[index]

	// 将index位置之后的元素向前移动
	for i := index + 1; i < a.size; i++ {
		a.arr[i-1] = a.arr[i]
	}

	a.size--

	// 如果元素数量减少到容量的一半，进行缩容
	if a.size == len(a.arr)/2 && len(a.arr)/2 != 0 {
		a.resize(len(a.arr) / 2)
	}

	return ret
}

// RemoveLast 删除数组最后一个元素
func (a *Array[E]) RemoveLast() E {
	return a.Remove(a.size - 1)
}

// RemoveFirst 删除数组第一个元素
func (a *Array[E]) RemoveFirst() E {
	return a.Remove(0)
}

// RemoveElement 删除指定的元素
// e: 要删除的元素
func (a *Array[E]) RemoveElement(e E) {
	index := a.Find(e)
	if index != -1 {
		a.Remove(index)
	}
}

// String 返回数组的字符串表示
func (a *Array[E]) String() string {
	if a.size == 0 {
		return "[]"
	}

	result := "["
	for i := 0; i < a.size; i++ {
		result += fmt.Sprintf("%v", a.arr[i])
		if i != a.size-1 {
			result += ", "
		}
	}
	result += "]"
	return result
}

// resize 调整数组容量
// newCapacity: 新的容量大小
func (a *Array[E]) resize(newCapacity int) {
	newArr := make([]E, newCapacity)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.arr[i]
	}
	a.arr = newArr
}

// ToSlice 将数组转换为切片
func (a *Array[E]) ToSlice() []E {
	result := make([]E, a.size)
	for i := 0; i < a.size; i++ {
		result[i] = a.arr[i]
	}
	return result
}

// Clear 清空数组
func (a *Array[E]) Clear() {
	a.size = 0
	// 重置为默认容量
	a.arr = make([]E, 10)
}

// 示例程序
func main() {
	fmt.Println("=== Go语言动态数组示例 ===")

	// 创建默认容量的数组
	arr := NewArrayDefault[int]()
	fmt.Printf("创建默认数组，容量: %d, 大小: %d, 是否为空: %t\n",
		arr.Capacity(), arr.Size(), arr.IsEmpty())

	// 添加元素
	fmt.Println("\n--- 添加元素 ---")
	arr.AddLast(1)
	arr.AddLast(2)
	arr.AddLast(3)
	arr.AddFirst(0)
	arr.Add(2, 10)
	fmt.Printf("数组内容: %s\n", arr.String())
	fmt.Printf("容量: %d, 大小: %d\n", arr.Capacity(), arr.Size())

	// 获取和设置元素
	fmt.Println("\n--- 获取和设置元素 ---")
	fmt.Printf("索引2的元素: %d\n", arr.Get(2))
	arr.Set(2, 20)
	fmt.Printf("设置索引2为20后: %s\n", arr.String())

	// 查找元素
	fmt.Println("\n--- 查找元素 ---")
	fmt.Printf("元素20的索引: %d\n", arr.Find(20))
	fmt.Printf("是否包含元素5: %t\n", arr.Contains(5))
	fmt.Printf("是否包含元素1: %t\n", arr.Contains(1))

	// 交换元素
	fmt.Println("\n--- 交换元素 ---")
	fmt.Printf("交换前: %s\n", arr.String())
	arr.Swap(0, 4)
	fmt.Printf("交换索引0和4后: %s\n", arr.String())

	// 删除元素
	fmt.Println("\n--- 删除元素 ---")
	fmt.Printf("删除前: %s\n", arr.String())
	removed := arr.Remove(2)
	fmt.Printf("删除索引2的元素(%d)后: %s\n", removed, arr.String())

	removedFirst := arr.RemoveFirst()
	fmt.Printf("删除第一个元素(%d)后: %s\n", removedFirst, arr.String())

	removedLast := arr.RemoveLast()
	fmt.Printf("删除最后一个元素(%d)后: %s\n", removedLast, arr.String())

	// 测试字符串数组
	fmt.Println("\n--- 字符串数组测试 ---")
	strArr := NewArrayFromSlice[string]([]string{"Hello", "World", "Go"})
	fmt.Printf("字符串数组: %s\n", strArr.String())
	strArr.AddLast("Language")
	fmt.Printf("添加元素后: %s\n", strArr.String())

	// 测试扩容
	fmt.Println("\n--- 扩容测试 ---")
	testArr := NewArray[int](2)
	fmt.Printf("初始容量: %d\n", testArr.Capacity())
	for i := 0; i < 10; i++ {
		testArr.AddLast(i)
		fmt.Printf("添加%d后，容量: %d, 大小: %d\n", i, testArr.Capacity(), testArr.Size())
	}

	// 测试缩容
	fmt.Println("\n--- 缩容测试 ---")
	for i := 0; i < 8; i++ {
		testArr.RemoveLast()
		fmt.Printf("删除后，容量: %d, 大小: %d\n", testArr.Capacity(), testArr.Size())
	}

	fmt.Println("\n=== 示例程序结束 ===")
}
