package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input_day string

func hasDupli(arr []byte) bool {
	for i, a := range arr {
		for j, b := range arr {
			if i != j && a == b {
				return true
			}
		}
	}
	return false
}

func Part1(input string) int {
	lastFour := []byte{}
	for i := 0; i < len(input); i++ {
		lastFour = append(lastFour, input[i])
		if i < 3 {
			continue
		}
		if len(lastFour) > 4 {
			lastFour = lastFour[1:]
		}
		if !hasDupli(lastFour) {
			return i + 1
		}
	}
	return -1
}

func Part2(input string) int {
	lastFour := []byte{}
	for i := 0; i < len(input); i++ {
		lastFour = append(lastFour, input[i])
		if i < 13 {
			continue
		}
		if len(lastFour) > 14 {
			lastFour = lastFour[1:]
		}
		if !hasDupli(lastFour) {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println("--2022 day 06 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
