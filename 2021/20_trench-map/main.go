// 실패1: algo[0] --> # 임...ㅠㅠ
// 더 빠른 방법: https://github.com/pemoreau/advent-of-code/blob/main/go/2021/20/day20.go

package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
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

	start := time.Now()
	fmt.Println("part01: ", solution(2))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("part02: ", solution(50))
	fmt.Println(time.Since(start))
}

func solution(count int) int {
	addLines('.')
	for i := 0; i < count; i++ {
		// 사방에 붙이기
		addLines(data[0][0])

		// 이미지 개선
		before := make([]string, len(data))
		copy(before, data)
		advance(before)
	}
	return getCount()
}

func addLines(dot byte) {
	for i := 0; i < len(data); i++ {
		data[i] = fmt.Sprintf("%s%s%s", string(dot), data[i], string(dot))
	}
	// up, down
	newBytes := []byte{}
	for i := 0; i < len(data[0]); i++ {
		newBytes = append(newBytes, dot)
	}
	data = append([]string{string(newBytes)}, data...)
	data = append(data, string(newBytes))
}

func advance(before []string) {
	defVal := before[0][0]
	for i := 0; i < len(before); i++ {
		for j := 0; j < len(before[i]); j++ {
			bins := getAround(i, j, defVal)
			num := getNum(bins)
			before[i] = fmt.Sprintf("%s%s%s", before[i][:j], string(algo[num]), before[i][j+1:])
		}
	}
	data = before
}

func getAround(row, col int, defByte byte) [9]int {
	defVal := 0
	if defByte == '#' {
		defVal = 1
	}
	result := [9]int{}
	cnt := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i < 0 || j < 0 || i > len(data)-1 || j > len(data)-1 {
				result[cnt] = defVal
			} else {
				if data[i][j] == '#' {
					result[cnt] = 1
				} else {
					result[cnt] = 0
				}
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
