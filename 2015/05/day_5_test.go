package main

import (
	"testing"
)

func Test1(t *testing.T) {
	testInputs := []string{"ugknbfddgicrmopn", "aaa"}
	if SolveDay5(testInputs, Part1) != 2 {
		t.Error(`SolveDay5(testInput, Part1) != 2`)
	}
}

func Test2(t *testing.T) {
	testInputs := []string{"jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb"}
	if SolveDay5(testInputs, Part1) != 0 {
		t.Error(`SolveDay5(testInput, Part1) != 0`)
	}
}

func Test3(t *testing.T) {
	testInputs := []string{"qjhvhtzxzqqjkmpb", "xxyxx", "qxxyxx", "xxyxxq", "xxxyxxx", "xyxqjqj"}
	if SolveDay5(testInputs, Part2) != 6 {
		t.Error(`SolveDay5(testInput, Part2) != 5`)
	}
}

func Test4(t *testing.T) {
	testInputs := []string{"uurcxstgmygtbstg", "ieodomkazucvgmuy", "xxyaxx", "aaaaa", "xxx"}
	if SolveDay5(testInputs, Part2) != 0 {
		t.Error(`SolveDay5(testInput, Part2) != 0`)
	}
}
