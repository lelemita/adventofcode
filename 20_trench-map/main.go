// 실패1: algo[0] --> # 임...ㅠㅠ

package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed example.txt
var puzzle string
var algo string
var data []string

func readInput() {
	lines := strings.Split(strings.ReplaceAll(puzzle, "\r\n", "\n"), "\n")
	algo = lines[0]

	for i := 2; i < len(lines); i++ {
		data = append(data, lines[i])
	}
}

func main() {
	data = []string{}
	readInput()

	count := 2
	for i := 0; i < count; i++ {
		// 사방에 두줄씩 붙이기
		addLines()

		// 이미지 개선
		before := make([]string, len(data))
		copy(before, data)
		advance(before)
		for _, line := range data {
			fmt.Println(line)
		}
	}

	fmt.Println("part01: ", getCount())
}

func addLines() {
	for i := 0; i < len(data); i++ {
		data[i] = fmt.Sprintf("..%s..", data[i])
	}
	// up, down
	newBytes := []byte{}
	for i := 0; i < len(data[0]); i++ {
		newBytes = append(newBytes, '.')
	}
	data = append([]string{string(newBytes), string(newBytes)}, data...)
	data = append(data, string(newBytes))
	data = append(data, string(newBytes))
}

func advance(before []string) {
	for i := 1; i < len(before)-1; i++ {
		for j := 1; j < len(before[i])-1; j++ {
			bins := getAround(i, j)
			num := getNum(bins)
			before[i] = fmt.Sprintf("%s%s%s", before[i][:j], string(algo[num]), before[i][j+1:])
		}
	}
	data = before
}

func getAround(row, col int) [9]int {
	result := [9]int{}
	cnt := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if data[i][j] == '#' {
				result[cnt] = 1
			} else {
				result[cnt] = 0
			}
			cnt += 1
		}
	}
	return result
}

func getNum(bins [9]int) int {
	result := 0
	for i := 0; i < len(bins); i++ {
		result += bins[i] * int(math.Pow(2, float64(8-i)))
	}
	return result
}

func getCount() int {
	result := 0
	for _, line := range data {
		for _, v := range line {
			if v == '#' {
				result += 1
			}
		}
	}
	return result
}
