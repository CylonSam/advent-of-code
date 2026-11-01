package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

type ChallengePart int

const (
	Part1 ChallengePart = 1
	Part2 ChallengePart = 2
)

func main() {
	input := "iwrupvqb"
	fmt.Println(SolveDay4(input, Part2))
}

func SolveDay4(input string, challengePart ChallengePart) int {
	prefix := "00000"
	cutoff := 5
	if challengePart == Part2 {
		cutoff = 6
		prefix = "000000"
	}

	num := 1
	for {
		numString := strconv.Itoa(num)

		key := input + numString

		md5Hash := md5.New()
		io.WriteString(md5Hash, key)
		hashed := fmt.Sprintf("%x", md5Hash.Sum(nil))

		if hashed[0:cutoff] == prefix {
			return num
		}

		num += 1
	}

}
