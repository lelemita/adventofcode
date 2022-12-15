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

var tooBigDirs = map[string]bool{}
var nodeMap = map[string]Node{}

func read(input string) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	lines = append(lines, "$ cd end")

	// read
	oldDir := ""
	newDir := ""
	ls := []string{}
	lsSum := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "$ ls" {
			continue
		} else if n, _ := fmt.Sscanf(lines[i], "$ cd %s", &newDir); n > 0 {
			if _, isExist := nodeMap[oldDir]; !isExist {
				nodeMap[oldDir] = Node{
					Size: 0,
					Ls:   ls,
				}
				if lsSum > maximum {
					tooBigDirs[oldDir] = true
				}
				ls = []string{}
				lsSum = 0
			}

			if newDir == ".." {
				idx := strings.LastIndex(oldDir, "/")
				oldDir = oldDir[:idx]
			} else {
				oldDir = fmt.Sprintf("%s/%s", oldDir, newDir)
			}
		} else {
			var sizeStr, name string
			fmt.Sscanf(lines[i], "%s %s", &sizeStr, &name)
			name = fmt.Sprintf("%s/%s", oldDir, name)
			ls = append(ls, name)
			if sizeStr != "dir" {
				size, err := strconv.Atoi(sizeStr)
				common.PanicIfErr(err)
				nodeMap[name] = Node{Size: size}
				lsSum += size
			}
		}
	}
}

func Part1(input string) int {
	read(input)

	// find
	result := 0
	for name, node := range nodeMap {
		if node.Size == 0 {
			oneDirSize := getSize(name)
			if oneDirSize <= maximum {
				result += oneDirSize
			}
		}
	}
	return result
}

func getSize(name string) int {
	node := nodeMap[name]
	result := node.Size
	for _, nm := range node.Ls {
		result += getSize(nm)
	}
	return result
}

func Part2(input string) int {
	read(input)
	need := getSize("//") - 40_000_000

	result := 70_000_000
	for name, node := range nodeMap {
		if node.Size == 0 {
			oneDirSize := getSize(name)
			if oneDirSize >= need {
				if oneDirSize < result {
					result = oneDirSize
				}
			}
		}
	}
	return result
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
