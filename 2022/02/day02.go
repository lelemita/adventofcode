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
	var score map[string]int = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	result := 0
	for _, v := range lines {
		var opp, mine string
		fmt.Sscanf(v, "%s %s", &opp, &mine)
		result += score[mine]
		if score[mine] == score[opp] {
			result += 3
		} else if mine == "Z" && opp == "A" {
			result += 0
		} else if mine == "X" && opp == "C" {
			result += 6
		} else if score[mine] > score[opp] {
			result += 6
		}
	}
	return result
}

func Part2(input string) int {
	var score map[string]int = map[string]int{
		"A X": 0 + 3,
		"A Y": 3 + 1,
		"A Z": 6 + 2,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 0 + 2,
		"C Y": 3 + 3,
		"C Z": 6 + 1,
	}
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	result := 0
	for _, v := range lines {
		result += score[v]
	}
	return result
}

func main() {
	fmt.Println("--2022 day 02 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
