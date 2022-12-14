package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lelemita/adventofcode/common"
)

//go:embed input.txt
var input_day string

const maximum = 100_000

type Node struct {
	Size int
	Ls   []string
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	tooBigDirs := map[string]bool{}

	// read
	nodeMap := map[string]Node{}
	oldDir := "/"
	newDir := ""
	ls := []string{}
	for i := 0; i < len(lines); i++ {
		if n, _ := fmt.Sscanf(lines[i], "$ cd %s", &newDir); n > 0 {
			if oldDir != ".." {
				nodeMap[oldDir] = Node{
					Size: 0,
					Ls:   ls,
				}
			}
			oldDir = newDir
		} else if lines[i] == "$ ls" {
			ls = []string{}
		} else {
			var sizeStr, name string
			fmt.Sscanf(lines[i], "%s %s", &sizeStr, &name)
			size := 0
			if sizeStr != "dir" {
				fileSize, err := strconv.Atoi(sizeStr)
				common.PanicIfErr(err)
				size = fileSize
			}
			nodeMap[name] = Node{Size: size}
			ls = append(ls, name)
			if size >= maximum {
				tooBigDirs[oldDir] = true
			}
		}
	}

	// find
	for name, node := range nodeMap {

	}
	return 0
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

}

func main() {
	fmt.Println("--2022 day 07 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
