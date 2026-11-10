package main

import (
	"testing"

	"github.com/cylonsam/advent-of-code/2015/auxlib"
)

func TestPart1(t *testing.T) {
	testInputs := []string{"turn on 0,0 through 999,999", "toggle 0,0 through 999,0", "turn off 499,499 through 500,500"}
	if SolveDay6(testInputs, auxlib.Part1) != 998996 {
		t.Error(`SolveDay6(testInput, Part1) != 998996`)
	}
}

func TestPart2(t *testing.T) {
	testInputs := []string{"turn on 0,0 through 0,0", "toggle 0,0 through 999,999"}
	if SolveDay6(testInputs, auxlib.Part2) != 2000001 {
		t.Error(`SolveDay6(testInput, Part2) != 2000001`)
	}
}
