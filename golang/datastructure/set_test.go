package datastructure

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	//l := Linklist()
	//l.AddHead("aaa")
	//l.AddHead("bbb")
	//l.AddHead(1111)
	//l.ToString()
	s := Set()
	s.Add("aaaa")
	s.Add("cccc")
	s.Add("1111")

	s.Del("1111")
	s.Del("cccc")
	s.Del("aaaa")
	s.Del("bbbb")

	s.Add("aaaa")

	fmt.Println(s.Contains("aaaa"))

	s.ToString()
}