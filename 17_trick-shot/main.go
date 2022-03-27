package main

import (
	"fmt"
)

var example = "target area: x=20..30, y=-10..-5"
var puzzle = "target area: x=85..145, y=-163..-108"

func main() {
	part01(-10)
	part01(-163)
	part02(20, 30, -10, -5)
	part02(85, 145, -163, -108)
}

func part01(miny int) {
	// 미적으로 풀면 안됨..40.5.....디지털 세계임..ㅠㅠ
	// https://gist.github.com/mdarrik/72835482b47e9b3e2827faa5789f8e6a
	result := miny * (miny + 1) / 2
	fmt.Println(result)
}

func part02(x1, x2, y1, y2 int) {
	count := map[string]bool{}
	maxVyzero := -y1 - 1
	maxStep := -y1*2 + 1

	for vy := y1; vy <= maxVyzero; vy++ {
		for vx := 0; vx <= x2; vx++ {
			for t := 0; t <= maxStep; t++ {
				x, y := getLoc(t, vx, vy)
				if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
					count[fmt.Sprintf("%d,%d", vx, vy)] = true
				}
				if x > x2 || y < y1 {
					break
				}
			}
		}
	}
	fmt.Println(len(count))
}

func getLoc(step, vxzero, vyzero int) (int, int) {
	return getX(vxzero, step), getY(vyzero, step)
}

func getX(vxzero, step int) int {
	x := 0
	for t := 0; t < step; t++ {
		x += getVx(vxzero, t)
	}
	return x
}

func getY(vyzero, step int) int {
	y := 0
	for t := 0; t < step; t++ {
		y += getVy(vyzero, t)
	}
	return y
}

func getVx(vxzero, step int) int {
	if step >= vxzero {
		return 0
	}
	return vxzero - step
}

func getVy(vyzero, step int) int {
	return vyzero - step
}
