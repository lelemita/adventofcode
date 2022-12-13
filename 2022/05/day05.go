package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/lelemita/adventofcode/common"
)

//go:embed input.txt
var input_day string
var init_day = [][]rune{
	{'S', 'M', 'R', 'N', 'W', 'J', 'V', 'T'},
	{'B', 'W', 'D', 'J', 'Q', 'P', 'C', 'V'},
	{'B', 'J', 'F', 'H', 'D', 'R', 'P'},
	{'F', 'R', 'P', 'B', 'M', 'N', 'D'},
	{'H', 'V', 'R', 'P', 'T', 'B'},
	{'C', 'B', 'P', 'T'},
	{'B', 'J', 'R', 'P', 'L'},
	{'N', 'C', 'S', 'L', 'T', 'Z', 'B', 'W'},
	{'L', 'S', 'G'},
}

func Part1(input string, init [][]rune) string {
	stacks := []common.Stack[rune]{}
	for _, s := range init {
		stacks = append(stacks, common.NewRuneStackWithData(s))
	}

	// read move
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var count, from, to int
		n, _ := fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
		if n == 0 {
			continue
		}
		from -= 1
		to -= 1
		for i := 0; i < count; i++ {
			stacks[to].Push(*stacks[from].Pop())
		}
	}

	// return status
	result := []rune{}
	for _, v := range stacks {
		result = append(result, *v.Pop())
	}
	return string(result)
}

func Part2(input string, init [][]rune) string {
	stacks := []common.Stack[rune]{}
	for _, s := range init {
		stacks = append(stacks, common.NewRuneStackWithData(s))
	}

	// read move
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var count, from, to int
		n, _ := fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
		if n == 0 {
			continue
		}
		from -= 1
		to -= 1
		temp := common.NewRuneStack()
		for i := 0; i < count; i++ {
			temp.Push(*stacks[from].Pop())
		}
		for !temp.IsEmpty() {
			stacks[to].Push(*temp.Pop())
		}
	}

	// return status
	result := []rune{}
	for _, v := range stacks {
		result = append(result, *v.Pop())
	}
	return string(result)
}

func main() {
	fmt.Println("--2022 day 05 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day, init_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day, init_day))
	fmt.Println(time.Since(start))
}
