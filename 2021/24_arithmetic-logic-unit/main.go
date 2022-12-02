// 참고(시작은 비슷했다며...ㅋㅋ): https://github.com/womogenes/AoC-2021-Solutions/blob/main/day_24
// 참고: xlsx 파일
//// upWard: 26z + w + b
//// downWard: (w == (z%26 - a)) ? z/26 : 26(z/26) + w + b

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Step struct {
	numA, numB int
}

var steps [14]Step = [14]Step{
	{numA: 10, numB: 5},
	{numA: 13, numB: 9},
	{numA: 12, numB: 4},
	{numA: -12, numB: 4},
	{numA: 11, numB: 10},
	{numA: -13, numB: 14},
	{numA: -9, numB: 14},
	{numA: -12, numB: 12},
	{numA: 14, numB: 14},
	{numA: -9, numB: 14},
	{numA: 15, numB: 5},
	{numA: 11, numB: 10},
	{numA: -16, numB: 8},
	{numA: -2, numB: 15},
}

type Memory struct {
	W, Z int
}

func (m *Memory) upWard(s Step) {
	m.Z = 26*m.Z + m.W + s.numB
}

func (m *Memory) downWard(s Step) {
	m.Z = m.Z / 26
}

func hasZero(n int) bool {
	str := strconv.Itoa(n)
	return strings.Contains(str, "0")
}

func getIndexNum(num, idx int) int {
	pow := int(math.Pow10(idx))
	return (num % pow) / (pow / 10)
}

func printResult(res []int) {
	for _, v := range res {
		fmt.Print(v)
	}
	fmt.Println()
}

func calc(n int) []int {
	if hasZero(n) {
		return nil
	}
	m := Memory{W: 0, Z: 0}
	result := []int{}
	idx := 0
	for i := 0; i < 14; i++ {
		s := steps[i]
		if s.numA > 0 {
			m.W = getIndexNum(n, 7-idx)
			idx += 1
			m.upWard(s)
		} else {
			m.W = m.Z%26 + s.numA
			if m.W < 1 || m.W > 9 {
				break
			}
			m.downWard(s)
		}
		result = append(result, m.W)
	}
	if m.Z == 0 {
		return result
	}
	return nil
}

func part01() {
	for n := 9999999; n >= 1111111; n-- {
		if res := calc(n); res != nil {
			printResult(res)
			break
		}
	}
}

func part02() {
	for n := 1111111; n <= 9999999; n++ {
		if res := calc(n); res != nil {
			printResult(res)
			break
		}
	}
}

func main() {
	part01()
	part02()
}
