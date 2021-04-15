package datastructure

import (
	"fmt"
	"testing"
)

func TestQueuePush(t *testing.T) {
	q := Queue()
	q.Push("aaaa")
	q.Push("bbb")

	fmt.Println(q.Pop())
	fmt.Println(q.Size())
	fmt.Println(q.Pop())
	fmt.Println(q.Size())
	fmt.Println(q.Pop())

	fmt.Println(q.Size())
}

