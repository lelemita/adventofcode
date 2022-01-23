// https://adventofcode.com/2021/day/7
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func part01(puzzle string) int {
	lines := strings.Split(puzzle, ",")
	crabs := make([]int, len(lines))
	for i, line := range lines {
		dis, _ := strconv.Atoi(line)
		crabs[i] = dis
	}
	sort.Ints(crabs)
	middle := crabs[len(crabs)/2]
	dists := make([]int, len(crabs))
	for c := 0; c < len(crabs); c++ {
		if crabs[c] > middle {
			dists[c] = crabs[c] - middle
		} else {
			dists[c] = middle - crabs[c]
		}

	}

	sum := 0
	for _, v := range dists {
		sum += v
	}
	return sum
}

func main() {
	example := readInput("example.txt")
	fmt.Println(part01(example))

	puzzle := readInput("input.txt")
	fmt.Println(part01(puzzle))

}
