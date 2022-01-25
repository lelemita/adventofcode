// https://adventofcode.com/2021/day/8
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func part02(exs, vals []string) int {
	count := 0
	if bi, err := strconv.ParseInt("001001", 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bi)
		fmt.Println(2 | bi)
	}

	// 1: 길이2
	// 7: 길이3
	// 4: 길이4
	// 8: 길이7
	// 길이5 중(2,3,5)에서
	//	 1성분 다 있으면: 3
	//	 없는 거 중에서
	//		4-X == 1 : 5
	//	 	4-X == 2 : 2
	// 길이6 중(6,9,0)에서
	//	 4-X == 0 : 9
	//	 4-X == 1 : 0
	//	 4-X == 3 : 6
	return count
}

func main() {
	exs, vals := readInput("example.txt")
	// fmt.Println(part01(exs, vals))
	fmt.Println(part02(exs, vals))

	// exs, vals = readInput("input.txt")
	// fmt.Println(part01(exs, vals))
	// fmt.Println(part02(exs, vals))
}
