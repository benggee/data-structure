package datastructure

import (
	"testing"
)

func TestRlAdd(t *testing.T) {
	l1 := RlLinklist()
	for i := 0; i < 100; i++ {
		l1.Add(0, i)
	}
	l1.RlToString()
}


func TestRlYeild(t *testing.T) {
	l1 := RlLinklist()
	for i := 0; i < 100; i++ {
		l1.Add(0, i)
	}

	l1.RlToString()
}

func TestRlRemove(t *testing.T) {
	l1 := RlLinklist()
	for i := 0; i < 20; i++ {
		l1.Add(0, i)
	}

	l1.Remove(1)

	l1.RlToString()
}