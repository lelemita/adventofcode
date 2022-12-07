package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {
		var a, b, x, y int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &x, &y)
		if a <= x && y <= b {
			result += 1
		} else if x <= a && b <= y {
			result += 1
		}
	}
	return result
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	result := len(lines)
	for _, line := range lines {
		var a, b, x, y int
		fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &x, &y)
		if b < x || y < a {
			result -= 1
		}
	}
	return result
}

func main() {
	fmt.Println("--2022 day 04 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
