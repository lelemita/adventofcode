package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solution("example.txt")
	solution("input.txt")
}

func solution(path string) {
	polymer, rule := readInput(path)
	for i := 0; i < 10; i++ {
		next := []byte{polymer[0]}
		for j := 0; j < len(polymer)-1; j++ {
			a := polymer[j]
			c := polymer[j+1]
			b, isExist := rule[string([]byte{a, c})]
			if isExist {
				next = append(next, b)
			}
			next = append(next, c)
		}
		polymer = next
	}
	countMap := countAtom(polymer)
	fmt.Println(getMaxGap(countMap))
}

func readInput(path string) (initial []byte, rule map[string]byte) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	if scan.Scan() {
		initial = scan.Bytes()
		scan.Scan()
	} else {
		panic("no input")
	}

	rule = map[string]byte{}
	for scan.Scan() {
		var a, b, c byte
		fmt.Sscanf(scan.Text(), "%c%c -> %c", &a, &c, &b)
		rule[string([]byte{a, c})] = b
	}
	return
}

func countAtom(polymer []byte) map[byte]int {
	countMap := map[byte]int{}
	for _, v := range polymer {
		countMap[v] += 1
	}
	return countMap
}

func getMaxGap(countMap map[byte]int) int {
	max := 0
	min := 9999999999
	for _, v := range countMap {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}
