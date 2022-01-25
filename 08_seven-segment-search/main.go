// https://adventofcode.com/2021/day/8
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(path string) ([]string, []string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var digits []string
	var output []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " | ")
		digits = append(digits, strs[0])
		output = append(output, strs[1])
	}

	return digits, output
}

func part01(exs, vals []string) int {
	count := 0
	for _, nums := range vals {
		for _, n := range strings.Split(nums, " ") {
			l := len(n)
			if (l >= 2 && l <= 4) || l == 7 {
				count += 1
			}
		}
	}
	return count
}

func main() {
	exs, vals := readInput("example.txt")
	fmt.Println(part01(exs, vals))
	// 	fmt.Println(part02(example))

	exs, vals = readInput("input.txt")
	fmt.Println(part01(exs, vals))
	// 	fmt.Println(part02(puzzle))
}
