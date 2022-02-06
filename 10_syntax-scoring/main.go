// https://adventofcode.com/2021/day/9
package main

import (
	"bufio"
	"fmt"
	"os"
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

var score = map[byte]int{
	byte(')'): 3,
	byte('}'): 1197,
	byte(']'): 57,
	byte('>'): 25137,
}
var pair = map[byte]byte{
	byte(')'): byte('('),
	byte('}'): byte('{'),
	byte(']'): byte('['),
	byte('>'): byte('<'),
}

func part01(input [][]byte) int {
	result := 0
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
					// fmt.Printf("(%2d,%2d) %s %s\n", i, j, stack, []byte{p})
					break Loop
				}
				stack = stack[:n]
			default:
				stack = append(stack, p)
			}
		}
		if isCorrupted != 0 {
			result += score[isCorrupted]
		}
	}
	return result
}

func main() {
	// fmt.Println(byte('('), byte(')'))
	// fmt.Println(byte('<'), byte('>'))
	// fmt.Println(byte('['), byte(']'))
	// fmt.Println(byte('{'), byte('}'))

	input := readInput("example.txt")
	fmt.Println(part01(input))
	// fmt.Println(part02(input))
	input = readInput("input.txt")
	fmt.Println(part01(input))
	// fmt.Println(part02(input))
}
