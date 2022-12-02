package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed example.txt
var puzzle string

func readInput() map[int][][3]int {
	data := map[int][][3]int{}
	lines := strings.Split(strings.ReplaceAll(puzzle, "\r\n", "\n"), "\n")
	scanner := 0
	data[scanner] = [][3]int{}
	for i := 1; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			i += 1
			fmt.Sscanf(lines[i], "--- scanner %d ---", &scanner)
			data[scanner] = [][3]int{}
		} else {
			var x, y, z int
			fmt.Sscanf(lines[i], "%d,%d,%d", &x, &y, &z)
			data[scanner] = append(data[scanner], [3]int{x, y, z})
		}
	}
	return data
}

func distance(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

func main() {
	// dataMap := readInput()

	// 각 스캐너별로
	// // 3방향 정사영
	// // // +- 두방향 진행하다가
	// // // // 점이 12개 이상이 처음 되는 순간
	// // // // // 거리 조합 66개 이상씩 구해두기 : 6방향 각각 66개이상씩 세트

	// 모든 스캐너 서로 비교하면서
	// 66개 이상씩

	// 음... 이렇게 풀면 안됨.
}
