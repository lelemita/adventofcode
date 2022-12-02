// 문제: https://adventofcode.com/2021/day/18
// 실패1: explode 다하고 split...해야했음
// 참고1: 정규식 사용하는 것: https://github.com/BorisLeMeec/adventofcode/blob/main/35/main.go
package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var puzzle string

func main() {
	fs := strings.Split(puzzle, "\n")

	// PART 01
	formula := fs[0]
	for _, f := range fs[1:] {
		formula = fmt.Sprintf("[%s,%s]", formula, f)
		formula = react(formula)
	}
	fmt.Println("part01:", getMagnitue(formula))

	// PART 02
	maxMagn := 0
	for i := 0; i < len(fs); i++ {
		for j := 0; j < len(fs); j++ {
			if i == j {
				continue
			}
			result := getMagnitue(react(fmt.Sprintf("[%s,%s]", fs[i], fs[j])))
			if result > maxMagn {
				maxMagn = result
			}
		}
	}
	fmt.Println("part02:", maxMagn)
}

func getMagnitue(formula string) int {
	compile, _ := regexp.Compile(`([\[][0-9]*,[0-9]*[\]])`)
	for {
		if formula[0] != '[' {
			break
		}
		idx := compile.FindStringIndex(formula)
		pair := formula[idx[0]:idx[1]]
		comma := strings.IndexByte(pair, ',')
		frontNumStr := string(pair[1:comma])
		backNumStr := string(pair[comma+1 : len(pair)-1])
		frontNum, err := strconv.Atoi(frontNumStr)
		if err != nil {
			panic(err)
		}
		backNum, err := strconv.Atoi(backNumStr)
		if err != nil {
			panic(err)
		}
		num := (frontNum * 3) + (backNum * 2)
		formula = formula[:idx[0]] + fmt.Sprintf("%d", num) + formula[idx[1]:]
	}
	result, err := strconv.Atoi(formula)
	if err != nil {
		panic(err)
	}
	return result
}

func react(formula string) string {
	var didSome bool
	for {
		didSome, formula = explode(formula)
		if !didSome {
			formula, didSome = split(formula)
			if !didSome {
				break
			}
		}
	}
	return formula
}

func explode(formula string) (bool, string) {
	count := 0
	for i, v := range formula {
		if v == '[' {
			count += 1
		} else if v == ']' {
			count -= 1
		}
		if count >= 5 {
			formula = boom(i, formula)
			return true, formula
		}
	}
	return false, formula
}

func split(formula string) (string, bool) {
	compile, _ := regexp.Compile(`[1-9][0-9]+`)
	idx := compile.FindStringIndex(formula)
	if idx == nil {
		return formula, false
	}
	bigNum, err := strconv.Atoi(formula[idx[0]:idx[1]])
	if err != nil {
		panic(err)
	}

	formula = formula[:idx[0]] + fmt.Sprintf("[%d,%d]", bigNum/2, bigNum/2+bigNum%2) + formula[idx[1]:]
	return formula, true
}

func boom(idx int, formula string) string {
	// 정규식 쓰는거 ref: https://github.com/BorisLeMeec/adventofcode/blob/main/35/main.go
	compile, _ := regexp.Compile(`([\[][0-9]*,[0-9]*[\]])`)
	jdx := compile.FindStringIndex(formula[idx:])
	pair := formula[idx+jdx[0] : idx+jdx[1]]
	comma := strings.IndexByte(pair, ',')
	frontNumStr := string(pair[1:comma])
	backNumStr := string(pair[comma+1 : len(pair)-1])
	frontNum, err := strconv.Atoi(frontNumStr)
	backNum, err := strconv.Atoi(backNumStr)
	if err != nil {
		panic(err)
	}

	front := formula[:idx]
	target := []byte{}
	findNum := false
	for i := idx - 1; i >= 0; i-- {
		if byte('0') <= front[i] && front[i] <= byte('9') {
			findNum = true
			for j := i; j >= 0; j-- {
				if front[j] == '[' || front[j] == ',' || front[j] == ']' {
					break
				} else {
					target = append(target, front[j])
				}
			}
			if findNum {
				newNum := getReverseNum(target) + frontNum
				front = front[:i-len(target)+1] + fmt.Sprintf("%d", newNum) + front[i+1:]
				break
			}
		}
	}

	back := formula[idx+len(pair):]
	target = []byte{}
	findNum = false
	for i := 0; i < len(back); i++ {
		if byte('0') <= back[i] && back[i] <= byte('9') {
			findNum = true
			for j := i; j < len(back); j++ {
				if back[j] == '[' || back[j] == ',' || back[j] == ']' {
					break
				} else {
					target = append(target, back[j])
				}
			}
			if findNum {
				targetNum, err := strconv.Atoi(string(target))
				if err != nil {
					panic(err)
				}
				newNum := targetNum + backNum
				back = back[:i] + fmt.Sprintf("%d", newNum) + back[i+len(target):]
				break
			}
		}
	}

	formula = front + "0" + back
	return formula
}

func getReverseNum(victims []byte) int {
	rev := []byte{}
	for i := len(victims) - 1; i >= 0; i-- {
		rev = append(rev, victims[i])
	}
	num, err := strconv.Atoi(string(rev))
	if err != nil {
		panic(err)
	}
	return num
}
