// https://adventofcode.com/2021/day/9
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const WALL = byte('9')

func readInput(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var output = [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bs := scanner.Bytes()
		line := []byte{}
		line = append(line, bs...)
		output = append(output, line)
	}
	return output
}

func part01(input [][]byte) int {
	result := 0
	maxI := len(input) - 1
	for i, line := range input {
		var arrI []int
		if i == 0 {
			arrI = []int{1}
		} else if i == maxI {
			arrI = []int{i - 1}
		} else {
			arrI = []int{i - 1, i + 1}
		}
		maxJ := len(line) - 1
		for j, p := range line {
			var arrJ []int
			if j == 0 {
				arrJ = []int{1}
			} else if j == maxJ {
				arrJ = []int{j - 1}
			} else {
				arrJ = []int{j - 1, j + 1}
			}
			isMin := true
			for _, y := range arrI {
				if p >= input[y][j] {
					isMin = false
					break
				}
			}
			for _, x := range arrJ {
				if p >= input[i][x] {
					isMin = false
					break
				}
			}
			if isMin {
				result += int(p) - 47
			}
		}
	}
	return result
}

func part02(input [][]byte) int {
	res := []int{0, 0, 0}
	for i, line := range input {
		for j, p := range line {
			if p < WALL {
				cnt := 0
				getSize(&input, i, j, &cnt)
				if cnt > 0 {
					sort.Ints(res)
					if cnt > res[0] {
						res[0] = cnt
					}
				}

			}
		}
	}
	return res[0] * res[1] * res[2]
}

func getSize(input *[][]byte, i, j int, cnt *int) {
	if (*input)[i][j] == WALL {
		return
	}
	(*input)[i][j] = WALL
	(*cnt) += 1

	maxI := len(*input) - 1
	line := (*input)[i]
	maxJ := len(line) - 1
	var arrI []int
	if i == 0 {
		arrI = []int{1}
	} else if i == maxI {
		arrI = []int{i - 1}
	} else {
		arrI = []int{i - 1, i + 1}
	}
	var arrJ []int
	if j == 0 {
		arrJ = []int{1}
	} else if j == maxJ {
		arrJ = []int{j - 1}
	} else {
		arrJ = []int{j - 1, j + 1}
	}

	// 상하
	for _, y := range arrI {
		if (*input)[y][j] < WALL {
			getSize(input, y, j, cnt)
		}
	}
	// 좌우
	for _, x := range arrJ {
		if (*input)[i][x] < WALL {
			getSize(input, i, x, cnt)
		}
	}
}

func main() {
	input := readInput("example.txt")
	fmt.Println(part01(input))
	fmt.Println(part02(input))
	input = readInput("input.txt")
	fmt.Println(part01(input))
	fmt.Println(part02(input))
}
