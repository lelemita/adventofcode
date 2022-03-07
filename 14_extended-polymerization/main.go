package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// solution("example.txt", 10)
	// solution("input.txt", 10)
	solution("input.txt", 40)
}

var RULE map[string]byte

func solution(path string, totalStep int) {
	initial := readInput(path)
	// 초기 화합물에 각 단위가 얼마나 들었는지 확인
	fmt.Println("step_01")
	unitCount := getUnitCount(initial)

	// 각 단위가 어떻게 변하는지 확인
	fmt.Println("step_02")
	unitResult := getUnitResult(totalStep)

	// unit 별 변화 반영
	fmt.Println("step_03")
	totalResult := countPerUnit(unitCount, unitResult)

	// 중복 값 제거
	fmt.Println("step_04")
	countMap := countAtom(initial)
	for atom, cnt := range countMap {
		totalResult[atom] -= cnt
	}
	totalResult[initial[0]] += 1
	totalResult[initial[len(initial)-1]] += 1

	fmt.Println(getMaxGap(totalResult))
}

func countPerUnit(unitCount map[string]int, unitResult map[string]map[byte]int) map[byte]int {
	totalResult := map[byte]int{}
	for unit, uCnt := range unitCount {
		for u, atomMap := range unitResult {
			if u == unit {
				for atom, aCnt := range atomMap {
					totalResult[atom] += (uCnt * aCnt)
				}
			}
		}
	}
	return totalResult
}

func getUnitCount(polymer []byte) map[string]int {
	unitCount := map[string]int{}
	for i := 0; i < len(polymer)-1; i++ {
		unit := string([]byte{polymer[i], polymer[i+1]})
		unitCount[unit] += 1
	}
	return unitCount
}

func getUnitResult(totalStep int) map[string]map[byte]int {
	unitResult := map[string]map[byte]int{}
	for strUnit := range RULE {
		unit := []byte(strUnit)
		for i := 0; i < totalStep; i++ {
			unit = step(unit)
		}
		unitResult[strUnit] = countAtom(unit)
	}
	return unitResult
}

func step(polymer []byte) []byte {
	next := []byte{polymer[0]}
	for j := 0; j < len(polymer)-1; j++ {
		a := polymer[j]
		c := polymer[j+1]
		b, isExist := RULE[string([]byte{a, c})]
		if isExist {
			next = append(next, b)
		}
		next = append(next, c)
	}
	return next
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
