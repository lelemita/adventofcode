package main

import (
	"bufio"
	"fmt"
	"os"
)

var RULE map[string]byte
var OVERLAP map[byte]int

func main() {
	OVERLAP = map[byte]int{}
	initial := readInput("example.txt")
	solution(initial)
}

func solution(initial []byte) {
	// 초기 화합물에 각 단위가 얼마나 들었는지 확인
	unitCount := getUnitCount(initial)

	// 각 단위가 어떻게 변하는지 확인
	unitResult := getUnitResult()
	// for unit, m := range unitResult {
	// 	fmt.Println("-- ", unit, " ---")
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	// 각 단위별 원자 개수 결과 구하기
	unitResultAtom := countResultAtom(unitResult)
	// for unit, m := range unitResultAtom {
	// 	fmt.Println("-- ", unit, " ---")
	// 	for k, v := range m {
	// 		fmt.Println(string([]byte{k}), v)
	// 	}
	// }

	// unit 별 변화 반영
	totalResult := countPerUnit(unitCount, unitResultAtom)
	// for k, v := range totalResult {
	// 	fmt.Println(string([]byte{k}), v)
	// }

	// 중복 값 제거
	// countMap := countAtom(initial)
	for atom, cnt := range OVERLAP {
		totalResult[atom] -= cnt
	}
	totalResult[initial[0]] += 1
	totalResult[initial[len(initial)-1]] += 1

	fmt.Println(getMaxGap(totalResult))
}

func countResultAtom(unitResult map[string]map[string]int) map[string]map[byte]int {
	result := map[string]map[byte]int{}
	for unit, unitMap := range unitResult {
		result[unit] = map[byte]int{}
		for poly, cnt := range unitMap {
			result[unit][poly[0]] += cnt
			result[unit][poly[1]] += cnt
		}
	}
	return result
}

func getUnitResult() map[string]map[string]int {
	// 유닛별10번돈결과(unit) map[유닛]int
	// for 각 유닛
	//// 각 유닛을 10번 진행시켰을 때 polymer 구하기: afterTen(unit) []byte
	//// 그 폴리머에서 유닛 개수 구하기: map[유닛]int

	// 유닛별20번돈결과(unit) map[유닛]int
	// for 각 유닛
	//// 유닛A.유닛a += 10[A] * 10[a]
	map10 := map[string]map[string]int{}
	for strUnit := range RULE {
		map10[strUnit] = afterTen([]byte(strUnit))
	}

	// unitResult := map[string]map[string]int{}
	// for key, val := range map10 {
	// 	unitResult[key] = map[string]int{}
	// 	for k, v := range val {
	// 		unitResult[key][k] = v
	// 	}
	// }
	// for i := 0; i < 1; i++ {
	// 	for b := range map10 {
	// 		unitResult[b] = afterTwenty(b, map10, unitResult)
	// 	}
	// }
	// return unitResult
	return map10
}

func afterTwenty(host string, map10, unitResult map[string]map[string]int) map[string]int {
	result := map[string]int{}
	for part, cntPart := range unitResult[host] {
		for sub, cntSub := range map10[part] {
			result[sub] += cntPart * cntSub
		}
	}
	return result
}

func afterTen(polymer []byte) map[string]int {
	for i := 0; i < 4; i++ {
		saveOverlap(polymer)

		polymer = step(polymer)
	}
	return getUnitCount(polymer)
}

func saveOverlap(polymer []byte) {
	flag := string(polymer) == "CC"
	if flag {
		fmt.Println(string(polymer))
	}
	countMap := countAtom(polymer)
	for atom, cnt := range countMap {
		OVERLAP[atom] -= cnt
	}
	OVERLAP[polymer[0]] += 1
	OVERLAP[polymer[len(polymer)-1]] += 1
}

func getUnitCount(polymer []byte) map[string]int {
	unitCount := map[string]int{}
	for i := 0; i < len(polymer)-1; i++ {
		unit := string([]byte{polymer[i], polymer[i+1]})
		unitCount[unit] += 1
	}
	return unitCount
}

func step(polymer []byte) []byte {
	next := []byte{}
	for j := 0; j < len(polymer)-1; j++ {
		a := polymer[j]
		c := polymer[j+1]
		next = append(next, a)
		b, isExist := RULE[string([]byte{a, c})]
		if isExist {
			next = append(next, b)
		}
		next = append(next, c)
	}
	return next
}

func countPerUnit(unitCount map[string]int, unitResultAtom map[string]map[byte]int) map[byte]int {
	totalResult := map[byte]int{}
	for unit, uCnt := range unitCount {
		for u, atomMap := range unitResultAtom {
			if u == unit {
				for atom, aCnt := range atomMap {
					totalResult[atom] += (uCnt * aCnt)
				}
			}
		}
	}
	return totalResult
}

func readInput(path string) (initial []byte) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	if scan.Scan() {
		initial = scan.Bytes()
		scan.Scan()
	} else {
		panic("no input")
	}

	RULE = map[string]byte{}
	for scan.Scan() {
		var a, b, c byte
		fmt.Sscanf(scan.Text(), "%c%c -> %c", &a, &c, &b)
		RULE[string([]byte{a, c})] = b
	}
	return
}

func countAtom(polymer []byte) map[byte]int {
	countMap := map[byte]int{}
	for _, v := range polymer {
		countMap[v] += 1
	}
	return countMap
}

func getMaxGap(countMap map[byte]int) int {
	max := 0
	min := 9999999
	for _, v := range countMap {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}
