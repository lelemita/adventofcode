package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	lines = append(lines, "")
	result := 0
	one := 0
	for _, v := range lines {
		if len(v) == 0 {
			if one > result {
				result = one
			}
			one = 0
		}
		cal, _ := strconv.Atoi(v)
		one += cal
	}
	return result
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	lines = append(lines, "")
	result := sort.IntSlice{}
	one := 0
	for _, v := range lines {
		if len(v) == 0 {
			if result.Len() < 3 {
				result = append(result, one)
			} else if one > result[0] {
				result[0] = one
				result.Sort()
			}
			one = 0
		}
		cal, _ := strconv.Atoi(v)
		one += cal
	}

	return result[0] + result[1] + result[2]
}

func main() {
	fmt.Println("--2022 day 01 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
