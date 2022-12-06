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
		ch := getDuplicate(line) + 1
		if ch > 'a' {
			ch -= 'a'
		} else {
			ch -= 'A' - 26
		}
		result += int(ch)
	}
	return result
}

func getDuplicate(line string) byte {
	for i := 0; i < len(line)/2; i++ {
		for j := len(line) / 2; j < len(line); j++ {
			if line[i] == line[j] {
				return line[i]
			}
		}
	}
	return 0
}

func getTriplicate(lines [3]string) rune {
	candis := []rune{}
	for _, x := range lines[0] {
		for _, y := range lines[1] {
			if x == y {
				candis = append(candis, x)
			}
		}
	}
	for _, v := range candis {
		for _, z := range lines[2] {
			if v == z {
				return v
			}
		}
	}
	return 0
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	result := 0
	threeLines := [3]string{}
	for i, line := range lines {
		threeLines[i%3] = line
		if i%3 != 2 {
			continue
		}
		ch := getTriplicate(threeLines) + 1
		if ch > 'a' {
			ch -= 'a'
		} else {
			ch -= 'A' - 26
		}
		result += int(ch)
	}
	return result
}

func main() {
	fmt.Println("--2022 day 03 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
