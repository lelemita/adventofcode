// https://adventofcode.com/2021/day/12
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var caves = map[string][]string{}
var count = 0

func readInput(path string) map[string][]string {
	var input = map[string][]string{}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "-")
		input[arr[0]] = append(input[arr[0]], arr[1])
		input[arr[1]] = append(input[arr[1]], arr[0])
	}
	return input
}

func part(input map[string][]string) int {
	count = 0
	caves = input
	pass([]string{}, "start")
	return count
}

func pass(route []string, here string) {
	route = append(route, here)
	if here == "end" {
		count += 1
		return
	}
	for _, next := range caves[here] {
		if chkRoute(route, next) {
			pass(route, next)
		}
	}
}

func chkRoute(route []string, next string) bool {
	if []byte(next)[0] > 'Z' {
		for _, v := range route {
			if next == v {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(part(readInput("./12_passage-pathing/ex01.txt")) == 10)
	fmt.Println(part(readInput("./12_passage-pathing/ex02.txt")) == 19)
	fmt.Println(part(readInput("./12_passage-pathing/ex03.txt")) == 226)
	fmt.Println(part(readInput("./12_passage-pathing/input.txt")))

	// fmt.Println(part02(readInput("./12_passage-pathing/ex01.txt")) == 36)
	// 	fmt.Println(part02(readInput("./12_passage-pathing/ex02.txt")) == 103)
	// 	fmt.Println(part02(readInput("./12_passage-pathing/ex03.txt")) == 3509)
	// 	fmt.Println(part02(readInput("./12_passage-pathing/input.txt")))
}
