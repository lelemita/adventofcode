// https://adventofcode.com/2021/day/11
package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		for _, v := range bs {
			line = append(line, int(v)-int(byte('0')))
		}
		output = append(output, line)
	}
	return output
}

func part01(input [][]int) int {
	result := 0
	for s := 0; s < 100; s++ {
		// for _, v := range input {
		// 	fmt.Printf("%v\n", v)
		// }
		// fmt.Println()

		var memo = map[string]bool{}
		for i, line := range input {
			for j := range line {
				flash(&input, i, j, &memo)
			}
		}

		// 결과확인
		for i, line := range input {
			for j, p := range line {
				if p > 9 {
					input[i][j] = 0
					result += 1
				}
			}
		}
	}
	return result
}

func flash(input *[][]int, i, j int, memo *map[string]bool) {
	key := fmt.Sprintf("%d,%d", i, j)
	if _, ex := (*memo)[key]; ex {
		return
	} else if (*input)[i][j] < 9 {
		(*input)[i][j] += 1
		return
	}

	(*memo)[key] = true
	maxI := len(*input) - 1
	var arrI []int
	if i == 0 {
		arrI = []int{0, 1}
	} else if i == maxI {
		arrI = []int{i - 1, i}
	} else {
		arrI = []int{i - 1, i, i + 1}
	}

	var arrJ []int
	if j == 0 {
		arrJ = []int{0, 1}
	} else if j == len((*input)[i])-1 {
		arrJ = []int{j - 1, j}
	} else {
		arrJ = []int{j - 1, j, j + 1}
	}

	for _, y := range arrI {
		for _, x := range arrJ {
			if i == y && j == x {
				(*input)[y][x] += 1
			} else {
				flash(input, y, x, memo)
			}
		}
	}
}

func main() {
	// input := readInput("sample.txt")
	// fmt.Println(part01(input))
	input := readInput("example.txt")
	fmt.Println(part01(input))
	// fmt.Println(part02(input))

	input = readInput("input.txt")
	fmt.Println(part01(input))
	// fmt.Println(part02(input))
}
