package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var input_test string
var init_test = [][]rune{
	{'Z', 'N'},
	{'M', 'C', 'D'},
	{'P'},
}

func TestPart1(t *testing.T) {
	result := Part1(string(input_test), init_test)
	expected := "CMZ"
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(string(input_day), init_day)
	expected := "LJSVLTWQM"
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(string(input_test), init_test)
	expected := "MCD"
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := Part2(string(input_day), init_day)
	expected := "BRQWDBBJM"
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(input_test, init_test)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(input_day, init_day)
	}
}
