// https://adventofcode.com/2021/day/10
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

var pair = map[byte]byte{
	byte(')'): byte('('),
	byte('}'): byte('{'),
	byte(']'): byte('['),
	byte('>'): byte('<'),
}

func part01(input [][]byte) ([][]byte, int) {
	var score = map[byte]int{
		byte(')'): 3,
		byte(']'): 57,
		byte('}'): 1197,
		byte('>'): 25137,
	}

	result := 0
	after := [][]byte{}
	for _, line := range input {
		stack := []byte{}
		isCorrupted := byte(0)
	Loop:
		for _, p := range line {
			switch p {
			case byte(')'), byte('>'), byte(']'), byte('}'):
				n := len(stack) - 1
				if stack[n] != pair[p] {
					isCorrupted = p
					break Loop
				}
				stack = stack[:n]
			default:
				stack = append(stack, p)
			}
		}
		if isCorrupted != 0 {
			result += score[isCorrupted]
		} else {
			after = append(after, line)
		}
	}
	return after, result
}

func part02(input [][]byte) int {
	var score = map[byte]int{
		byte('('): 1,
		byte('['): 2,
		byte('{'): 3,
		byte('<'): 4,
	}

	results := []int{}
	for _, line := range input {
		stack := []byte{}
		for _, p := range line {
			switch p {
			case byte(')'), byte('>'), byte(']'), byte('}'):
				n := len(stack) - 1
				stack = stack[:n]
			default:
				stack = append(stack, p)
			}
		}
		res := 0
		for i := len(stack) - 1; i >= 0; i-- {
			res *= 5
			res += score[stack[i]]
		}
		// fmt.Printf("%d %s \n", res, stack)
		results = append(results, res)
	}
	sort.Ints(results)
	return results[len(input)/2]
}

func main() {
	input := readInput("example.txt")
	input, res01 := part01(input)
	fmt.Println(res01)
	fmt.Println(part02(input))

	input = readInput("input.txt")
	input, res01 = part01(input)
	fmt.Println(res01)
	fmt.Println(part02(input))
}
