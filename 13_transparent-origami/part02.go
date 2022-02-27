package main

import (
	"bufio"
	"fmt"
	"os"
)

func sol(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dots := map[string]bool{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		dots[sc.Text()] = true
	}

	for sc.Scan() {
		var axis rune
		var foldLine int
		fmt.Sscanf(sc.Text(), "fold along %c=%d", &axis, &foldLine)

		for dot := range dots {
			var x, y int
			fmt.Sscanf(dot, "%d,%d", &x, &y)
			if axis == 'y' && y >= foldLine {
				y = 2*foldLine - y
			} else if axis == 'x' && x >= foldLine {
				x = 2*foldLine - x
			}
			dots[fmt.Sprintf("%d,%d", x, y)] = true
		}
	}
	draw(dots, 50, 10)
}

func draw(dots map[string]bool, xMax, yMax int) {
	for y := 0; y <= yMax; y++ {
		for x := 0; x <= xMax; x++ {
			if dots[fmt.Sprintf("%d,%d", x, y)] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	sol("example.txt")
	fmt.Println("------------------")
	sol("input.txt")
}
