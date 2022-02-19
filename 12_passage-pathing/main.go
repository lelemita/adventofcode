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

func readInput(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "-")
		caves[arr[0]] = append(caves[arr[0]], arr[1])
		caves[arr[1]] = append(caves[arr[1]], arr[0])
	}
}

func part01() int {
	count = 0
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
	// readInput("./12_passage-pathing/ex01.txt")
	// readInput("./12_passage-pathing/ex02.txt")
	// readInput("./12_passage-pathing/ex03.txt")
	readInput("./12_passage-pathing/input.txt")
	fmt.Println(part01())
	// fmt.Println(part02(input))

	// fmt.Println(part01(input))
	// fmt.Println(part02(input))
}
