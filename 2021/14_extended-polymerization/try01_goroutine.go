package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	// solution("example.txt")
	solution("input.txt")
	// solution("input.txt")
}

const STEP = 40
const UNIT_NUM = 10000000

var wg sync.WaitGroup

func solution(path string) {
	initial, rule := readInput(path)
	result := make(chan []byte)
	wg.Add(1)
	go reaction(STEP/2, initial, rule, result)
	polyes, jointMap := seprate(<-result)

	countMaps := []map[byte]int{}
	for i, v := range polyes {
		fmt.Println(i)
		res := make(chan []byte)
		wg.Add(1)
		go reaction(STEP/2, v, rule, res)
		countMaps = append(countMaps, countAtom(<-res))
	}
	wg.Wait()

	totalCountMap := sumCounts(countMaps, jointMap)
	fmt.Println(getMaxGap(totalCountMap))
}

func sumCounts(countMaps []map[byte]int, jointMap map[byte]int) map[byte]int {
	total := map[byte]int{}
	for _, m := range countMaps {
		for k, v := range m {
			total[k] += v
		}
	}
	for k, v := range jointMap {
		total[k] -= v
	}
	return total
}

func reaction(step int, polymer []byte, rule map[string]byte, done chan<- []byte) {
	for i := 0; i < step; i++ {
		next := []byte{polymer[0]}
		for j := 0; j < len(polymer)-1; j++ {
			a := polymer[j]
			c := polymer[j+1]
			b, isExist := rule[string([]byte{a, c})]
			if isExist {
				next = append(next, b)
			}
			next = append(next, c)
		}
		polymer = next
	}
	done <- polymer
	wg.Done()
}

func readInput(path string) (initial []byte, rule map[string]byte) {
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

	rule = map[string]byte{}
	for scan.Scan() {
		var a, b, c byte
		fmt.Sscanf(scan.Text(), "%c%c -> %c", &a, &c, &b)
		rule[string([]byte{a, c})] = b
	}
	return
}

func seprate(polymer []byte) ([][]byte, map[byte]int) {
	polyes := [][]byte{}
	unit := len(polymer) / UNIT_NUM
	cnt := 0
	for i := 0; i < UNIT_NUM; i++ {
		polyes = append(polyes, polymer[unit*i:unit*(i+1)])
		cnt += 1
	}
	polyes = append(polyes, polymer[cnt*unit:])

	joints := [][]byte{}
	jointMap := map[byte]int{}
	for i := 0; i < len(polyes)-1; i++ {
		jointFront := polyes[i][len(polyes[i])-1]
		jointBack := polyes[i+1][0]
		joints = append(joints, []byte{jointFront, jointBack})
		jointMap[jointFront] += 1
		jointMap[jointBack] += 1
	}
	for _, v := range joints {
		polyes = append(polyes, v)
	}
	return polyes, jointMap
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
	min := 9999999999
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
