package datastructure

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := Mapcontains()
	m.Add("aa", "bb")
	m.Add("bb", "cc")
	m.Add("name", "LiuFajun")

	fmt.Println(m.Get("name"))

	 m.Del("name")
	fmt.Println(m.Get("name"))

	fmt.Println(m.Size())

	m.ToString()
}