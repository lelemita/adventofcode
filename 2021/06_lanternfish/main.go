// https://adventofcode.com/2021/day/6
// 피보나치로는 왜 안되지ㅠㅠ
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func solution(puzzle string, day int) int {
	lines := strings.Split(puzzle, ",")
	fishes := make([]int, 9)
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		fishes[i] += 1
	}

	for d := 0; d < day; d++ {
		newFish := fishes[0]
		for f := 1; f < len(fishes); f++ {
			fishes[f-1] = fishes[f]
		}
		fishes[6] += newFish
		fishes[8] = newFish
	}

	sum := 0
	for _, v := range fishes {
		sum += v
	}
	return sum
}

func main() {
	example := readInput("example.txt")
	fmt.Println(solution(example, 80))
	fmt.Println(solution(example, 256))
	puzzle := readInput("input.txt")
	fmt.Println(solution(puzzle, 80))
	fmt.Println(solution(puzzle, 256))
}
