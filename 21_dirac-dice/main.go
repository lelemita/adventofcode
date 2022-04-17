/*
	** example -> 739785
	Player 1 starting position: 4
	Player 2 starting position: 8

	** input
	Player 1 starting position: 2
	Player 2 starting position: 8
*/

package main

import "fmt"

type Player struct {
	loc   int
	Score int
}

func (p *Player) drowDice(t int) {
	p.loc = (p.loc + (9*t + 6)) % 10
	if p.loc == 0 {
		p.Score += 10
	} else {
		p.Score += p.loc
	}
}

func solution() int {
	A := Player{loc: 2, Score: 0}
	B := Player{loc: 8, Score: 0}
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

func main() {
	fmt.Println(solution())
}
