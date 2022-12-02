// part1 방식을 part2로 확장하려다 실패함... 왜?
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dot struct {
	X, Y int
}

type fold struct {
	ToLeft bool //false: toUp
	Num    int
}

func readInput(path string) (map[dot]bool, []fold) {
	dots := map[dot]bool{}
	folds := []fold{}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, ",")
		if len(arr) == 2 {
			x, _ := strconv.Atoi(arr[0])
			y, _ := strconv.Atoi(arr[1])
			dots[dot{X: x, Y: y}] = true
		}
		if strings.Contains(arr[0], "=") {
			fs := strings.Split(arr[0], "=")
			num, _ := strconv.Atoi(fs[1])
			folds = append(folds, fold{ToLeft: strings.Contains(fs[0], "x"), Num: num})
		}
	}
	return dots, folds
}

func getSize(dots map[dot]bool) (int, int, int, int) {
	xMin, yMin, xMax, yMax := 9999, 9999, -9999, -9999
	for d := range dots {
		if d.X > xMax {
			xMax = d.X
		}
		if d.X < xMin {
			xMin = d.X
		}
		if d.Y > yMax {
			yMax = d.Y
		}
		if d.Y < yMin {
			yMin = d.Y
		}
	}
	return xMin, yMin, xMax, yMax
}

func solution(dots map[dot]bool, folds []fold) int {
	for _, f := range folds {
		nextDots := map[dot]bool{}
		if f.ToLeft {
			for d := range dots {
				if d.X < f.Num {
					nextDots[dot{X: d.X + (2 * (f.Num - d.X)), Y: d.Y}] = true
				} else if d.X > f.Num {
					nextDots[dot{X: d.X, Y: d.Y}] = true
				}
			}
		} else {
			for d := range dots {
				if d.Y < f.Num {
					nextDots[dot{X: d.X, Y: d.Y}] = true
				} else if d.Y > f.Num {
					nextDots[dot{X: d.X, Y: d.Y - (2 * (d.Y - f.Num))}] = true
				}
			}
		}
		fmt.Println(f, len(nextDots), "------")
		dots = nextDots
	}

	draw(dots)
	return 0
}

func draw(dots map[dot]bool) {
	xMin, yMin, xMax, yMax := getSize(dots)
	fmt.Println(xMin, yMin, xMax, yMax)
	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if dots[dot{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	solution(readInput("example.txt"))
	solution(readInput("input.txt"))
}
