package datastructure

import "testing"

func TestLoopLinkListAdd(t *testing.T) {
	l := LoopLinkList(10)

	for i := 0; i < 10; i++ {
		l.Add(i, i)
	}

	l.ToString()
}

func TestLoopLinkListRemove(t *testing.T) {
	l := LoopLinkList(10)

	for i := 0; i < 10; i++ {
		l.Add(i, i)
	}

	l.Remove(10)

	l.ToString()
}