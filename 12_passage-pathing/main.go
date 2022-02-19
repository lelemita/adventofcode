// https://adventofcode.com/2021/day/12
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var partNum = 1
var count = 0
var caves = map[string][]string{}

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

func part(pa int, input map[string][]string) int {
	partNum = pa
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
	if next == "start" {
		return false
	}
	if []byte(next)[0] > 'Z' {
		haveTwice := true
		if partNum == 2 {
			haveTwice = chkTwice(route)
		}
		for _, v := range route {
			if next == v && haveTwice {
				return false
			}
		}
	}
	return true
}

func chkTwice(route []string) bool {
	memo := map[string]bool{}
	for _, v := range route {
		if []byte(v)[0] > 'Z' {
			_, exist := memo[v]
			if exist {
				return true
			}
			memo[v] = true
		}
	}
	return false
}

func main() {
	fmt.Println(part(1, readInput("ex01.txt")) == 10)
	fmt.Println(part(1, readInput("ex02.txt")) == 19)
	fmt.Println(part(1, readInput("ex03.txt")) == 226)
	fmt.Println(part(1, readInput("input.txt")) == 4912)

	fmt.Println(part(2, readInput("ex01.txt")) == 36)
	fmt.Println(part(2, readInput("ex02.txt")) == 103)
	fmt.Println(part(2, readInput("ex03.txt")) == 3509)
	fmt.Println(part(2, readInput("input.txt")) == 150004)
}
