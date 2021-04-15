package datastructure

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestHashData(t *testing.T) {
	h := Hash()
	for i := 0; i < 1000; i++ {
		key := "key"+strconv.Itoa(i)
		h.Add(key, i)
	}

	ret, _ := json.Marshal(h.hashTable)

	fmt.Println(string(ret))
}