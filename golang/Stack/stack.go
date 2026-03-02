package main

import (
	"fmt"
	"strings"
)

// Stack 栈接口
type Stack[E any] interface {
	Size() int
	IsEmpty() bool
	Push(e E)
	Pop() E
	Peek() E
}

// ArrayStack 基于动态数组的栈实现
type ArrayStack[E any] struct {
	array *Array[E]
}

// NewArrayStack 创建一个新的数组栈
func NewArrayStack[E any](capacity int) *ArrayStack[E] {
	return &ArrayStack[E]{
		array: NewArray[E](capacity),
	}
}

// NewArrayStackDefault 创建默认容量的数组栈
func NewArrayStackDefault[E any]() *ArrayStack[E] {
	return &ArrayStack[E]{
		array: NewArrayDefault[E](),
	}
}

func (s *ArrayStack[E]) Size() int {
	return s.array.Size()
}

func (s *ArrayStack[E]) IsEmpty() bool {
	return s.array.IsEmpty()
}

func (s *ArrayStack[E]) Push(e E) {
	s.array.AddLast(e)
}

func (s *ArrayStack[E]) Pop() E {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	return s.array.RemoveLast()
}

func (s *ArrayStack[E]) Peek() E {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	return s.array.Get(s.array.Size() - 1)
}

func (s *ArrayStack[E]) String() string {
	var builder strings.Builder
	builder.WriteString("Stack: [")
	for i := 0; i < s.array.Size(); i++ {
		builder.WriteString(fmt.Sprintf("%v", s.array.Get(i)))
		if i < s.array.Size()-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("] <- top")
	return builder.String()
}

func main() {
	fmt.Println("=== Go语言栈示例 ===")

	stack := NewArrayStackDefault[int]()
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Printf("入栈 %d: %s\n", i, stack.String())
	}

	fmt.Printf("栈顶元素: %d\n", stack.Peek())
	fmt.Printf("出栈: %d\n", stack.Pop())
	fmt.Printf("出栈后: %s\n", stack.String())

	fmt.Printf("栈大小: %d\n", stack.Size())
	fmt.Printf("是否为空: %t\n", stack.IsEmpty())

	// 清空栈
	for !stack.IsEmpty() {
		stack.Pop()
	}
	fmt.Printf("清空后是否为空: %t\n", stack.IsEmpty())

	fmt.Println("\n=== 示例程序结束 ===")
}
