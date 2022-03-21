// reference: https://getsturdy.com/advent-of-code-2021-uoeIDQk/browse/day16/gustav/main.go
package main

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var contents string

func main() {
	contents = strings.TrimSpace(contents)
	bin := hexBin(contents)
	// for {
	// 	if len(bin)%4 == 0 {
	// 		break
	// 	}
	// 	// 왜 붙이는 거지?
	// 	bin = "0" + bin
	// }
	_, sum, val := parse(bin, 0)
	fmt.Println(sum, val)
}

// pc: 현위치
func parse(bin string, pc int) (retPc int, sumVersion, ret int64) {
	version := binDec(bin[pc+0 : pc+0+3])
	pkgType := binDec(bin[pc+3 : pc+3+3])

	// Literal
	if pkgType == 4 {
		pc += 6
		var by string
		for {
			last := bin[pc] == '0'
			by += bin[pc+1 : pc+5]
			pc += 5
			if last {
				break
			}
		}
		return pc, version, binDec(by)
	}

	lengthType := bin[pc+6]
	var subVals []int64
	if lengthType == '0' {
		l := binDec(bin[pc+7 : pc+7+15])
		pc = pc + 7 + 15
		stopAt := pc + int(l)
		for {
			nextpc, addV, val := parse(bin, pc)
			version += addV
			pc = nextpc
			subVals = append(subVals, val)
			if pc >= stopAt {
				break
			}
		}
	} else {
		l := binDec(bin[pc+7 : pc+7+11])
		pc = pc + 7 + 11
		for i := int64(0); i < l; i++ {
			nextpc, addV, val := parse(bin, pc)
			version += addV
			pc = nextpc
			subVals = append(subVals, val)
		}
	}
	var s int64
	switch pkgType {
	case 0:
		for _, v := range subVals {
			s += v
		}
	case 1:
		s = 1
		for _, v := range subVals {
			s *= v
		}
	case 2:
		s = math.MaxInt64
		for _, v := range subVals {
			s = min(s, v)
		}
	case 3:
		for _, v := range subVals {
			s = max(s, v)
		}
	case 5:
		if subVals[0] > subVals[1] {
			s = 1
		}
	case 6:
		if subVals[0] < subVals[1] {
			s = 1
		}
	case 7:
		if subVals[0] == subVals[1] {
			s = 1
		}
	}
	return pc, version, s
}

func hexBin(in string) string {
	decoded, err := hex.DecodeString(in)
	if err != nil {
		log.Fatal(err)
	}
	var res string
	for _, d := range decoded {
		res += fmt.Sprintf("%08b", d)
	}
	return res
}

func binDec(bin string) int64 {
	num, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
