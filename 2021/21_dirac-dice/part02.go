// struct, cache 참고: https://github.com/pemoreau/advent-of-code/blob/main/go/2021/21/day21.go
// 재미있는 oop 풀이: https://github.com/CodeHex/advent-of-code-2021/blob/main/day21/main.go
// 의문) 배열도 키가 되고, 3개 합으로 계산해도 되는거면... 처음 풀이는 어디가 틀렸던 거지.....ㅠㅠ

package main

import "fmt"

type state struct {
	player int
	dice   int
	score  [2]int
	loc    [2]int
}

var cache map[state][2]int
var nums map[int]int

const winScore = 21

func (s *state) move() {
	i := s.player
	s.loc[i] = (s.loc[i] + s.dice) % 10
	if s.loc[i] == 0 {
		s.score[i] += 10
	} else {
		s.score[i] += s.loc[i]
	}
}

func (s *state) changePlayer() {
	if s.player == 0 {
		s.player = 1
	} else {
		s.player = 0
	}
}

func play(s state) [2]int {
	if val, isExist := cache[s]; isExist {
		return val

	}
	initS := s

	s.move()
	if s.score[s.player] >= winScore {
		win := [2]int{0, 0}
		win[s.player] = 1
		cache[initS] = win
		return win
	}
	s.changePlayer()

	win := [2]int{0, 0}
	for n, cnt := range nums {
		s.dice = n
		res := play(s)
		win[0] += res[0] * cnt
		win[1] += res[1] * cnt
	}

	cache[initS] = win
	return win
}

func part2(start0, start1 int) int {
	var w0, w1 int
	for n, cnt := range nums {
		s := state{
			player: 0,
			dice:   n,
			score:  [2]int{0, 0},
			loc:    [2]int{start0, start1},
		}
		res := play(s)
		w0 += res[0] * cnt
		w1 += res[1] * cnt
	}

	if w0 > w1 {
		return w0
	}
	return w1
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
	nums = getNums()
	cache = make(map[state][2]int)

	result := part2(4, 8)
	fmt.Println("PART02: ", result)
	fmt.Println(444356092776315 == result)
}
