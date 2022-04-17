package main

import (
	"fmt"
)

func main() {
	idx := 1
	for i := 0; i < 5; i++ {
		idx = idx << 1
		fmt.Println(idx)
	}
}
