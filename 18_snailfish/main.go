package main

import (
	"fmt"
)

func main() {
	// 문자열 합치기
	// formula := makeFormula()

	fs := []string{
		"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",

		"[[[7,7],[[[3,7],[4,3]],[[6,3],[8,8]]]]]",

		"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]",
		"[[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]],[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]]",
		"[[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]],[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]]",
	}

	ans := []string{
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",

		"[[8,[7,7]],[[7,9],[5,0]]]",

		"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		"[[[[7,0],[7,7]],[[7,7],[7,8]]],[[[7,7],[8,8]],[[7,7],[8,7]]]]",
	}

	for i, f := range fs {
		result := part01(f)
		fmt.Println(ans[i] == result)
		fmt.Println(ans[i])
		fmt.Println(result)
	}

	// 크기 계산
}

func part01(formula string) string {
	// 멈출때까지 반응
	isEnd := false
	for !isEnd {
		isEnd, formula = react(formula)
	}
	return formula
}

func react(formula string) (bool, string) {
	count := 0
	for i, v := range formula {
		if v == '[' {
			count += 1
		} else if v == ']' {
			count -= 1
		}
		if count >= 5 {
			if checkOneDepthPair(formula[i : i+5]) {
				formula = explode(i, formula)
				return false, formula
			}
		}
	}
	return true, formula
}

func checkOneDepthPair(str string) bool {
	if str[0] != '[' {
		return false
	}
	if byte('0') > str[1] || str[1] > byte('9') {
		return false
	}
	if str[2] != ',' {
		return false
	}
	if byte('0') > str[3] || str[3] > byte('9') {
		return false
	}
	if str[4] != ']' {
		return false
	}
	return true
}

func explode(idx int, formula string) string {
	// explode, split
	pair := formula[idx : idx+5]
	frontNum := int(pair[1] - byte('0'))
	backNum := int(pair[3] - byte('0'))

	front := formula[:idx]
	for i := len(front) - 1; i >= 0; i-- {
		if byte('0') <= front[i] && front[i] <= byte('9') {
			num := int(front[i]-byte('0')) + frontNum
			results := ""
			if num >= 10 {
				results = fmt.Sprintf("[%d,%d]", num/2, num/2+num%2)
			} else {
				results = fmt.Sprintf("%d", num)
			}
			front = front[:i] + results + front[i+1:]
			break
		}
	}

	back := formula[idx+5:]
	for i := 0; i < len(back); i++ {
		if byte('0') <= back[i] && back[i] <= byte('9') {
			num := int(back[i]-byte('0')) + backNum
			results := ""
			if num >= 10 {
				results = fmt.Sprintf("[%d,%d]", num/2, num/2+num%2)
			} else {
				results = fmt.Sprintf("%d", num)
			}
			back = back[:i] + results + back[i+1:]
			break
		}
	}

	// fmt.Println(front, "0", back)
	formula = front + "0" + back
	return formula
}
