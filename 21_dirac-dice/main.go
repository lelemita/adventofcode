/*
	** example -> 739785, 444356092776315
	Player 1 starting position: 4
	Player 2 starting position: 8

	** input
	Player 1 starting position: 2
	Player 2 starting position: 8
*/

package main

import (
	"fmt"
)

type Player struct {
	loc   int
	Score int
	Index int
}

func (p *Player) drowDice(t int) {
	p.loc = (p.loc + (9*t + 6)) % 10
	if p.loc == 0 {
		p.Score += 10
	} else {
		p.Score += p.loc
	}
}

func part01(startA, startB int) int {
	A := Player{loc: startA, Score: 0}
	B := Player{loc: startB, Score: 0}
	cnt := 0
	for true {
		if cnt%2 == 0 {
			A.drowDice(cnt)
		} else {
			B.drowDice(cnt)
		}
		cnt += 1

		if A.Score >= 1000 {
			fmt.Println(cnt, A.Score, B.Score)
			return 3 * cnt * B.Score
		}
		if B.Score >= 1000 {
			fmt.Println(cnt, A.Score, B.Score)
			return 3 * cnt * A.Score
		}
	}
	return -1
}

var winCount map[int]int

func (p *Player) clone() Player {
	return Player{loc: p.loc, Score: p.Score, Index: p.Index}
}

func (p *Player) move(num int) {
	p.loc = (p.loc + num) % 10
	if p.loc == 0 {
		p.Score += 10
	} else {
		p.Score += p.loc
	}
}

var cache map[string][2]int
var nums map[int]int

const winScore int = 3

func (p *Player) SplitUniverse(num, cnt int, next Player) {
	status := fmt.Sprintf("%d_%d_%d_%d_%d", p.Index, p.Score, p.loc, next.Score, next.loc)
	fmt.Println(num, status)
	if res, isExist := cache[status]; isExist {
		fmt.Println(">> ", status)
		winCount[0] += res[0]
		winCount[1] += res[1]
	} else {
		p.move(num)
		if p.Score >= winScore {
			// winCount[p.Index] += 1
			res := [2]int{0, 0}
			res[p.Index] = cnt
			cache[status] = res
		} else {
			for k, cnt := range nums {
				clone := next.clone()
				clone.SplitUniverse(k, cnt, p.clone())
			}
		}
	}
}

func part02(startA, startB int) int {
	nums = getNums()
	cache = map[string][2]int{}
	winCount = map[int]int{0: 0, 1: 0}
	A := Player{loc: startA, Score: 0, Index: 0}
	B := Player{loc: startB, Score: 0, Index: 1}
	for k, cnt := range nums {
		clone := A.clone()
		clone.SplitUniverse(k, cnt, B.clone())
	}

	fmt.Println("--------------------")
	for k, v := range winCount {
		fmt.Println(k, v)
	}
	return -1
}

func getNums() map[int]int {
	arr := []int{}
	cnts := map[int]int{}
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				arr = append(arr, i+j+k)
				cnts[i+j+k] += 1
			}
		}
	}
	return cnts
}

func main() {
	// fmt.Println(part01(4, 8))
	// fmt.Println(part01(2, 8))

	fmt.Println(part02(4, 8))
}
