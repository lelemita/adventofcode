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

func part01() {
	for n := 9999999; n >= 1111111; n-- {
		if hasZero(n) {
			continue
		}
		result := []int{}
		m := Memory{W: 0, Z: 0}

		for i := 14; i > 0; i-- {

			이렇게 하면 안됨. 
			downward 는 정해진 숫자임
			그걸 이용하면 9^7 경우만 확인하면 됨

			pow := int(math.Pow10(i))
			m.W = (n % pow) / (pow / 10)

			s := steps[14-i]
			if s.numA > 0 {
				m.upWard(s)
				result = append(result, m.W)
			} else {
				if m.Z%26 != s.numA {
					break
				}
				m.downWard(s)
			}
		}
		if m.Z == 0 {
			fmt.Println(n)
			break
		}
	}
}

func hasZero(n int) bool {
	str := strconv.Itoa(n)
	return strings.Contains(str, "0")
}

func main() {
	for i := 0; i <= 26; i++ {

		fmt.Println(i % 26)
	}
	// part01()
}
