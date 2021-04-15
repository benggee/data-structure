package datastructure

import "testing"

func TestAdd(t *testing.T) {
	l1 := Linklist()
	for i := 0; i < 100; i++ {
		l1.Add(0, i)
	}

	l1.ToString()
}

func TestAddTail(t *testing.T) {
	l1 := Linklist()
	for i := 0; i < 50; i++ {
		l1.AddTail(i)
	}
	l1.ToString()
}

func TestLinkListAddHead(t *testing.T) {
	l1 := Linklist()
	for i := 0; i < 40; i++ {
		l1.AddHead(i)
	}
	l1.ToString()
}

func TestLinklistRemove(t *testing.T) {
	l1 := Linklist()
	for i := 0; i < 40; i++ {
		l1.AddHead(i)
	}
	l1.ToString()
	l1.Remove(3)
	l1.ToString()
}