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
	return compareInt(a, b)
}

func compareInt(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func isSameAbs(a, b int) bool {
	if a == b {
		return true
	} else if a+b == 0 {
		return true
	}
	return false
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

func part02(path string) {
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
		fmt.Println(line)
		strPs := strings.FieldsFunc(line, split)
		p := [4]int{}
		for i := 0; i < 4; i++ {
			p[i], _ = strconv.Atoi(strPs[i])
		}
		if p[0] == p[2] {
			start, end := compareInt(p[1], p[3])
			for i := start; i <= end; i++ {
				counts[fmt.Sprintf("%d,%d", i, p[0])] += 1
			}
		} else if p[1] == p[3] {
			start, end := compareInt(p[0], p[2])
			for i := start; i <= end; i++ {
				counts[fmt.Sprintf("%d,%d", p[1], i)] += 1
			}
		} else if isSameAbs(p[0]-p[2], p[1]-p[3]) {
			start, end := compareInt(p[1], p[3])
			dif := end - start
			dx := (p[3] - p[1]) / dif
			dy := (p[2] - p[0]) / dif
			fmt.Println("  ", dx, dy, dif)
			for i := 0; i <= end-start; i++ {
				counts[fmt.Sprintf("%d,%d", p[0]+(dx*i), p[1]+(dy*i))] += 1
				fmt.Printf("  %d,%d\n", p[0]+(dx*i), p[1]+(dy*i))
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
	// part01("./example")
	// part01("./input")
	part02("./example")
	// part02("./input")
}
