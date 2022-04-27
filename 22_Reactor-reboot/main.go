package main

import (
	"bufio"
	"fmt"
	"os"
)

var steps []step

type step struct {
	on     bool
	cuboid [3][2]int
}

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		var kind string
		cube := [3][2]int{{}, {}, {}}
		fmt.Sscanf(scan.Text(), "%s x=%d..%d,y=%d..%d,z=%d..%d",
			&kind, &cube[0][0], &cube[0][1], &cube[1][0], &cube[1][1], &cube[2][0], &cube[2][1])
		steps = append(steps, step{kind == "on", cube})
	}
}

var region map[axies]bool

type axies struct {
	x, y, z int
}

func part01(min, max int) {
	for _, step := range steps {
		fmt.Println(step.on)
		for i, v := range step.cuboid {
			fmt.Println("  ", i, v)
		}
	}
}

func main() {
	steps = []step{}
	region = map[axies]bool{}
	readFile("example.txt")
	part01(-50, 50)
}
