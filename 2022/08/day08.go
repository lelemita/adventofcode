package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lelemita/adventofcode/common"
)

//go:embed input.txt
var input_day string

const maximum = 9

var forest = [][]int{}
var isCounted = map[string]bool{}

func read(input string) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trees := []int{}
		for _, ch := range line {
			tree, err := strconv.Atoi(string(ch))
			common.PanicIfErr(err)
			trees = append(trees, tree)
		}
		forest = append(forest, trees)
	}
}

func Part1(input string) int {
	read(input)

	// 서쪽
	for i := 1; i < len(forest)-1; i++ {
		before := -1
		for j := 0; j < len(forest[0])-1; j++ {
			before = checkTree(before, i, j)
			if before == maximum {
				break
			}
		}
	}

	// 동쪽
	for i := 1; i < len(forest)-1; i++ {
		before := -1
		for j := len(forest[0]) - 1; j > 0; j-- {
			if before < forest[i][j] {
				before = checkTree(before, i, j)
				if before == maximum {
					break
				}
			}
		}
	}

	// 북쪽
	for j := 1; j < len(forest[0])-1; j++ {
		before := -1
		for i := 0; i < len(forest)-1; i++ {
			if before < forest[i][j] {
				before = checkTree(before, i, j)
				if before == maximum {
					break
				}
			}
		}
	}

	// 남쪽
	for j := 1; j < len(forest[0])-1; j++ {
		before := -1
		for i := len(forest) - 1; i > 0; i-- {
			before = checkTree(before, i, j)
			if before == maximum {
				break
			}
		}
	}
	return len(isCounted) + 4
}

func checkTree(before int, i int, j int) int {
	if before < forest[i][j] {
		before = forest[i][j]
		axies := fmt.Sprintf("%d,%d", i, j)
		isCounted[axies] = true
	}
	return before
}

func Part2(input string) int {
	if len(forest) == 0 {
		read(input)
	}

	result := 0
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			// 동서남북
			score := 1
			cnt := 0
			for x := j + 1; x < len(forest[i]); x++ {
				cnt += 1
				if forest[i][j] <= forest[i][x] {
					break
				}
			}
			score *= cnt

			cnt = 0
			for x := j - 1; x >= 0; x-- {
				cnt += 1
				if forest[i][j] <= forest[i][x] {
					break
				}
			}
			score *= cnt

			cnt = 0
			for y := i + 1; y < len(forest); y++ {
				cnt += 1
				if forest[i][j] <= forest[y][j] {
					break
				}
			}
			score *= cnt

			cnt = 0
			for y := i - 1; y >= 0; y-- {
				cnt += 1
				if forest[i][j] <= forest[y][j] {
					break
				}
			}
			score *= cnt

			// 갱신
			if score > result {
				result = score
			}
		}
	}
	return result
}

func main() {
	fmt.Println("--2022 day 08 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
