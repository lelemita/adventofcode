// https://adventofcode.com/2021/day/15
// 1차: Dynamic Programming으로만 했다가 틀림.. 다시 올라가는 경로 반영 안됨
// 2차: Dijkstra Algorithm
package main

import (
	"bufio"
	priorityQueue "chiton/priority"
	"container/heap"
	"errors"
	"fmt"
	"math"
	"os"
)

var cave [][]int

func main() {
	// cave = readInput("input.txt")
	cave = readInput("example.txt")
	// cave = readInput("example2.txt") // for part01
	// cave = readInput("example3.txt") // for part02

	solution()
	cave = fiveTimes() // for part02
	solution()
}

func solution() {
	idx := 0
	pq := priorityQueue.PriorityQueue{&priorityQueue.Item{
		Axis:  "0,0",
		Cost:  0,
		Index: idx,
	}}
	heap.Init(&pq)

	cost := map[string]int{"0,0": 0}
	for pq.Len() > 0 {
		here := heap.Pop(&pq).(*priorityQueue.Item)
		old, exist := cost[here.Axis]
		if exist && old < here.Cost {
			continue
		}
		var i, j int
		_, err := fmt.Sscanf(here.Axis, "%d,%d", &i, &j)
		if err != nil {
			panic(err)
		}

		// 인접한 정점들을 모두 검사한다.
		for x := 0; x < 2; x++ {
			for d := -1; d <= 1; d += 2 {
				var err error
				var nextCost int
				var there string
				if x == 0 {
					nextCost, err = getNextCost(here.Cost, i+d, j)
					there = fmt.Sprintf("%d,%d", i+d, j)
				} else {
					nextCost, err = getNextCost(here.Cost, i, j+d)
					there = fmt.Sprintf("%d,%d", i, j+d)
				}
				if err != nil {
					continue
				}
				old, exist := cost[there]
				if !exist || old > nextCost {
					cost[there] = nextCost
					idx += 1
					pq.Push(&priorityQueue.Item{
						Axis:  there,
						Cost:  nextCost,
						Index: idx,
					})
					heap.Init(&pq)
				}
			}
		}
	}
	goalAxis := fmt.Sprintf("%d,%d", len(cave)-1, len(cave[0])-1)
	fmt.Println(cost[goalAxis])
}

func getNextCost(cost, i, j int) (int, error) {
	if i < 0 || j < 0 || i >= len(cave) || j >= len(cave[0]) {
		return math.MaxInt, errors.New("out")
	}
	return cave[i][j] + cost, nil
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

func fiveTimes() [][]int {
	realCave := [][]int{}

	// 가로로 5회 반복
	for _, row := range cave {
		temp := []int{}
		for i := 0; i < 5; i++ {
			for _, n := range row {
				temp = append(temp, higher(n, i))
			}
		}
		realCave = append(realCave, temp)
	}
	// 세로로 5회 반복
	for i := 1; i < 5; i++ {
		for r := 0; r < len(cave); r++ {
			temp := []int{}
			for _, n := range realCave[r] {
				temp = append(temp, higher(n, i))
			}
			realCave = append(realCave, temp)
		}
	}
	return realCave
}

func higher(n, times int) int {
	for i := 0; i < times; i++ {
		n = next(n)
	}
	return n
}

func next(n int) int {
	if n == 9 {
		return 1
	}
	return n + 1
}
