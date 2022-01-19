package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compare(strA, strB string) (int, int) {
	a, _ := strconv.Atoi(strA)
	b, _ := strconv.Atoi(strB)
	if a > b {
		return b, a
	}
	return a, b
}

func split(r rune) bool {
	return r == ' ' || r == ',' || r == '-' || r == '>'
}

func part01(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	counts := map[string]int{}
	for _, line := range lines {
		p := strings.FieldsFunc(line, split)
		if p[0] == p[2] {
			y, _ := strconv.Atoi(p[0])
			start, end := compare(p[1], p[3])
			for i := start; i <= end; i++ {
				counts[fmt.Sprintf("%d-%d", i, y)] += 1
			}
		} else if p[1] == p[3] {
			x, _ := strconv.Atoi(p[1])
			start, end := compare(p[0], p[2])
			for i := start; i <= end; i++ {
				counts[fmt.Sprintf("%d-%d", x, i)] += 1
			}
		}
	}
	count := 0
	for _, v := range counts {
		if v > 1 {
			count += 1
		}
	}
	fmt.Println(count)
}

func main() {
	part01("./example")
	part01("./input")
	// part02("./example")
	// part02("./input")
}
