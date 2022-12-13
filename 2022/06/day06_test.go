package main

import (
	_ "embed"
	"testing"
)

type testCase struct {
	input   string
	answer1 int
	answer2 int
}

var inputTests []testCase = []testCase{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7, 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5, 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6, 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, 26},
}

func TestPart1(t *testing.T) {
	for _, tc := range inputTests {
		result := Part1(string(tc.input))
		if result != tc.answer1 {
			t.Errorf("Result is incorrect, got: %v, want: %v.", result, tc.answer1)
		}
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(string(input_day))
	expected := 1848
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	for _, tc := range inputTests {
		result := Part2(string(tc.input))
		if result != tc.answer2 {
			t.Errorf("Result is incorrect, got: %v, want: %v.", result, tc.answer2)
		}
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(string(input_day))
	expected := 2308
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input_day)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(input_day)
	}
}
