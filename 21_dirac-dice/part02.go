package main

import "fmt"

type state struct {
	player         int
	dice           int
	score0, score1 int
	loc0, loc1     int
}

type win struct {
	win0 int
	win1 int
}

var cache map[state]win
var nums map[int]int

const winScore = 21

func play(s state) win {
	if val, isExist := cache[s]; isExist {
		return val

	}
	initS := s
	switch s.player {
	case 0:
		s.loc0, s.score0 = getScore(s.loc0, s.score0, s.dice)
		s.player = 1
		if s.score0 >= winScore {
			cache[initS] = win{win0: 1, win1: 0}
			return win{win0: 1, win1: 0}
		}
	case 1:
		s.loc1, s.score1 = getScore(s.loc1, s.score1, s.dice)
		s.player = 0
		if s.score1 >= winScore {
			cache[initS] = win{win0: 0, win1: 1}
			return win{win0: 0, win1: 1}
		}
	}

	var w0, w1 int
	for n, cnt := range nums {
		s.dice = n
		res := play(s)
		w0 += res.win0 * cnt
		w1 += res.win1 * cnt
	}

	cache[initS] = win{win0: w0, win1: w1}
	return win{win0: w0, win1: w1}
}

func getScore(loc, score, dice int) (int, int) {
	loc = (loc + dice) % 10
	if loc == 0 {
		score += 10
	} else {
		score += loc
	}
	return loc, score
}

func part2(start0, start1 int) {
	var w0, w1 int
	for n, cnt := range nums {
		s := state{
			player: 0,
			dice:   n,
			score0: 0,
			score1: 0,
			loc0:   start0,
			loc1:   start1,
		}
		res := play(s)
		w0 += res.win0 * cnt
		w1 += res.win1 * cnt
	}

	fmt.Println("-----------------")
	fmt.Println("0: ", w0)
	fmt.Println("1: ", w1)
	fmt.Println("-----------------")
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
	cache = make(map[state]win)

	fmt.Println(nums)
	part2(4, 8)

	fmt.Println("21: 444356092776315")
}
