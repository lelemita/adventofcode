package main

import (
	"fmt"
)

func main() {
	arr := []int{}
	cnts := map[int]int{}
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				fmt.Print(i+j+k, ", ")
				arr = append(arr, i+j+k)
				cnts[i+j+k] += 1
			}
		}
	}
	fmt.Println()
	fmt.Println(len(arr))
	fmt.Println(cnts)
}
