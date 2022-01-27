// https://adventofcode.com/2021/day/9
package main

import (
	"bufio"
	"fmt"
	"os"
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

func part01(input [][]byte) int {
	for i, line := range input {
		for j, p := range line {
			for 
		}
	}
	return 0
}

func main() {
	input := readInput("example.txt")
	fmt.Println(part01(input))
	// fmt.Println(part02(input))
	// input = readInput("input.txt")
	// fmt.Println(part01(input))
	// fmt.Println(part02(input))
}
