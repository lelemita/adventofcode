package main

import (
	"bufio"
	"fmt"
	"os"
)

var steps []step

const (
	X = 0
	Y = 1
	Z = 2
)

type step struct {
	on     bool
	cuboid [3][2]int
}

func readFile(path string) {
	steps = []step{}
	region = map[axies]bool{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		var kind string
		cu := [3][2]int{{}, {}, {}}
		fmt.Sscanf(scan.Text(), "%s x=%d..%d,y=%d..%d,z=%d..%d",
			&kind, &cu[X][0], &cu[X][1], &cu[Y][0], &cu[Y][1], &cu[Z][0], &cu[Z][1])
		steps = append(steps, step{kind == "on", cu})
	}
}

var region map[axies]bool

type axies struct {
	x, y, z int
}

func part01(min, max int) {
	for _, s := range steps {
		if !chkStep(min, max, s.cuboid) {
			continue
		}
		for i := s.cuboid[X][0]; i <= s.cuboid[X][1]; i++ {
			for j := s.cuboid[Y][0]; j <= s.cuboid[Y][1]; j++ {
				for k := s.cuboid[Z][0]; k <= s.cuboid[Z][1]; k++ {
					region[axies{i, j, k}] = s.on
				}
			}
		}
	}

	fmt.Println("Part01: ", getOnCount(min, max))
}

func chkStep(min, max int, c [3][2]int) bool {
	for _, x := range [3]int{X, Y, Z} {
		if c[x][0] < min && c[x][1] < min {
			return false
		} else if c[x][0] > max && c[x][1] > max {
			return false
		}
	}
	return true
}

func getOnCount(min, max int) int {
	count := 0
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			for k := min; k <= max; k++ {
				if val, isExist := region[axies{i, j, k}]; isExist && val {
					count += 1
				}
			}
		}
	}
	return count
}

func main() {
	// readFile("example1.txt")
	// part01(-50, 50)

	// readFile("example.txt")
	// part01(-50, 50)

	readFile("input.txt")
	part01(-50, 50)
}
