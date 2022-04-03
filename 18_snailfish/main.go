package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// explode 다하고 split...
// 정규식 사용하는 방식 참고: https://github.com/BorisLeMeec/adventofcode/blob/main/35/main.go
func main() {
	// 문자열 합치기
	// formula := makeFormula()

	fs := []string{
		"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
		"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]",
		"[[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]],[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]]",
		"[[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]],[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]]",
	}

	ans := []string{
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
	}

	for i, f := range fs {
		result := part01(f)
		fmt.Println(ans[i] == result)
		if ans[i] != result {
			fmt.Println(ans[i])
			fmt.Println(result)
		}
	}

	// 크기 계산
}

func part01(formula string) string {
	// 멈출때까지 반응
	var didSome bool
	for {
		didSome, formula = react(formula)
		if !didSome {
			formula, didSome = split(formula)
			if !didSome {
				break
			}
		}
	}
	return formula
}

func react(formula string) (bool, string) {
	count := 0
	for i, v := range formula {
		//explode
		if v == '[' {
			count += 1
		} else if v == ']' {
			count -= 1
		}
		if count >= 5 {
			formula = explode(i, formula)
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

	// fmt.Println(bigNum)
	formula = formula[:idx[0]] + fmt.Sprintf("[%d,%d]", bigNum/2, bigNum/2+bigNum%2) + formula[idx[1]:]
	return formula, true
}

func explode(idx int, formula string) string {
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
				front = front[:i] + fmt.Sprintf("%d", newNum) + front[i+len(target):]
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
