// 미적으로 풀면 안됨..40.5.....디지털 세계임..ㅠㅠ
// https://gist.github.com/mdarrik/72835482b47e9b3e2827faa5789f8e6a
package main

import (
	"fmt"
)

var example = "target area: x=20..30, y=-10..-5"
var puzzle = "target area: x=85..145, y=-163..-108"

func part01(y1, y2 int) {
	miny := -y1
	if y1 > y2 {
		miny = -y2
	}
	result := miny * (miny - 1) / 2
	fmt.Println(result)
}

func main() {
	part01(-10, -5)
	part01(-163, -108)
}
