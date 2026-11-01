package main

import (
	"testing"
)

func Test1(t *testing.T) {
	testInput := "abcdef"
	if SolveDay4(testInput, Part1) != 609043 {
		t.Error(`SolveDay4(testInput) != 609043`)
	}
}

func Test2(t *testing.T) {
	testInput := "pqrstuv"
	if SolveDay4(testInput, Part1) != 1048970 {
		t.Error(`SolveDay4(testInput) != 1048970`)
	}
}
