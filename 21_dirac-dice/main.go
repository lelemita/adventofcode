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
	"sync"
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

var wg sync.WaitGroup
var mutex sync.Mutex
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

func (p *Player) SplitUniverse(num int, next Player) {
	// fmt.Printf("%d (num: %d) loc:%d, Score:%d -> ", p.Index, num, p.loc, p.Score)
	p.move(num)
	// fmt.Printf("%d / %d\r\n", p.loc, p.Score)
	if p.Score >= 21 {
		mutex.Lock()
		winCount[p.Index] += 1
		mutex.Unlock()
	} else {
		for _, v := range [27]int{3, 4, 5, 4, 5, 6, 5, 6, 7, 4, 5, 6, 5, 6, 7, 6, 7, 8, 5, 6, 7, 6, 7, 8, 7, 8, 9} {
			wg.Add(1)
			clone := next.clone()
			go clone.SplitUniverse(v, p.clone())
		}
	}
	wg.Done()
}

func part02(startA, startB int) int {
	winCount = map[int]int{0: 0, 1: 0}
	A := Player{loc: startA, Score: 0, Index: 0}
	B := Player{loc: startB, Score: 0, Index: 1}
	// map[3:1 4:3 5:6 6:7 7:6 8:3 9:1]
	// for _, v := range [27]int{3, 4, 5, 6, 7, 8, 9} {
	for _, v := range [27]int{3, 4, 5, 4, 5, 6, 5, 6, 7, 4, 5, 6, 5, 6, 7, 6, 7, 8, 5, 6, 7, 6, 7, 8, 7, 8, 9} {
		wg.Add(1)
		clone := A.clone()
		go clone.SplitUniverse(v, B.clone())
	}

	wg.Wait()
	fmt.Println("--------------------")
	for k, v := range winCount {
		fmt.Println(k, v)
	}
	return -1
}

func main() {
	// fmt.Println(part01(4, 8))
	// fmt.Println(part01(2, 8))

	fmt.Println(part02(4, 8))
}
