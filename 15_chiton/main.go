package main

import (
	"bufio"
	"fmt"
	"os"
)

var memo map[string]int

func main() {
	cave := readInput("example2.txt")
	// cave := readInput("input.txt")
	memo = map[string]int{}
	memo["0,0"] = 0
	for j := 1; j < len(cave[0]); j++ {
		axis := fmt.Sprintf("0,%d", j)
		before := fmt.Sprintf("0,%d", j-1)
		memo[axis] = memo[before] + cave[0][j]
	}
	height := len(cave)
	width := len(cave[0])
	for i := 1; i < height; i++ {
		for j := 0; j < width; j++ {
			up := fmt.Sprintf("%d,%d", i-1, j)
			left := fmt.Sprintf("%d,%d", i, j-1)
			val := memo[up]
			if j > 0 && memo[up] > memo[left] {
				val = memo[left]
			}
			memo[fmt.Sprintf("%d,%d", i, j)] = cave[i][j] + val
			axis := fmt.Sprintf("%d,%d", i, j)
			fmt.Print(axis, "=", memo[axis], " / ")
		}
		fmt.Println()
	}
	fmt.Println(memo[fmt.Sprintf("%d,%d", height-1, width-1)])
}

func readInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var output = [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bs := scanner.Bytes()
		line := []int{}
		for _, b := range bs {
			line = append(line, int(b)-int('0'))
		}
		output = append(output, line)
	}
	return output
}
