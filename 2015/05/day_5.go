package main

import (
	"fmt"
	"strings"

	"github.com/cylonsam/advent-of-code/2015/auxlib"
)

func main() {
	filePath := "/home/samlas/workspace/advent-of-code/2015/05/input.csv"
	csvRecords := auxlib.ReadCsv(filePath)
	inputs := parseInputs(csvRecords)
	fmt.Println(SolveDay5(inputs, auxlib.Part2))
}

func SolveDay5(inputs []string, challengePart auxlib.ChallengePart) int {
	niceCount := 0

	if challengePart == auxlib.Part1 {
		for inputIndex := range len(inputs) {
			if isNice1(inputs[inputIndex]) {
				niceCount += 1
			}
		}
	} else {
		for inputIndex := range len(inputs) {
			if isNice2(inputs[inputIndex]) {
				niceCount += 1
			}
		}
	}

	return niceCount
}

func isNice1(str string) bool {

	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}
	// 1. check forbidden string
	for forbiddenIndex := range len(forbiddenStrings) {
		if strings.Contains(str, forbiddenStrings[forbiddenIndex]) {
			// fmt.Println("forbidden!")
			return false
		}
	}

	// 2. check vowel count >= 3
	vowelCount := 0
	vowels := "aeiou"

	for j := range len(str) {
		for vowelIndex := range len(vowels) {
			if str[j:j+1] == vowels[vowelIndex:vowelIndex+1] {
				vowelCount += 1
			}
		}
	}
	// fmt.Println(vowelCount)
	if vowelCount < 3 {
		return false
	}

	// 3. check successive letters
	successiveLetterCount := 0
	inputSize := len(str)
	for i := range inputSize {
		if i+1 < inputSize && str[i:i+1] == str[i+1:i+2] {
			successiveLetterCount += 1
		}
	}

	// fmt.Println(successiveLetterCount)
	if successiveLetterCount == 0 {
		return false
	}

	return true

}

func isNice2(str string) bool {
	pairRepeat := false

	var currentPair string
	for i := 0; i < len(str); i += 1 {
		if i == len(str)-1 || len(str) <= 3 {
			break
		}

		currentPair = str[i : i+2]
		fmt.Println(currentPair)
		fmt.Printf("i:%d\n", i)

		otherPairRelativeIndex := strings.Index(str[i+2:], currentPair)

		if otherPairRelativeIndex != -1 {
			pairRepeat = true
		}

		if pairRepeat {
			break
		}
	}

	// fmt.Println("aaaaaaaaaa")
	uniqueLetterBetween := false
	for j := range len(str) {
		if j == len(str)-2 {
			break
		}
		if str[j] == str[j+2] {
			uniqueLetterBetween = true
			break
		}
	}
	return (true == pairRepeat) && (true == uniqueLetterBetween)
}

func parseInputs(csvRecords [][]string) []string {
	result := []string{}

	for i := range len(csvRecords) {
		result = append(result, csvRecords[i][0])
	}

	return result
}
