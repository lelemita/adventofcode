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
	IsVertical bool
	Line       int
}

func readInput(path string) ([]dot, []fold) {
	dots := []dot{}
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
			dots = append(dots, dot{X: x, Y: y})
		}
		if strings.Contains(arr[0], "=") {
			fs := strings.Split(arr[0], "=")
			num, _ := strconv.Atoi(fs[1])
			folds = append(folds, fold{IsVertical: strings.Contains(fs[0], "x"), Line: num})
		}
	}
	return dots, folds
}

func part01(dots []dot, folds []fold) int {
	for _, v := range dots {
		fmt.Println(v)
	}
	for _, v := range folds {
		fmt.Println(v)
	}
	return 0
}

func main() {
	part01(readInput("example.txt"))
}
