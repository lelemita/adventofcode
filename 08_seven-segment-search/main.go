// https://adventofcode.com/2021/day/8
package main

import (
	"bufio"
	"fmt"
	"math"
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

func decode(code string, exs *string) int {
	switch len(code) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	}
	var one, four string
	for _, val := range strings.Split(*exs, " ") {
		if len(val) == 2 {
			one = val
		} else if len(val) == 4 {
			four = val
		}
	}
	if len(code) == 5 {
		if strings.Contains(code, string(one[0])) && strings.Contains(code, string(one[1])) {
			return 3
		}
		count := len(four)
		for _, f := range four {
			for _, x := range code {
				if x == f {
					count -= 1
					if count == 1 {
						return 5
					}
				}
			}
		}
		return 2
	}
	// len(code) == 6
	if strings.Contains(code, string(one[0])) != strings.Contains(code, string(one[1])) {
		return 6
	}
	count := len(four)
	for _, f := range four {
		for _, x := range code {
			if x == f {
				count -= 1
				if count == 0 {
					return 9
				}
			}
		}
	}
	return 0
}

func part02(exs, vals []string) int {
	sum := 0
	for i, arr := range vals {
		strs := strings.Split(arr, " ")
		result := 0
		for j, code := range strs {
			result += int(math.Pow10(len(strs)-1-j)) * decode(code, &exs[i])
		}
		sum += result
	}
	return sum
}

func main() {
	exs, vals := readInput("example.txt")
	fmt.Println(part01(exs, vals))
	fmt.Println(part02(exs, vals))

	exs, vals = readInput("input.txt")
	fmt.Println(part01(exs, vals))
	fmt.Println(part02(exs, vals))
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
