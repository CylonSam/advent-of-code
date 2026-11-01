package main

import (
	"testing"
)

func Test1(t *testing.T) {
	if Solve(">", Part1) != 2 {
		t.Error(`Solve(">", Part1) != 2`)
	}
}
