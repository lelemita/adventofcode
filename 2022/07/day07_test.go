package main

import (
	_ "embed"
	"testing"
)

//go:embed input_test.txt
var input_test string

//go:embed input_test2.txt
var input_test2 string

func TestPart1(t *testing.T) {
	result := Part1(string(input_test))
	expected := 95437
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

// 동일한 폴더 이름 이 있는 예시
func TestPart1_9999(t *testing.T) {
	result := Part1(string(input_test2))
	expected := 99999
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestPart1Input(t *testing.T) {
	result := Part1(string(input_day))
	expected := 0
	if result != expected {
		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
	}
}

//func TestPart2(t *testing.T) {
//	result := Part2(string(input_test))
//	expected := 0
//	if result != expected {
//		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
//	}
//}

//func TestPart2Input(t *testing.T) {
//	result := Part2(string(input_day))
//	expected := 0
//	if result != expected {
//		t.Errorf("Result is incorrect, got: %v, want: %v.", result, expected)
//	}
//}

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
