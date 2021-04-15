package tool

import (
	"fmt"
	"time"
)

func EchoCurrenNaTime() int64 {
	t := time.Now().UnixNano() / 1e6
	fmt.Println()
	return t
}